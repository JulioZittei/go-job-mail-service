package controller

import (
	"net/http"
	"testing"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldStartCampaignById(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedCampaign := &contract.CampaignOutput{
		ID:      "idtest",
		Name:    "Teste",
		Content: "Content Test",
		Status:  model.Pending,
	}

	serviceMocked.On("GetById", mock.Anything).Return(expectedCampaign, nil)

	req, res := newReqAndRecord("GET", "/campaign/idtest", nil)

	json, status, err := controller.CampaignGetById(res, req)

	assert.Nil(err)
	assert.NotNil(json)
	assert.Equal(200, status)
	assert.Equal(expectedCampaign.ID, json.(*contract.CampaignOutput).ID)
	assert.Equal(expectedCampaign.Name, json.(*contract.CampaignOutput).Name)
}

func TestShouldReturnErrorWhenGetCampaignById(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedError := internalerrors.NewErrInternal()

	serviceMocked.On("GetById", mock.Anything).Return(nil, expectedError)

	req, res := newReqAndRecord("GET", "/campaign/idtest", nil)

	json, status, err := controller.CampaignGetById(res, req)

	assert.Nil(json)
	assert.NotNil(err)
	assert.Equal(500, status)
	assert.Equal(expectedError.Error(), err.Error())
}

func TestShouldReturnCampaignNotFoundErrorWhemGettingById(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedError := internalerrors.NewErrCampaignNotFound()

	serviceMocked.On("GetById", mock.Anything).Return(nil, expectedError)

	req, res := newReqAndRecord("GET", "/campaign/idtest", nil)

	json, status, err := controller.CampaignGetById(res, req)

	assert.Nil(json)
	assert.NotNil(err)
	assert.Equal(http.StatusInternalServerError, status)
	assert.Equal(expectedError.Error(), err.Error())
}
