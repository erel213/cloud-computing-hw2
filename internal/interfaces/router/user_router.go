package router

import (
	"log/slog"
	"whatsapp-like/contracts"
	"whatsapp-like/internal/application"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		slog.Error("Error creating user: %v", err)
		return c.Status(err.Code()).JSON(err.Error())
	}

	slog.Info("User created:", "userId", user.UserId)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (router *UserRouter) BlockUser(c *fiber.Ctx) error {
	request := new(contracts.BlockUserRequest)
	parseErr := c.BodyParser(request)
	if parseErr != nil {
		slog.Error("Error parsing request: %v", parseErr)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": parseErr.Error()})
	}

	user, err := router.userService.BlockUser(request.UserId, request.BlockedUserId)

	if err != nil {
		slog.Error("Error blocking user: %v", err)
		return c.Status(err.Code()).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Info("User blocked:", "userId", request.UserId, "blockedUserId", request.BlockedUserId)
	return c.Status(fiber.StatusOK).JSON(user)
}

func (router *UserRouter) GetUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")
	userIdParsed, parseErr := uuid.Parse(userId)

	if parseErr != nil {
		slog.Error("Error parsing userId: %v", parseErr)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"errors": parseErr.Error()})
	}

	response, responseErr := router.userService.GetUserById(userIdParsed)
	if responseErr != nil {
		slog.Error("Error getting user by id: %v", responseErr)
		return c.Status(responseErr.Code()).JSON(fiber.Map{"errors": responseErr.Error()})
	}

	slog.Info("User retrieved:", "userId", userId)
	return c.Status(fiber.StatusOK).JSON(response)
}
