package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/walleksmr/golang-emailn/internal/contract"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	return map[string]string{"id": id}, 201, err
}
