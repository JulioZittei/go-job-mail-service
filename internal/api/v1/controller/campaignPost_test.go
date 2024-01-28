package controller

import (
	"strings"
	"testing"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldCreateAndSaveACampaign(t *testing.T) {
	assert := assert.New(t)
	setup()

	bodyExpected := map[string]string{
		"id": "idtest",
	}

	createdByExpected := "teste@teste.com"

	requestBody := contract.NewCampaignInput{
		Name:    "Teste",
		Content: "Content Test",
		Emails: []string{
			"john@mail.com",
		},
	}

	serviceMocked.On("Create", mock.MatchedBy(func(campaign *contract.NewCampaignInput) bool {
		if campaign.Name != requestBody.Name || campaign.Content != requestBody.Content || len(campaign.Emails) != len(requestBody.Emails) || campaign.CreatedBy != createdByExpected {
			return false
		}
		return true
	})).Return("idtest", nil)

	controller := CampaignController{
		CampaignService: serviceMocked,
	}

	req, res := newReqAndRecord("POST", "/campaign", requestBody)

	req = addParamToContext(req, EMAIL_KEY, createdByExpected)

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(err)
	assert.Equal(201, status)
	assert.Equal(bodyExpected, json)
}

func TestShouldReturnErrorWhenSaveACampaign(t *testing.T) {
	assert := assert.New(t)
	setup()

	createdByExpected := "teste@teste.com"

	requestBody := contract.NewCampaignInput{
		Name:    "Teste",
		Content: "Content Test",
		Emails: []string{
			"john@mail.com",
		},
	}

	serviceMocked.On("Create", mock.Anything).Return("", internalerrors.NewErrInternal())

	req, res := newReqAndRecord("POST", "/campaign", requestBody)

	req = addParamToContext(req, EMAIL_KEY, createdByExpected)

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(json)
	assert.Equal(500, status)
	assert.Equal(internalerrors.NewErrInternal().Error(), err.Error())
}

func TestShouldReturnErrorWhenSendMalFormedBody(t *testing.T) {
	assert := assert.New(t)
	setup()

	requestBody := strings.NewReader(`{
		"name": "Campanha de teste",
		"content": "Conteúdo da campanha"
		"emails": [
			"john@mail.com"
		]
	}`)

	req, res := newReqAndRecord("POST", "/campaign", requestBody)

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(json)
	assert.Equal(400, status)
	assert.Equal(internalerrors.NewErrBadRequest().Error(), err.Error())
}
