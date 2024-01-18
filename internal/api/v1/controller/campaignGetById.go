package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (c *CampaignController) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	campaign, err := c.CampaignService.GetById(id)
	if err != nil {
		return nil, 500, err
	}

	return campaign, 200, err
}