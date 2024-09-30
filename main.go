package main

import (
	"EmailGO/internal/domain/campaign"
	"EmailGO/internal/endpoints"
	"EmailGO/internal/infra/config"
	"EmailGO/internal/infra/database"
	"EmailGO/internal/infra/mail"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"

	_ "EmailGO/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title EmailGo
// @version 1.0
// @description API para envio de emails em massa.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(config.CorsConfig()) // middleware para o cors

	db := database.NewDb()
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
		SendMail:   mail.SendMail,
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition
	))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Post("/login", endpoints.HandlerError(endpoints.Login))

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
		r.Patch("/{id}", endpoints.HandlerError(handler.CampaignStart))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
	})

	http.ListenAndServe(":3000", r)
}
