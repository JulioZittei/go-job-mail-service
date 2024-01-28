package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	mockTests "github.com/JulioZittei/go-job-mail-service/internal/test/mocks"
	"github.com/go-chi/chi/v5"
)

var (
	serviceMocked *mockTests.CampaignServiceMock
	controller    = CampaignController{}
)

func setup() {
	serviceMocked = new(mockTests.CampaignServiceMock)
	controller.CampaignService = serviceMocked
}

func newReqAndRecord(method string, url string, requestBody interface{}) (*http.Request, *httptest.ResponseRecorder) {

	var buf bytes.Buffer
	var req *http.Request
	if _, ok := requestBody.(io.Reader); !ok {
		json.NewEncoder(&buf).Encode(requestBody)
		req, _ = http.NewRequest(method, url, &buf)
	} else {
		req, _ = http.NewRequest(method, url, requestBody.(io.Reader))
	}

	res := httptest.NewRecorder()

	return req, res
}

func addParameter(req *http.Request, urlParamKey string, urlParamValue string) *http.Request {
	chiContext := chi.NewRouteContext()

	chiContext.URLParams.Add(urlParamKey, urlParamValue)

	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
}

func addParamToContext(req *http.Request, key interface{}, value interface{}) *http.Request {
	ctx := context.WithValue(req.Context(), key, value)
	return req.WithContext(ctx)
}
