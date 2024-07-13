package controller

import (
	"errors"
	"go-ci/app/usecase"
	derror "go-ci/domain/error"
	"go-ci/dto"
	"go-ci/presentation/middleware"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
	auth           *middleware.AuthMiddleware
}

func NewProductController(usecase usecase.ProductUseCase, auth *middleware.AuthMiddleware) *ProductController {
	return &ProductController{usecase, auth}
}

func (pc *ProductController) Setup() *fiber.App {
	app := fiber.New()
	app.Post("/products", func(c fiber.Ctx) error {
		var dto dto.InsertProductRequest
		if err := c.Bind().Body(&dto); err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.Send([]byte(err.Error()))
		}
		user := pc.auth.GetUser(c)

		if product, err := pc.productUseCase.Insert(dto.Name, dto.Price, user.ID); err != nil {
			if errors.Is(err, derror.ErrUserNotFound) {
				c.SendStatus(fiber.StatusBadRequest)
				return c.Send([]byte("Seller not found"))
			}

			return err
		} else {
			return c.JSON(fiber.Map{
				"product": product,
			})
		}
	}, pc.auth.ParseToken, pc.auth.AuthRequired)

	app.Get("/products", func(c fiber.Ctx) error {
		products, err := pc.productUseCase.GetAll()
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"products": products,
		})
	})

	app.Delete("/products/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.Send([]byte(err.Error()))
		}
		user := pc.auth.GetUser(c)

		if err := pc.productUseCase.Delete(user.ID, uint(id)); err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"status": "success",
		})
	}, pc.auth.ParseToken, pc.auth.AuthRequired)
	return app
}
