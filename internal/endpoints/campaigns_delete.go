package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

// CampaignDelete godoc
// @Summary Deletar campanha por ID
// @Description Remove uma campanha com base no ID fornecido
// @Tags campaigns
// @Accept json
// @Produce json
// @Param id path string true "ID da campanha" example:"12345"
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} map[string]string "Mensagem de sucesso"
// @Failure 401 {object} internalerror.ErrorResponse "Não autorizado"
// @Failure 404 {object} internalerror.ErrorResponse "Campanha não encontrada"
// @Failure 500 {object} internalerror.ErrorResponse "Erro interno"
// @Router /campaigns/{id} [delete]
func (h *Handler) CampaignDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Delete(id)

	return nil, 200, err
}
