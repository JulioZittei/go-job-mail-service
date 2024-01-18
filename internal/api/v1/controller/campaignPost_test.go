package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	mockTests "github.com/JulioZittei/go-job-mail-service/internal/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestShouldCreateAndSaveACampaign(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(mockTests.CampaignServiceMock)

	bodyExpected := map[string]string{
		"id": "idtest",
	}

	requestBody := contract.NewCampaignInput{
		Name: "Teste",
		Content: "Content Test",
		Emails: []string{
			"john@mail.com",
		},
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(requestBody)

	serviceMocked.On("Create", mock.MatchedBy(func(campaign *contract.NewCampaignInput) bool {
		if campaign.Name != requestBody.Name || campaign.Content != requestBody.Content || len(campaign.Emails) != len(requestBody.Emails) {
			return false
		}
		return true
	})).Return("idtest", nil)

	controller := CampaignController{
		CampaignService: serviceMocked,
	}

	req, _ := http.NewRequest("POST", "/campaign", &buf)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(err)
	assert.Equal(201, status)
	assert.Equal(bodyExpected, json)
}

func TestShouldReturnErrorWhenSaveACampaign(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(mockTests.CampaignServiceMock)

	requestBody := contract.NewCampaignInput{
		Name: "Teste",
		Content: "Content Test",
		Emails: []string{
			"john@mail.com",
		},
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(requestBody)

	serviceMocked.On("Create", mock.Anything).Return("", internalerrors.NewErrInternal())

	controller := CampaignController{
		CampaignService: serviceMocked,
	}

	req, _ := http.NewRequest("POST", "/campaign", &buf)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(json)
	assert.Equal(500, status)
	assert.Equal(internalerrors.NewErrInternal().Error(), err.Error())
}

func TestShouldReturnErrorWhenSendMalFormedBody(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(mockTests.CampaignServiceMock)
	
	requestBody  := strings.NewReader(`{
		"name": "Campanha de teste",
		"content": "Conteúdo da campanha"
		"emails": [
			"john@mail.com"
		]
	}`)

	controller := CampaignController{
		CampaignService: serviceMocked,
	}

	req, _ := http.NewRequest("POST", "/campaign", requestBody)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(json)
	assert.Equal(400, status)
	assert.Equal(internalerrors.NewErrBadRequest().Error(), err.Error())
}