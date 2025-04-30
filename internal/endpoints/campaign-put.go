package endpoints

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/walleksmr/golang-emailn/internal/contract"
)

func (h *Handler) CampaignPutCancel(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.CampaingUpateInput
	request.ID = chi.URLParam(r, "id")
	if request.ID == "" {
		return nil, 400, errors.New("id is required")
	}

	render.DecodeJSON(r.Body, &request)

	err := h.CampaignService.Cancel(request)

	if err != nil {
		return nil, 400, err
	}

	return nil, 201, nil
}
