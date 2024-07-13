package middleware

import (
	"go-ci/domain/entity"
	irepository "go-ci/domain/repository"
	"go-ci/infra/config"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	config   *config.Config
	userRepo irepository.IUserRepository
}

func NewAuthMiddleware(config *config.Config, userRepo irepository.IUserRepository) *AuthMiddleware {
	return &AuthMiddleware{config, userRepo}
}

func (auth *AuthMiddleware) ParseToken(c fiber.Ctx) error {
	tokenString := c.Cookies("token")

	if tokenString == "" {
		return c.Next()
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.config.JWTSecret), nil
	})

	if err != nil {
		return c.Next()
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	id, err := strconv.ParseUint(claims.ID, 10, 64)

	if err != nil {
		return c.Next()
	}

	user := auth.userRepo.Get(uint(id))
	c.Locals("user", user)
	return c.Next()
}

func (auth *AuthMiddleware) AuthRequired(c fiber.Ctx) error {
	if c.Locals("user") == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}

func (auth *AuthMiddleware) GetUser(c fiber.Ctx) *entity.User {
	return c.Locals("user").(*entity.User)
}
