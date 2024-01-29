package main

import (
	"log"
	"time"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/service"
	"github.com/JulioZittei/go-job-mail-service/internal/infrastructure/database"
	"github.com/JulioZittei/go-job-mail-service/internal/infrastructure/mail"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDB()
	repository := database.CampaignRepository{
		Db: db,
	}
	service := service.CampaignServiceImpl{
		Repository: &repository,
		SendMail:   mail.SendMail,
	}

	for {
		campaigns, err := repository.GetCampaignToBeSent()
		if err != nil {
			println(err.Error())
		}

		println("Amount of campaigns: ", len(campaigns))

		for _, campaign := range campaigns {
			go service.SendMailAndUpdateStatus(&campaign)
			println("Campaign sent: ", campaign.ID)
		}
		time.Sleep(10 * time.Second)
	}
}
