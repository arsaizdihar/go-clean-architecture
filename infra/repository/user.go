package repository

import (
	"errors"
	helpers "go-ci/common"
	"go-ci/domain/entity"
	irepository "go-ci/domain/repository"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) irepository.IUserRepository {
	return &UserRepository{db}
}

type UserRepository struct {
	db *gorm.DB
}

func (ur *UserRepository) Register(username, password string) (entity.User, error) {
	password, err := helpers.HashPassword(password)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Username: username,
		Password: password,
	}

	ur.db.Create(&user)

	return user, nil
}

func (ur *UserRepository) CheckPassword(user *entity.User, password string) (bool, error) {
	return helpers.CheckPassword(user.Password, password)
}

func (ur *UserRepository) GetByUsername(username string) *entity.User {
	var user entity.User

	result := ur.db.First(&user, "username = ?", username)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func (ur *UserRepository) Get(id uint) *entity.User {
	var user entity.User

	result := ur.db.First(&user, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}
