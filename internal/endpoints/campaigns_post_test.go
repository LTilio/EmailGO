package endpoints

import (
	"EmailGO/internal/contract"
	internalMock "EmailGO/internal/test/internal-mock"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(body contract.NewCampaign, CreatedByExpected string) (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	ctx := context.WithValue(req.Context(), emailContextKey, CreatedByExpected)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	return req, rr

}

func Test_campaignsPost_saveNewCampaign(t *testing.T) {
	assert := assert.New(t)
	CreatedByExpected := "teste1@email.com"
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "teste de conteudo",
		Emails:  []string{"teste@email.com"},
	}
	service := new(internalMock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name && request.Content == body.Content && request.CreatedBy == CreatedByExpected {
			return true
		} else {
			return false
		}
	})).Return("2x", nil)
	handler := Handler{CampaignService: service}

	req, rr := setup(body, CreatedByExpected)

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_campaignsPost_InformeError(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "teste de conteudo",
		Emails:  []string{"teste@email.com"},
	}
	service := new(internalMock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	req, rr := setup(body, "teste@email.com")

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)

}
