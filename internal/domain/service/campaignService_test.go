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
	"gorm.io/gorm"
)

var (
	newCampaign = &contract.NewCampaignInput{
		Name:      "Campaign X",
		Content:   "Body Content",
		Emails:    []string{"john@mail.com", "mary@mail.com"},
		CreatedBy: "teste@teste.com",
	}
	mockedRepository *mockTests.CampaignRepositoryMock
	service          = CampaignServiceImpl{}
)

func setup() {
	mockedRepository = new(mockTests.CampaignRepositoryMock)
	service.Repository = mockedRepository
}

func TestShouldCreateAndSaveCampaign(t *testing.T) {
	assert := assert.New(t)

	setup()

	mockedRepository.On("Save", mock.MatchedBy(func(campaign *model.Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	id, err := service.Create(newCampaign)

	assert.Nil(err)
	assert.NotEmpty(id)
	mockedRepository.AssertExpectations(t)
}

func TestShouldValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	setup()

	id, err := service.Create(&contract.NewCampaignInput{
		Name:    "",
		Content: newCampaign.Content,
		Emails:  newCampaign.Emails,
	})

	assert.Error(err)
	assert.Empty(id)
	assert.Equal(err.Error(), "validation error")
}

func TestShouldReturnErrorWhenRepositorySave(t *testing.T) {
	assert := assert.New(t)
	setup()

	mockedRepository.On("Save", mock.Anything).Return(errors.New("error while saving campaign on database"))
	id, err := service.Create(newCampaign)

	expectedError := internalerrors.ErrInternal{}

	assert.Error(err)
	assert.Empty(id)
	assert.Equal(expectedError.Error(), err.Error())
}

func TestShouldGetCampaignById(t *testing.T) {
	assert := assert.New(t)
	setup()

	createdCampaign, _ := model.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	mockedRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == createdCampaign.ID
	})).Return(createdCampaign, nil)

	campaign, err := service.GetById(createdCampaign.ID)

	assert.Nil(err)
	assert.NotNil(campaign)
	assert.Equal(createdCampaign.ID, campaign.ID)
	assert.Equal(createdCampaign.CreatedBy, campaign.CreatedBy)
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErroWhenRepositoryGetCampaignById(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedError := internalerrors.NewErrInternal()

	mockedRepository.On("GetById", mock.Anything).Return(nil, errors.New("error while finding campaign"))

	_, err := service.GetById("idTest")

	assert.NotNil(err)
	assert.Equal(expectedError.Error(), err.Error())
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErroWhenCampaignDoesNotExists(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedError := internalerrors.NewErrCampaignNotFound()

	mockedRepository.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	_, err := service.GetById("idTest")

	assert.NotNil(err)
	assert.Equal(expectedError.Error(), err.Error())
	mockedRepository.AssertExpectations(t)
}

func TestShouldDeleteCampaign(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedCampaign, _ := model.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	mockedRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == expectedCampaign.ID
	})).Return(expectedCampaign, nil)

	mockedRepository.On("Delete", mock.MatchedBy(func(campaign *model.Campaign) bool {
		if campaign.Name != expectedCampaign.Name || campaign.Content != expectedCampaign.Content || campaign.ID != expectedCampaign.ID || len(campaign.Contacts) != len(expectedCampaign.Contacts) {
			return false
		}
		return true
	})).Return(nil)

	err := service.Delete(expectedCampaign.ID)

	assert.Nil(err)
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErrorWhenGettingCampaignById(t *testing.T) {
	assert := assert.New(t)
	setup()

	mockedRepository.On("GetById", mock.Anything).Return(nil, errors.New("unexpected error getting by id"))

	err := service.Delete("idtest")

	assert.NotNil(err)
	assert.Equal(internalerrors.NewErrInternal().Error(), err.Error())
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnCampaignNotFoundError(t *testing.T) {
	assert := assert.New(t)
	setup()

	mockedRepository.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Delete("idtest")

	assert.NotNil(err)
	assert.Equal(internalerrors.NewErrCampaignNotFound().Error(), err.Error())
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErrorWhenStatusIsNotPending(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedCampaign, _ := model.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	expectedCampaign.Status = model.Started

	mockedRepository.On("GetById", mock.Anything).Return(expectedCampaign, nil)

	err := service.Delete("idtest")

	assert.NotNil(err)
	assert.Equal(errors.New("campaign could not be deleted, because is not pending").Error(), err.Error())
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErrorWhenDeletingCampaign(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedCampaign, _ := model.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	mockedRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == expectedCampaign.ID
	})).Return(expectedCampaign, nil)

	mockedRepository.On("Delete", mock.MatchedBy(func(campaign *model.Campaign) bool {
		if campaign.Name != expectedCampaign.Name || campaign.Content != expectedCampaign.Content || campaign.ID != expectedCampaign.ID || len(campaign.Contacts) != len(expectedCampaign.Contacts) {
			return false
		}
		return true
	})).Return(errors.New("unexpected error while deleting"))

	err := service.Delete(expectedCampaign.ID)

	assert.NotNil(err)
	assert.Equal(internalerrors.NewErrInternal().Error(), err.Error())
	mockedRepository.AssertExpectations(t)
}

