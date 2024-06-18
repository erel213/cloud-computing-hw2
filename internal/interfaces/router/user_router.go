package router

import (
	"cmd/main.go/internal/application"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	userService *application.UserService
}

func NewUserRouter(userService *application.UserService) *UserRouter {
	return &UserRouter{
		userService: userService,
	}
}

func (router *UserRouter) CreateUser(c *fiber.Ctx) error {
	user, err := router.userService.CreateUser()

	if err != nil {
		return c.Status(err.Code()).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
