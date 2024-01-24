package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	mockTests "github.com/JulioZittei/go-job-mail-service/internal/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldDeleteACampaignById(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(mockTests.CampaignServiceMock)

	serviceMocked.On("Delete", mock.Anything).Return(nil)

	controller := CampaignController{
		CampaignService: serviceMocked,
	}

	req, _ := http.NewRequest("Delete", "/campaign/idtest", nil)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignDelete(res, req)

	assert.Nil(err)
	assert.Nil(json)
	assert.Equal(http.StatusNoContent, status)
}

func TestShouldReturnErrorWhenDeletingACampaign(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(mockTests.CampaignServiceMock)

	expectedError := internalerrors.NewErrCampaignNotFound()

	serviceMocked.On("Delete", mock.Anything).Return(expectedError)

	controller := CampaignController{
		CampaignService: serviceMocked,
	}

	req, _ := http.NewRequest("Delete", "/campaign/idtest", nil)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignDelete(res, req)

	assert.Nil(json)
	assert.NotNil(err)
	assert.Equal(http.StatusInternalServerError, status)
	assert.Equal(expectedError.Error(), err.Error())
}
