package controller

import (
	"net/http"
	"testing"

	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldDeleteACampaignById(t *testing.T) {
	assert := assert.New(t)
	setup()

	serviceMocked.On("Delete", mock.Anything).Return(nil)

	req, res := newReqAndRecord("Delete", "/campaign/idtest", nil)

	json, status, err := controller.CampaignDelete(res, req)

	assert.Nil(err)
	assert.Nil(json)
	assert.Equal(http.StatusNoContent, status)
}

func TestShouldReturnErrorWhenDeletingACampaign(t *testing.T) {
	assert := assert.New(t)
	setup()

	expectedError := internalerrors.NewErrCampaignNotFound()

	serviceMocked.On("Delete", mock.Anything).Return(expectedError)

	req, res := newReqAndRecord("Delete", "/campaign/idtest", nil)

	json, status, err := controller.CampaignDelete(res, req)

	assert.Nil(json)
	assert.NotNil(err)
	assert.Equal(http.StatusInternalServerError, status)
	assert.Equal(expectedError.Error(), err.Error())
}
