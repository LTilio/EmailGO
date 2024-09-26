package endpoints

import (
	"EmailGO/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

// CampaignPost godoc
// @Summary Criar nova campanha
// @Description Cria uma nova campanha de envio de emails
// @Tags campaigns
// @Accept json
// @Produce json
// @Param campaign body contract.NewCampaignRequest true "Dados da nova campanha"
// @Param Authorization header string true "Bearer {token}"
// @Success 201 {object} map[string]string "Campanha criada"
// @Failure 400 {object} internalerror.ErrorResponse "Requisição inválida"
// @Failure 500 {object} internalerror.ErrorResponse "Erro interno"
// @Router /campaigns [post]
func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaignRequest
	render.DecodeJSON(r.Body, &request)
	email := r.Context().Value(emailContextKey).(string)
	request.CreatedBy = email
	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id}, 201, err
}
