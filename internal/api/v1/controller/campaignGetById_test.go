package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	mockTests "github.com/JulioZittei/go-job-mail-service/internal/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldACampaignById(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(mockTests.CampaignServiceMock)

	expectedCampaign := &contract.CampaignOutput{
		ID: "idtest",
		Name: "Teste",
		Content: "Content Test",
		Status: model.Pending,
	}

	serviceMocked.On("GetById", mock.Anything).Return(expectedCampaign, nil)

	controller := CampaignController{
		CampaignService: serviceMocked,
	}

	req, _ := http.NewRequest("GET", "/campaign/idtest", nil)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignGetById(res, req)

	assert.Nil(err)
	assert.NotNil(json)
	assert.Equal(200, status)
	assert.Equal(expectedCampaign.ID, json.(*contract.CampaignOutput).ID)
	assert.Equal(expectedCampaign.Name, json.(*contract.CampaignOutput).Name)
}

func TestShouldReturnErrorWhenGetCampaignById(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(mockTests.CampaignServiceMock)

	expectedError := internalerrors.NewErrInternal()

	serviceMocked.On("GetById", mock.Anything).Return(nil, expectedError)

	controller := CampaignController{
		CampaignService: serviceMocked,
	}

	req, _ := http.NewRequest("GET", "/campaign/idtest", nil)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignGetById(res, req)

	assert.Nil(json)
	assert.NotNil(err)
	assert.Equal(500, status)
	assert.Equal(expectedError.Error(), err.Error())
}
