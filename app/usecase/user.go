package usecase

import (
	"go-ci/domain/entity"
	irepository "go-ci/domain/repository"
	"go-ci/infra/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserUseCase interface {
	Register(username, password string) (entity.User, error)
	GetByUsername(username string) *entity.User
	Login(username, password string) (*string, bool)
}

type UserUseCaseImpl struct {
	userRepo irepository.IUserRepository
	config   *config.Config
}

func NewUserUseCase(userRepo irepository.IUserRepository, config *config.Config) UserUseCase {
	return &UserUseCaseImpl{userRepo, config}
}

func (uc *UserUseCaseImpl) Register(username, password string) (entity.User, error) {
	return uc.userRepo.Register(username, password)
}

func (uc *UserUseCaseImpl) GetByUsername(username string) *entity.User {
	return uc.userRepo.GetByUsername(username)
}

func (uc *UserUseCaseImpl) Login(username, password string) (*string, bool) {
	user := uc.userRepo.GetByUsername(username)

	if user == nil {
		return nil, false
	}

	isValid, _ := uc.userRepo.CheckPassword(user, password)
	if !isValid {
		return nil, false
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			ID:        strconv.FormatUint(uint64(user.ID), 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.config.JWTExpire)),
		},
	)

	tokenString, err := token.SignedString([]byte(uc.config.JWTSecret))

	if err != nil {
		return nil, false
	}

	return &tokenString, true
}
