package router

import (
	"whatsapp-like/contracts"
	"whatsapp-like/internal/application"

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

func (router *UserRouter) BlockUser(c *fiber.Ctx) error {
	request := new(contracts.BlockUserRequest)
	parseErr := c.BodyParser(request)
	if parseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": parseErr.Error()})
	}

	user, err := router.userService.BlockUser(request.UserId, request.BlockedUserId)

	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
