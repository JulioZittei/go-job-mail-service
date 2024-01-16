package service

import (
	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/repository"
)

type CampaignService struct {
	Repository repository.Repository
}

func (s *CampaignService) Create(campaignInput *contract.NewCampaignInput) (id string, err error) {
	campaign, err := model.NewCampaign(campaignInput.Name, campaignInput.Content, campaignInput.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.NewErrInternal()
	}
	
	return campaign.ID, err
}