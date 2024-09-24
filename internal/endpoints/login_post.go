package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func Login(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid request")
	}

	token, err := getKeycloakToken(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return nil, http.StatusUnauthorized, errors.New("invalid credentials")
	}

	return TokenResponse{AccessToken: token}, http.StatusOK, nil
}

func getKeycloakToken(username, password string) (string, error) {
	identityProvider := "http://localhost:8080/realms/provideremailgo/protocol/openid-connect/token"
	clientId := "emailgo"
	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("username", username)
	data.Set("password", password)
	data.Set("grant_type", "password")

	req, err := http.NewRequest("POST", identityProvider, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to authenticate with Keycloak")
	}
	defer resp.Body.Close()

	var tokenResponse TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
