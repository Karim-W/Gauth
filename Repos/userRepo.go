package repos

import (
	"log"

	"gauth.com/Dbmanager/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	GetUser(id string) (models.User, error)
	GetAll() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUserDetails(id string, user models.User) (models.User, error)
	DeleteUser(id string) error
	Migrate() error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) GetUser(id string) (user models.User, err error) {
	log.Print("[UserRepository]...Get User")
	err = u.DB.Where("id=?", id).First(&user).Error
	return user, err
}

func (u userRepository) GetAll() (users []models.User, err error) {
	log.Print("[UserRepository]...Get All")
	err = u.DB.Find(&users).Error
	return users, err

}
func (u userRepository) CreateUser(user models.User) (models.User, error) {
	log.Print("[UserRepository]...Create User")
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) UpdateUserDetails(id string, user models.User) (models.User, error) {
	log.Print("[UserRepository]...Update User")
	err := u.DB.Where("id=?", id).Updates(user).Error
	return user, err
}

func (u userRepository) DeleteUser(id string) error {
	log.Print("[UserRepository]...Delete User")
	err := u.DB.Where("id=?", id).Delete(&models.User{}).Error
	return err
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&models.User{})
}
