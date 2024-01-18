package database

import (
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
)

type CampaignRepository struct {
	campaigns []model.Campaign
}


func (cr *CampaignRepository) Save(campaign *model.Campaign) error {
	cr.campaigns = append(cr.campaigns, *campaign)
	return nil
}

func (cr *CampaignRepository) Get() ([]model.Campaign, error) {
	if len(cr.campaigns) == 0 {
		return []model.Campaign{}, nil
	}
	return cr.campaigns, nil
}

func (cr *CampaignRepository) GetById(id string) (*model.Campaign, error) {
	for _, v := range cr.campaigns {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, internalerrors.NewErrCampaignNotFound()
}