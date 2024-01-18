package model

import (
	"testing"
	"time"

	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name = "Campaign X"
	content = "Body Content"
	contacts = []string{"john@email.com", "mary@email.com"}
	createdAt = time.Now()
	fake = faker.New()
)

func TestShouldCreateANewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil( campaign.ID)
	assert.GreaterOrEqual( campaign.CreatedAt, createdAt)
	assert.Equal(name, campaign.Name)
	assert.Equal( content, campaign.Content )
	assert.Equal( len(contacts), len(campaign.Contacts))
	assert.Equal(Pending, campaign.Status)
}

func TestShouldValidateNameMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", content, contacts)

	assert.Equal("name", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at least 5.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldValidateNameMax(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at most 24.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}


func TestShouldValidateContentMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at least 5.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldValidateContactsMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at least 1.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at most 1024.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldValidateContactEmail(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{"email"})

	assert.Equal("email", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be well-formed.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}