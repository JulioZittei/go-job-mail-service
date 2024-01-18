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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) Create(campaignInput *contract.NewCampaignInput) (id string, err error) {
	args := s.Called(campaignInput)
	return args.String(0), args.Error(1)
}

func (s *serviceMock) GetById(id string) (*contract.CampaignOutput, error) {
	args := s.Called(id)
	return &contract.CampaignOutput{}, args.Error(1)
}

var (
	controller = CampaignController{}
)

func TestShouldCreateAndSaveACampaign(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(serviceMock)

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

	controller.CampaignService = serviceMocked

	req, _ := http.NewRequest("POST", "/campaign", &buf)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(err)
	assert.Equal(201, status)
	assert.Equal(bodyExpected, json)
}

func TestShouldReturnErrorWhenSomethingWrong(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(serviceMock)

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

	controller.CampaignService = serviceMocked

	req, _ := http.NewRequest("POST", "/campaign", &buf)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(json)
	assert.Equal(500, status)
	assert.Equal(internalerrors.NewErrInternal().Error(), err.Error())
}

func TestShouldReturnErrorWhenSendMalFormedBody(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(serviceMock)
	
	requestBody  := strings.NewReader(`{
		"name": "Campanha de teste",
		"content": "Conteúdo da campanha"
		"emails": [
			"john@mail.com"
		]
	}`)

	controller.CampaignService = serviceMocked

	req, _ := http.NewRequest("POST", "/campaign", requestBody)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(json)
	assert.Equal(400, status)
	assert.Equal(internalerrors.NewErrBadRequest().Error(), err.Error())
}

func TestShouldGetListOfCampaigns(t *testing.T) {
	assert := assert.New(t)
	serviceMocked := new(serviceMock)

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

	controller.CampaignService = serviceMocked

	req, _ := http.NewRequest("POST", "/campaign", &buf)
	res := httptest.NewRecorder()

	json, status, err := controller.CampaignPost(res, req)

	assert.Nil(err)
	assert.Equal(201, status)
	assert.Equal(bodyExpected, json)
}