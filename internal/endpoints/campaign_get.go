package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 200)
	campaigns, err := h.CampaignService.Repository.ListAll()
	if err != nil {
		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.JSON(w, r, campaigns)
}
