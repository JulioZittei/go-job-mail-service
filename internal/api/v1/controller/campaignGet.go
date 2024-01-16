package controller

import (
	"net/http"
)

func (c *CampaignController) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := c.CampaignService.Repository.Get()
	if err != nil {
		return nil, 500, err
	}

	return map[string]interface{}{
		"campaigns": campaigns,
	}, 200, nil
}