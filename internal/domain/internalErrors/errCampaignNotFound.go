package internalerrors

import (
	"net/http"
	"strings"

	"github.com/JulioZittei/go-job-mail-service/internal/locale/message"
)

type ErrCampaignNotFound struct {
	StatusCode int
	StatusText string
	Title string
	Detail string
}

func NewErrCampaignNotFound() *ErrBadRequest {
	statusCode := http.StatusNotFound
	statusText := strings.ToUpper(http.StatusText(statusCode))
	statusText = strings.Replace(statusText, " ", "_", -1)
	title, _ := message.GetMessage(statusText)

	return &ErrBadRequest{
		StatusCode: statusCode,
		StatusText: statusText,
		Title: title,
		Detail: "Campaign not found.",
	}
}

func (ev *ErrCampaignNotFound) Error() string {
	return "campaign not found"
}