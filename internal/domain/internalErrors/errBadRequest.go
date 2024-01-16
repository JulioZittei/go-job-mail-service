package internalerrors

import (
	"net/http"
	"strings"

	"github.com/JulioZittei/go-job-mail-service/internal/locale/message"
)

type ErrBadRequest struct {
	StatusCode int
	StatusText string
	Title string
	Detail string
}

func NewErrBadRequest() *ErrBadRequest {
	statusCode := http.StatusBadRequest
	statusText := strings.ToUpper(http.StatusText(statusCode))
	statusText = strings.Replace(statusText, " ", "_", -1)
	title, _ := message.GetMessage(statusText)

	return &ErrBadRequest{
		StatusCode: statusCode,
		StatusText: statusText,
		Title: title,
		Detail: "",
	}
}

func (ev *ErrBadRequest) Error() string {
	return "bad request error"
}