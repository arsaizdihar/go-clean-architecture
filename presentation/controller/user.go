package controller

import (
	"go-ci/app/usecase"
	"go-ci/dto"
	"go-ci/infra/config"
	"go-ci/presentation/middleware"
	"time"

	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	userUseCase usecase.UserUseCase
	config      *config.Config
	auth        *middleware.AuthMiddleware
}

func NewUserController(userUseCase usecase.UserUseCase, config *config.Config, auth *middleware.AuthMiddleware) *UserController {
	return &UserController{userUseCase, config, auth}
}

func (uc *UserController) Setup() *fiber.App {
	app := fiber.New()

	app.Post("/register", func(c fiber.Ctx) error {
		var dto dto.UserRegisterRequest

		if err := c.Bind().Body(&dto); err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.Send([]byte(err.Error()))
		}

		user, err := uc.userUseCase.Register(dto.Username, dto.Password)
		if err != nil {
			return err
		}
		return c.JSON(user)
	})

	app.Post("/login", func(c fiber.Ctx) error {
		var dto dto.UserRegisterRequest

		if err := c.Bind().Body(&dto); err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.Send([]byte(err.Error()))
		}

		token, valid := uc.userUseCase.Login(dto.Username, dto.Password)
		if !valid {
			c.SendStatus(fiber.StatusBadRequest)
			return c.Send([]byte("Invalid username or password"))
		}
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    *token,
			HTTPOnly: true,
			Expires:  time.Now().Add(uc.config.JWTExpire),
		})

		return c.JSON(fiber.Map{
			"status": "success",
		})
	})

	app.Get("/me", func(c fiber.Ctx) error {
		return c.JSON(uc.auth.GetUser(c))
	}, uc.auth.ParseToken, uc.auth.AuthRequired)

	return app
}
