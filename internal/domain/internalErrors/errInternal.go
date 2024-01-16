package internalerrors

import (
	"net/http"
	"strings"

	"github.com/JulioZittei/go-job-mail-service/internal/locale/message"
)

type ErrInternal struct {
	StatusCode int
	StatusText string
	Title string
	Detail string
}

func (ev *ErrInternal) Error() string {
	return "internal server error"
}

func NewErrInternal() *ErrInternal {
	statusCode := http.StatusInternalServerError
	statusText := strings.ToUpper(http.StatusText(statusCode))
	statusText = strings.Replace(statusText, " ", "_", -1)
	title, _ := message.GetMessage(statusText)

	return &ErrInternal{
		StatusCode: statusCode,
		StatusText: statusText,
		Title: title,
		Detail: "",
	}
}