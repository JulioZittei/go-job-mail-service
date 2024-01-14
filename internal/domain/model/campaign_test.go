package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name = "Campaign X"
	content = "Body"
	contacts = []string{"john@email.com", "mary@email.com"}
	createdAt = time.Now()
)

func TestShouldCreateANewCampaign(t *testing.T) {
	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(t, campaign.ID)
	assert.GreaterOrEqual(t, campaign.CreatedAt, createdAt)
	assert.Equal(t, name, campaign.Name)
	assert.Equal(t, content, campaign.Content )
	assert.Equal(t, len(contacts), len(campaign.Contacts))
}

func TestShouldValidateName(t *testing.T) {
	_, err := NewCampaign("", content, contacts)

	assert.Equal(t,"name is required", err.Error())
}

func TestShouldValidateContent(t *testing.T) {
	_, err := NewCampaign(name, "", contacts)

	assert.Equal(t,"content is required", err.Error())
}

func TestShouldValidateContacts(t *testing.T) {
	_, err := NewCampaign(name, content, []string{})

	assert.Equal(t,"contacts is required", err.Error())
}