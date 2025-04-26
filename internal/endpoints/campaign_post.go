package endpoints

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/walleksmr/golang-emailn/internal/contract"
	"github.com/walleksmr/golang-emailn/internal/excptions"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaign
	err := render.DecodeJSON(r.Body, &request)
	if err != nil {
		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	id, err := h.CampaignService.Create(request)
	if err != nil {
		status := 400
		if errors.Is(err, excptions.ErrInternal) {
			status = 500
		}
		render.Status(r, status)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}
