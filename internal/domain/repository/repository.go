package repository

import "github.com/JulioZittei/go-job-mail-service/internal/domain/model"

type Repository interface {
	Save(campaign *model.Campaign) error
	Update(campaign *model.Campaign) error
	Get() ([]model.Campaign, error)
	GetById(id string) (*model.Campaign, error)
	Delete(campaign *model.Campaign) error
}
