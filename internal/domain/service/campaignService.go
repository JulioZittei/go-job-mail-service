package service

import (
	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/repository"
)

type CampaignService interface {
	Create(campaignInput *contract.NewCampaignInput) (id string, err error)
	GetById(id string) (*contract.CampaignOutput, error)
}

type CampaignServiceImpl struct {
	Repository repository.Repository
}

func (s *CampaignServiceImpl) Create(campaignInput *contract.NewCampaignInput) (id string, err error) {
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

func (s *CampaignServiceImpl) GetById(id string) (*contract.CampaignOutput, error) {
	campaign, err := s.Repository.GetById(id)
	if err != nil {
		return nil, internalerrors.NewErrInternal()
	}
	return &contract.CampaignOutput{
		ID: campaign.ID,
		Name: campaign.Name,
		Content: campaign.Content,
		Status: campaign.Status,
	}, nil
}