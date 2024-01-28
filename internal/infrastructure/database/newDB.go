package database

import (
	"os"

	"github.com/JulioZittei/go-job-mail-service/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to database")
	}

	db.AutoMigrate(&model.Campaign{}, &model.Contact{})

	return db
}
