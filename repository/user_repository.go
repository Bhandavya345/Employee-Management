package repository

import (
	"errors"

	"github.com/Bhandavya345/Employee-Management/database"
	"github.com/Bhandavya345/Employee-Management/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type UserRepo struct{}

func NewUserRepository() UserRepository {
	return &UserRepo{}
}

// Create User
func (r *UserRepo) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// Get User By Email
func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	err := database.DB.
		Where("email = ?", email).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return &user, err
}

// Get User By ID
func (r *UserRepo) GetUserByID(id uint) (*models.User, error) {

	var user models.User

	err := database.DB.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return &user, err
}
