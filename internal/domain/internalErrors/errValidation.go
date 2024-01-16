package internalerrors

import (
	"net/http"
	"strings"

	"github.com/JulioZittei/go-job-mail-service/internal/locale/message"
)

type ErrValidation struct {
	StatusCode int
	StatusText string
	Title string
	Detail string
	ErrorsParam []ErrorsParam
}

func NewErrValidation(errorsParam []ErrorsParam) *ErrValidation {
	statusCode := http.StatusUnprocessableEntity
	statusText := strings.ToUpper(http.StatusText(statusCode))
	statusText = strings.Replace(statusText, " ", "_", -1)
	title, _ := message.GetMessage(statusText)

	return &ErrValidation{
		StatusCode: statusCode,
		StatusText: statusText,
		Title: title,
		Detail: "",
		ErrorsParam: errorsParam,
	}
}

func (ev *ErrValidation) Error() string {
	return "validation error"
}