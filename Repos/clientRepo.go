package repos

import (
	"gauth.com/Database/Models"
	"gorm.io/gorm"
)

type ClientRepo interface {
	Migrate() error
}

type clientRepository struct {
	DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepo {
	return clientRepository{
		DB: db,
	}
}

func (c clientRepository) Migrate() error {
	return c.DB.AutoMigrate(&Models.Client{})
}
