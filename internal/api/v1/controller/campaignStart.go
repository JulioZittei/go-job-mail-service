package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (c *CampaignController) CampaignStart(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	err := c.CampaignService.Start(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return nil, http.StatusAccepted, nil
}
