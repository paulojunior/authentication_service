package repository

import (
	"authentication/data"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (u *UserRepository) FindByEmail(email string) (user data.User, err error) {
	result := u.DB.Where("email = ?", email).First(&user)
	if result != nil && result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
