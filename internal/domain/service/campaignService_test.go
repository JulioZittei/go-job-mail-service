package service

import (
	"errors"
	"testing"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	mockTests "github.com/JulioZittei/go-job-mail-service/internal/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



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
	mockedRepository := new(mockTests.CampaignRepositoryMock)
	
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

func TestShouldReturnErrorWhenRepositorySave(t *testing.T) {
	assert := assert.New(t)

	mockedRepository := new(mockTests.CampaignRepositoryMock)
	
	mockedRepository.On("Save", mock.Anything).Return(errors.New("error while saving campaign on database"))

	service.Repository = mockedRepository
	id, err := service.Create(newCampaign)

	expectedError := internalerrors.ErrInternal{}

	assert.Error(err)
	assert.Empty(id)
	assert.Equal(expectedError.Error(), err.Error())
}

func TestShouldGetCampaignById(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(mockTests.CampaignRepositoryMock)
	
	createdCampaign, _ := model.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	mockedRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == createdCampaign.ID
	})).Return(createdCampaign, nil)

	service.Repository = mockedRepository
	campaign, err := service.GetById(createdCampaign.ID)

	assert.Nil(err)
	assert.NotNil(campaign)
	assert.Equal(createdCampaign.ID, campaign.ID )
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErroWhenRepositoryGetCampaignById(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(mockTests.CampaignRepositoryMock)

	expectedError := internalerrors.NewErrInternal()

	mockedRepository.On("GetById", mock.Anything).Return(nil, errors.New("error while finding campaign"))

	service.Repository = mockedRepository
	_, err := service.GetById("idTest")

	assert.NotNil(err)
	assert.Equal(expectedError.Error(), err.Error() )
	mockedRepository.AssertExpectations(t)
}