package router

import (
	"log/slog"
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
		slog.Error("Error creating group: %v", err)
		return c.Status(err.Code()).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Info("Group created:", "groupId", groupResponse.GroupId)
	return c.Status(fiber.StatusCreated).JSON(groupResponse)
}

func (router *GroupRouter) AddUserToGroup(c *fiber.Ctx) error {
	addUserToGroupContract := new(contracts.AddUserToGroupRequest)
	if err := c.BodyParser(addUserToGroupContract); err != nil {
		slog.Error("Error parsing request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := router.GroupService.AddUserToGroup(*addUserToGroupContract)
	if err != nil {
		slog.Error("Error adding user to group: %v", err)
		return c.Status(err.Code()).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Info("User added to group:", "userId", addUserToGroupContract.UserId, "groupId", addUserToGroupContract.GroupId)
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
		slog.Error("Error removing user from group: %v", err)
		return c.Status(err.Code()).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Info("User removed from group:", "userId", userId, "groupId", groupId)
	c.Status(fiber.StatusOK)
	return nil
}
