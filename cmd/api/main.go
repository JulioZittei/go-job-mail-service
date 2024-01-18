package main

import (
	"net/http"

	"github.com/JulioZittei/go-job-mail-service/internal/api/v1/controller"
	exceptionhandler "github.com/JulioZittei/go-job-mail-service/internal/api/v1/exceptionHandler"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/service"
	"github.com/JulioZittei/go-job-mail-service/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := service.CampaignServiceImpl{
		Repository: &database.CampaignRepository{},
	}

	controller := &controller.CampaignController{
		CampaignService: &service,
	}
	
	r.Post("/campaign", exceptionhandler.ExceptionHandler(controller.CampaignPost))

	http.ListenAndServe(":3000", r)
}