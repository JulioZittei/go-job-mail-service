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
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil( campaign.ID)
	assert.GreaterOrEqual( campaign.CreatedAt, createdAt)
	assert.Equal(name, campaign.Name)
	assert.Equal( content, campaign.Content )
	assert.Equal( len(contacts), len(campaign.Contacts))
}

func TestShouldValidateName(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func TestShouldValidateContent(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func TestShouldValidateContacts(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required", err.Error())
}