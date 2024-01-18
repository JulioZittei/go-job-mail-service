package model

import (
	"time"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/util/validator"
	"github.com/rs/xid"
)

type Contact struct {
	Email string `json:"email" validate:"email"`
}

const (
	Pending string = "PENDING"
	Started string = "STARTED"
	Done string = "DONE"
)

type Campaign struct {
	ID string `json:"id" validate:"required"`
	Name string `json:"name" validate:"min=5,max=24"`
	Content string `json:"content" validate:"min=5,max=1024"`
	Contacts []*Contact `json:"contacts" validate:"min=1,dive"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]*Contact, len(emails))
	for i, value := range emails {
		contacts[i] = &Contact{
			Email: value,
		}
	}
	
	campaign := &Campaign{
		ID: xid.New().String(),
		Name: name,
		Content: content,
		CreatedAt: time.Now(),
		Contacts: contacts,
		Status: Pending,
	}

	err := validator.ValidateStruct(campaign)
	if err == nil {
		return campaign, err
	}
	return nil, err
}