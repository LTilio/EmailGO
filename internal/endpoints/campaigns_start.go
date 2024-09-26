package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

// CampaignStart godoc
// @Summary Iniciar campanha
// @Description Altera o status da campanha para "Iniciada" com base no ID fornecido
// @Tags campaigns
// @Accept json
// @Produce json
// @Param id path string true "ID da campanha" example:"12345"
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} contract.CampaignResponse "Campanha iniciada com sucesso"
// @Failure 401 {object} internalerror.ErrorResponse "Não autorizado"
// @Failure 404 {object} internalerror.ErrorResponse "Campanha não encontrada"
// @Failure 500 {object} internalerror.ErrorResponse "Erro interno"
// @Router /campaigns/{id}/start [patch]
func (h *Handler) CampaignStart(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Start(id)

	return nil, 200, err
}
