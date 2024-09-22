package main

import (
	"EmailGO/internal/campaign"
	"EmailGO/internal/endpoints"
	"EmailGO/internal/infra/database"
	"EmailGO/internal/infra/mail"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
		SendMail:   mail.SendMail,
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
		r.Patch("/{id}", endpoints.HandlerError(handler.CampaignStart))
	})

	http.ListenAndServe(":3000", r)
}
