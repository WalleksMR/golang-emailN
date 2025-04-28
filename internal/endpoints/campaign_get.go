package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.ListAll()
	return campaigns, 200, err
}

func (h *Handler) CampaignGetOne(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	campaing, err := h.CampaignService.GetById(id)
	return campaing, 200, err
}
