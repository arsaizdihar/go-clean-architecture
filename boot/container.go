package boot

import (
	"go-ci/infra/config"
	repo "go-ci/infra/repository"
	"go-ci/presentation"

	"go-ci/app/usecase"
	"go-ci/presentation/controller"
	"go-ci/presentation/middleware"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/dig"
)

func Setup() {
	c := dig.New()

	println("Setting up container")

	c.Provide(config.Setup)

	/* Repository */
	c.Provide(repo.Setup)
	c.Provide(repo.NewUserRepository)
	c.Provide(repo.NewProductRepository)

	/* Usecase */
	c.Provide(usecase.NewUserUseCase)
	c.Provide(usecase.NewProductUseCase)

	/* Middleware */
	c.Provide(middleware.NewAuthMiddleware)

	/* Controller */
	c.Provide(controller.NewUserController)
	c.Provide(controller.NewProductController)

	err := c.Invoke(func(uc *controller.UserController, pc *controller.ProductController) {
		println("Invoking controller")
		app := fiber.New(fiber.Config{
			StructValidator: presentation.NewStructValidator(),
		})
		app.Use("/api", uc.Setup())
		app.Use("/api", pc.Setup())
		app.Listen("localhost:3000")
		println("Server is running on port 3000")
	})

	if err != nil {
		panic(err)
	}
}
