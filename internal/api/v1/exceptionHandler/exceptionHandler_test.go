package exceptionhandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/stretchr/testify/assert"
)
func TestShouldHandlerResponseWhenCreateCampaign(t *testing.T) {
	assert := assert.New(t)

	type bodyTest struct {
		ID string `json:"id"`
	}

	objectExpected := bodyTest{
		ID: "idtest",
	}

	controller := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error ) {
		return objectExpected, 201, nil
	}

	handler := ExceptionHandler(controller)

	req, _ := http.NewRequest("GET", "/campaign", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	objectReturned := bodyTest{}
	json.Unmarshal(res.Body.Bytes(), &objectReturned)

	assert.Equal(http.StatusCreated, res.Code)
	assert.Equal(objectExpected, objectReturned)
}

func TestShouldHandlerErrorWhenControllerReturnsInternalError(t *testing.T) {
	assert := assert.New(t)
	controller := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error ) {
		return nil, 0, internalerrors.NewErrInternal()
	}

	handler := ExceptionHandler(controller)

	req, _ := http.NewRequest("GET", "/campaign", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.NewErrInternal().Title)
}

func TestShouldHandlerErrorWhenControllerReturnsValidationError(t *testing.T) {
	assert := assert.New(t)
	var errorsParams = make([]internalerrors.ErrorsParam, 1)
	
	errorsParams[0] = internalerrors.ErrorsParam{
			Param: "name",
			Message: "must be at least 5.",
	}
	

	controller := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error ) {
		return nil, 0, internalerrors.NewErrValidation(errorsParams)
	}

	handler := ExceptionHandler(controller)

	req, _ := http.NewRequest("POST", "/campaign", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	assert.Equal(http.StatusUnprocessableEntity, res.Code)
	assert.Contains(res.Body.String(), internalerrors.NewErrValidation(errorsParams).Title)
}

func TestShouldHandlerErrorWhenControllerReturnsBadRequestError(t *testing.T) {
	assert := assert.New(t)	

	controller := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error ) {
		return nil, 0, internalerrors.NewErrBadRequest()
	}

	handler := ExceptionHandler(controller)

	req, _ := http.NewRequest("POST", "/campaign", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), internalerrors.NewErrBadRequest().Title)
}

func TestShouldHandlerErrorWhenControllerReturnsUnmappedError(t *testing.T) {
	assert := assert.New(t)	

	controller := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error ) {
		return nil, 0, errors.New("unmapped error")
	}

	handler := ExceptionHandler(controller)

	req, _ := http.NewRequest("POST", "/campaign", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), "unmaped error")
}