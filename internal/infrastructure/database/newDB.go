package database

import (
	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "host=localhost user=mail_service_job_user password=mail_service_job_pass dbname=mail_job_service port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to database")
	}

	db.AutoMigrate(&model.Campaign{}, &model.Contact{})

	return db
}
