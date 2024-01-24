package model

import (
	"time"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/util/validator"
	"github.com/rs/xid"
)

type Contact struct {
	ID         string `gorm:"primaryKey,size:50"`
	Email      string `gorm:"size:100" validate:"email"`
	CampaignId string `gorm:"size:50"`
}

const (
	Pending  string = "PENDING"
	Canceled string = "CANCELED"
	Deleted  string = "DELETED"
	Started  string = "STARTED"
	Done     string = "DONE"
)

type Campaign struct {
	ID        string     `gorm:"primaryKey,size:50" validate:"required"`
	Name      string     `gorm:"size:100" validate:"min=5,max=24"`
	Content   string     `gorm:"size:1024" validate:"min=5,max=1024"`
	Contacts  []*Contact `validate:"min=1,dive"`
	Status    string     `gorm:"size:20"`
	CreatedAt time.Time  `validate:"required"`
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Delete() {
	c.Status = Deleted
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]*Contact, len(emails))
	for i, value := range emails {
		contacts[i] = &Contact{
			ID:    xid.New().String(),
			Email: value,
		}
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
	}

	err := validator.ValidateStruct(campaign)
	if err == nil {
		return campaign, err
	}
	return nil, err
}
