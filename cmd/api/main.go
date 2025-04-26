package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/walleksmr/golang-emailn/internal/domain/campaign"
	"github.com/walleksmr/golang-emailn/internal/endpoints"
	"github.com/walleksmr/golang-emailn/internal/infra/database"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	repository := database.CampaignRepository{}
	campaignService := campaign.Service{Repository: &repository}
	handler := endpoints.Handler{CampaignService: &campaignService}
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))

	fmt.Println("API is running on port 3000")
	http.ListenAndServe(":3000", r)
}
