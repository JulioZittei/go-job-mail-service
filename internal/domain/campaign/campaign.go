package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID string
	Name string
	Content string
	Contacts []Contact
	CreatedAt time.Time 
}

func NewCampaign(name string, content string, emails []string) *Campaign {
	contacts := make([]Contact, len(emails))
	for i, value := range emails {
		contacts[i] = Contact{
			Email: value,
		}
	}
	
	return &Campaign{
		ID: xid.New().String(),
		Name: name,
		Content: content,
		CreatedAt: time.Now(),
		Contacts: contacts,
	}
}