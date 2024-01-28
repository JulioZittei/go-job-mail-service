package controller

import (
	"net/http"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/go-chi/render"
)

func (c *CampaignController) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var requestBody = contract.NewCampaignInput{}
	err := render.DecodeJSON(r.Body, &requestBody)
	if err != nil {
		return nil, http.StatusBadRequest, internalerrors.NewErrBadRequest()
	}

	userEmail := r.Context().Value(EMAIL_KEY).(string)

	requestBody.CreatedBy = userEmail

	id, err := c.CampaignService.Create(&requestBody)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return map[string]string{
		"id": id,
	}, http.StatusCreated, nil
}
