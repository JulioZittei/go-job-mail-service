package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JulioZittei/go-job-mail-service/internal/api/v1/controller"
	exceptionhandler "github.com/JulioZittei/go-job-mail-service/internal/api/v1/exceptionHandler"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/service"
	"github.com/JulioZittei/go-job-mail-service/internal/infrastructure/database"
	"github.com/JulioZittei/go-job-mail-service/internal/infrastructure/mail"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := service.CampaignServiceImpl{
		Repository: &database.CampaignRepository{
			Db: database.NewDB(),
		},
		SendMail: mail.SendMail,
	}

	campaigncController := &controller.CampaignController{
		CampaignService: &service,
	}

	r.Route("/campaign", func(r chi.Router) {
		r.Use(controller.Auth)
		r.Post("/", exceptionhandler.ExceptionHandler(campaigncController.CampaignPost))
		r.Get("/{id}", exceptionhandler.ExceptionHandler(campaigncController.CampaignGetById))
		r.Delete("/{id}", exceptionhandler.ExceptionHandler(campaigncController.CampaignDelete))
		r.Patch("/{id}", exceptionhandler.ExceptionHandler(campaigncController.CampaignStart))
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		fmt.Println("Starting Server")
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			panic(err)
		}

	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)

	fmt.Println("Started")

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("ServerStopping...")

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Server Stopped")
}
