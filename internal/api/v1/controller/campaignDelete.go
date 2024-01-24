package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (c *CampaignController) CampaignDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	err := c.CampaignService.Delete(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return nil, http.StatusNoContent, nil
}
