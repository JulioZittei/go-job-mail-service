package model

import (
	"errors"
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
	Contacts []*Contact
	CreatedAt time.Time 
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if len(name) == 0 {
		return nil, errors.New("name is required")
	} else if len(content) == 0 {
		return nil, errors.New("content is required")
	} else if len(emails) == 0 {
		return nil, errors.New("contacts is required")
	}

	contacts := make([]*Contact, len(emails))
	for i, value := range emails {
		contacts[i] = &Contact{
			Email: value,
		}
	}
	
	return &Campaign{
		ID: xid.New().String(),
		Name: name,
		Content: content,
		CreatedAt: time.Now(),
		Contacts: contacts,
	}, nil
}