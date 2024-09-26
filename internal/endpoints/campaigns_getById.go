package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

// CampaignGetById godoc
// @Summary Obter campanha por ID
// @Description Recupera os detalhes de uma campanha com base no ID fornecido
// @Tags campaigns
// @Accept json
// @Produce json
// @Param id path string true "ID da campanha" example:"12345"
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} contract.CampaignResponse "Detalhes da campanha"
// @Failure 401 {object} internalerror.ErrorResponse "Não autorizado"
// @Failure 404 {object} internalerror.ErrorResponse "Campanha não encontrada"
// @Failure 500 {object} internalerror.ErrorResponse "Erro interno"
// @Router /campaigns/{id} [get]
func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	campaign, err := h.CampaignService.GetBy(id)
	return campaign, 200, err

}
