package irepository

import "go-ci/domain/entity"

type IUserRepository interface {
	Register(username, password string) (entity.User, error)
	GetByUsername(username string) *entity.User
	Get(id uint) *entity.User
	CheckPassword(user *entity.User, password string) (bool, error)
}
