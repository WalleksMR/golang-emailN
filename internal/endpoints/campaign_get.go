package endpoints

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.ListAll()
	return campaigns, 200, err
}

func (h *Handler) CampaignGetOne(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, 400, errors.New("id is required")
	}
	campaing, err := h.CampaignService.GetById(id)
	return campaing, 200, err
}
