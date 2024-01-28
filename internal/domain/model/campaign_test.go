package model_test

import (
	"testing"
	"time"

	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name      = "Campaign X"
	content   = "Body Content"
	contacts  = []string{"john@email.com", "mary@email.com"}
	createdAt = time.Now()
	createdBy = "teste@teste.com"
	fake      = faker.New()
)

func TestShouldCreateANewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := model.NewCampaign(name, content, contacts, createdBy)

	assert.NotNil(campaign.ID)
	assert.GreaterOrEqual(campaign.CreatedAt, createdAt)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
	assert.Equal(model.Pending, campaign.Status)
	assert.Equal(createdBy, campaign.CreatedBy)
}

func TestShouldValidateNameMin(t *testing.T) {
	assert := assert.New(t)
	_, err := model.NewCampaign("", content, contacts, createdBy)

	assert.Equal("name", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at least 5.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldValidateNameMax(t *testing.T) {
	assert := assert.New(t)
	_, err := model.NewCampaign(fake.Lorem().Text(30), content, contacts, createdBy)

	assert.Equal("name", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at most 24.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldValidateContentMin(t *testing.T) {
	assert := assert.New(t)
	_, err := model.NewCampaign(name, "", contacts, createdBy)

	assert.Equal("content", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at least 5.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldValidateContactsMin(t *testing.T) {
	assert := assert.New(t)
	_, err := model.NewCampaign(name, content, []string{}, createdBy)

	assert.Equal("contacts", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at least 1.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := model.NewCampaign(name, fake.Lorem().Text(1040), contacts, createdBy)

	assert.Equal("content", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be at most 1024.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldValidateContactEmail(t *testing.T) {
	assert := assert.New(t)
	_, err := model.NewCampaign(name, content, []string{"email"}, createdBy)

	assert.Equal("email", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("must be well-formed.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}

func TestShouldChangeStatusToCanceled(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := model.NewCampaign(name, content, contacts, createdBy)

	campaign.Cancel()

	assert.Equal(model.Canceled, campaign.Status)
}

func TestShouldChangeStatusToDeleted(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := model.NewCampaign(name, content, contacts, createdBy)

	campaign.Delete()

	assert.Equal(model.Deleted, campaign.Status)
}

func TestShouldValidateCreatedByRequired(t *testing.T) {
	assert := assert.New(t)
	_, err := model.NewCampaign(name, content, contacts, "")

	assert.Equal("createdBy", err.(*internalerrors.ErrValidation).ErrorsParam[0].Param)
	assert.Equal("is required.", err.(*internalerrors.ErrValidation).ErrorsParam[0].Message)
}