func TestShouldStartCampaign(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedCampaign, _ := model.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	mockedRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == expectedCampaign.ID
	})).Return(expectedCampaign, nil)

	mockedRepository.On("Update", mock.MatchedBy(func(campaign *model.Campaign) bool {
		if campaign.Name != expectedCampaign.Name || campaign.Content != expectedCampaign.Content || campaign.ID != expectedCampaign.ID || len(campaign.Contacts) != len(expectedCampaign.Contacts) {
			return false
		}
		return true
	})).Return(nil)

	mailIsSent := false
	sendMail := func(campaign *model.Campaign) error {
		if campaign.ID == expectedCampaign.ID {
			mailIsSent = true
		}
		return nil
	}
	service.SendMail = sendMail

	err := service.Start(expectedCampaign.ID)

	assert.Nil(err)
	assert.True(mailIsSent)
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErrorWhileSendingMail(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedCampaign, _ := model.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	mockedRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == expectedCampaign.ID
	})).Return(expectedCampaign, nil)

	mailIsSent := false
	sendMail := func(campaign *model.Campaign) error {
		return errors.New("error while sending mail")
	}
	service.SendMail = sendMail

	err := service.Start(expectedCampaign.ID)

	assert.NotNil(err)
	assert.Equal(internalerrors.NewErrInternal().Error(), err.Error())
	assert.False(mailIsSent)
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErrorWhenStartingCampaignById(t *testing.T) {
	assert := assert.New(t)
	setup()

	mockedRepository.On("GetById", mock.Anything).Return(nil, errors.New("unexpected error getting by id"))

	err := service.Start("idtest")

	assert.NotNil(err)
	assert.Equal(internalerrors.NewErrInternal().Error(), err.Error())
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnCampaignNotFoundErrorWhenStarting(t *testing.T) {
	assert := assert.New(t)
	setup()

	mockedRepository.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Start("idtest")

	assert.NotNil(err)
	assert.Equal(internalerrors.NewErrCampaignNotFound().Error(), err.Error())
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErrorWhenStatusIsNotPendingWhileStarting(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedCampaign, _ := model.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	expectedCampaign.Status = model.Started

	mockedRepository.On("GetById", mock.Anything).Return(expectedCampaign, nil)

	err := service.Start("idtest")

	assert.NotNil(err)
	assert.Equal(errors.New("campaign could not be started, because is not pending").Error(), err.Error())
	mockedRepository.AssertExpectations(t)
}

func TestShouldReturnErrorWhenStartingCampaign(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedCampaign, _ := model.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	mockedRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == expectedCampaign.ID
	})).Return(expectedCampaign, nil)

	mockedRepository.On("Update", mock.MatchedBy(func(campaign *model.Campaign) bool {
		if campaign.Name != expectedCampaign.Name || campaign.Content != expectedCampaign.Content || campaign.ID != expectedCampaign.ID || len(campaign.Contacts) != len(expectedCampaign.Contacts) {
			return false
		}
		return true
	})).Return(errors.New("unexpected error while starting"))

	mailIsSent := false
	sendMail := func(campaign *model.Campaign) error {
		if campaign.ID == expectedCampaign.ID {
			mailIsSent = true
		}
		return nil
	}
	service.SendMail = sendMail
	err := service.Start(expectedCampaign.ID)

	assert.NotNil(err)
	assert.Equal(internalerrors.NewErrInternal().Error(), err.Error())
	assert.True(mailIsSent)
	mockedRepository.AssertExpectations(t)
}
