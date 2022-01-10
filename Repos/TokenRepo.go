package repos

import (
	"gauth.com/Database/Models"
	"gorm.io/gorm"
)

type TokenRepo interface {
	Migrate() error
}

type tokenRepository struct {
	DB *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepo {
	return tokenRepository{
		DB: db,
	}
}

func (t tokenRepository) Migrate() error {
	return t.DB.AutoMigrate(&Models.Token{})
}
