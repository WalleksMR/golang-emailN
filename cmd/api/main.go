package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/walleksmr/golang-emailn/internal/contract"
	"github.com/walleksmr/golang-emailn/internal/domain/campaign"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.Service{}
	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaign
		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		id, err := service.Create(request)
		if err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
	})

	fmt.Println("API is running on port 3000")
	http.ListenAndServe(":3000", r)
}
