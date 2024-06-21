package router

import (
	"whatsapp-like/contracts"
	"whatsapp-like/internal/application"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (router *GroupRouter) AddUserToGroup(c *fiber.Ctx) error {
	addUserToGroupContract := new(contracts.AddUserToGroupRequest)
	if err := c.BodyParser(addUserToGroupContract); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := router.GroupService.AddUserToGroup(*addUserToGroupContract)
	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusCreated)
	return nil
}

func (router *GroupRouter) RemoveUserFromGroup(c *fiber.Ctx) error {
	removeUserFromGroupContract := new(contracts.RemoveUserFromGroupRequest)
	userId := c.Params("userID")
	groupId := c.Params("groupID")

	removeUserFromGroupContract.UserId, _ = uuid.Parse(userId)
	removeUserFromGroupContract.GroupId, _ = uuid.Parse(groupId)

	err := router.GroupService.RemoveUserFromGroup(*removeUserFromGroupContract)
	if err != nil {
		return c.Status(err.Code()).JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusOK)
	return nil
}
