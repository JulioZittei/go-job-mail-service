package service

import (
	"errors"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/contract"
	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"github.com/JulioZittei/go-job-mail-service/internal/domain/repository"
	"gorm.io/gorm"
)

type CampaignService interface {
	Create(campaignInput *contract.NewCampaignInput) (id string, err error)
	GetById(id string) (*contract.CampaignOutput, error)
	Delete(id string) error
}

type CampaignServiceImpl struct {
	Repository repository.Repository
}

func (s *CampaignServiceImpl) Create(campaignInput *contract.NewCampaignInput) (id string, err error) {
	campaign, err := model.NewCampaign(campaignInput.Name, campaignInput.Content, campaignInput.Emails, campaignInput.CreatedBy)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internalerrors.NewErrCampaignNotFound()
		}
		return nil, internalerrors.NewErrInternal()
	}
	return &contract.CampaignOutput{
		ID:           campaign.ID,
		Name:         campaign.Name,
		Content:      campaign.Content,
		Status:       campaign.Status,
		EmailsToSend: len(campaign.Contacts),
		CreatedBy:    campaign.CreatedBy,
	}, nil
}

func (s *CampaignServiceImpl) Delete(id string) error {
	campaign, err := s.Repository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return internalerrors.NewErrCampaignNotFound()
		}
		return internalerrors.NewErrInternal()
	}

	if campaign.Status != model.Pending {
		return errors.New("campaign could not be deleted, because is not pending")
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)
	if err != nil {
		return internalerrors.NewErrInternal()
	}

	return nil
}
