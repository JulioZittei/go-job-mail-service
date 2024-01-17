package service

import (
	"errors"
	"testing"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func(r *repositoryMock) Save(campaign *model.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func(r *repositoryMock) Get() ([]model.Campaign, error) {
	args := r.Called()
	return []model.Campaign{}, args.Error(0)
}

var (
	newCampaign = &contract.NewCampaignInput{
	Name: "Campaign X",
	Content: "Body Content",
	Emails: []string{"john@mail.com", "mary@mail.com"},
	}
	service = CampaignServiceImpl{}
)

func TestShouldCreateAndSaveCampaign(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(repositoryMock)
	
	mockedRepository.On("Save", mock.MatchedBy(func(campaign *model.Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Repository = mockedRepository
	id, err := service.Create(newCampaign)

	assert.Nil(err)
	assert.NotEmpty(id)
	mockedRepository.AssertExpectations(t)
}

func TestShouldValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(&contract.NewCampaignInput{
		Name: "",
	Content: newCampaign.Content,
	Emails: newCampaign.Emails,
	})

	assert.Error(err)
	assert.Empty(id)
	assert.Equal(err.Error(), "validation error")
}

func TestShouldValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	mockedRepository := new(repositoryMock)
	
	mockedRepository.On("Save", mock.Anything).Return(errors.New("error while saving campaign on database"))

	service.Repository = mockedRepository
	id, err := service.Create(newCampaign)

	expectedError := internalerrors.ErrInternal{}

	assert.Error(err)
	assert.Empty(id)
	assert.Equal(expectedError.Error(), err.Error())
}