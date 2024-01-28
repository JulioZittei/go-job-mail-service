package mail

import (
	"fmt"
	"os"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"gopkg.in/gomail.v2"
)

func SendMail(campaign *model.Campaign) error {
	fmt.Println("Sending mail...")

	emails := make([]string, len(campaign.Contacts))
	for key, contact := range campaign.Contacts {
		emails[key] = contact.Email
	}

	dialer := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("EMAIL_USER"))
	message.SetHeader("To", emails...)
	message.SetHeader("Subject", campaign.Name)
	message.SetBody("text/html", campaign.Content)

	return dialer.DialAndSend(message)
}
