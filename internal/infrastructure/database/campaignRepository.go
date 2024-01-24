package database

import (
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (cr *CampaignRepository) Save(campaign *model.Campaign) error {
	tx := cr.Db.Create(campaign)
	return tx.Error
}

func (cr *CampaignRepository) Update(campaign *model.Campaign) error {
	tx := cr.Db.Save(campaign)
	return tx.Error
}

func (cr *CampaignRepository) Get() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	tx := cr.Db.Find(&campaigns)
	return campaigns, tx.Error
}

func (cr *CampaignRepository) GetById(id string) (*model.Campaign, error) {
	var campaign model.Campaign
	tx := cr.Db.Preload("Contacts").First(&campaign, "id = ?", id)
	return &campaign, tx.Error
}

func (cr *CampaignRepository) Delete(campaign *model.Campaign) error {

	tx := cr.Db.Select("Contacts").Delete(campaign)
	return tx.Error
}
