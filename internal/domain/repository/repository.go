package repository

import "github.com/JulioZittei/go-job-mail-service/internal/domain/model"

type Repository interface {
	Save(campaign *model.Campaign) error
}
