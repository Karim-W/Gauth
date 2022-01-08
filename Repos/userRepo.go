package repos

import (
	"errors"
	"log"

	"gauth.com/dbmanager/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Save(models.User) (models.User, error)
	GetAll() ([]models.User, error)
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
	WithTrx(*gorm.DB) userRepository
	Migrate() error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&models.User{})
}

func (u userRepository) Save(user models.User) (models.User, error) {
	log.Print("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	return user, err

}

func (u userRepository) GetAll() (users []models.User, err error) {
	log.Print("[UserRepository]...Get All")
	err = u.DB.Find(&users).Error
	return users, err

}

func (u userRepository) WithTrx(trxHandle *gorm.DB) userRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}

func (u userRepository) IncrementMoney(receiver uint, amount float64) error {
	log.Print("[UserRepository]...Increment Money")
	return u.DB.Model(&models.User{}).Where("id=?", receiver).Update("wallet", gorm.Expr("wallet + ?", amount)).Error
}

func (u userRepository) DecrementMoney(giver uint, amount float64) error {
	log.Print("[UserRepository]...Decrement Money")
	return errors.New("something")
	// return u.DB.Model(&model.User{}).Where("id=?", giver).Update("wallet", gorm.Expr("wallet - ?", amount)).Error
}
