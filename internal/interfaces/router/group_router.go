package router

import (
	"whatsapp-like/contracts"
	"whatsapp-like/internal/application"

	"github.com/gofiber/fiber/v2"
)

type GroupRouter struct {
	GroupService *application.GroupService
}

func NewGroupRouter(groupService *application.GroupService) *GroupRouter {
	return &GroupRouter{
		GroupService: groupService,
	}
}

func (router *GroupRouter) CreateGroup(c *fiber.Ctx) error {
	groupContract := new(contracts.CreateGroupRequest)
	if err := c.BodyParser(groupContract); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	groupResponse, err := router.GroupService.CreateNewGroup(*groupContract)
	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(groupResponse)
}
