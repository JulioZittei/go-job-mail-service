package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldStartACampaignById(t *testing.T) {
	assert := assert.New(t)
	setup()

	serviceMocked.On("Start", mock.Anything).Return(nil)

	controller := CampaignController{
		CampaignService: serviceMocked,
	}

	req, _ := http.NewRequest("Patch", "/campaign/idtest", nil)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignStart(res, req)

	assert.Nil(err)
	assert.Nil(json)
	assert.Equal(http.StatusAccepted, status)
}

func TestShouldReturnErrorWhenStartingACampaign(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedError := internalerrors.NewErrCampaignNotFound()

	campaignId := "idtest"

	serviceMocked.On("Start", mock.MatchedBy(func(id string) bool {
		return id == campaignId
	})).Return(expectedError)

	req, res := newReqAndRecord("PATCH", "campaign/{id}", nil)

	req = addParameter(req, "id", campaignId)

	json, status, err := controller.CampaignStart(res, req)

	assert.Nil(json)
	assert.NotNil(err)
	assert.Equal(http.StatusInternalServerError, status)
	assert.Equal(expectedError.Error(), err.Error())
}
