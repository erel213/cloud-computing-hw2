package router

import (
	"whatsapp-like/contracts"
	"whatsapp-like/internal/application"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MessageRouter struct {
	messageService *application.MessageService
}

func NewMessageRouter(messageService *application.MessageService) *MessageRouter {
	return &MessageRouter{messageService}
}

func (mr *MessageRouter) SendMessage(c *fiber.Ctx) error {
	messageContract := new(contracts.SendMessageRequest)
	if err := c.BodyParser(messageContract); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := mr.messageService.SendMessage(*messageContract)
	if err != nil {
		return c.Status(err.Code()).JSON(err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (mr *MessageRouter) GetMessagesForUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userId is required"})
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userId is not valid"})
	}

	messages, getMessagesErr := mr.messageService.GetMessagesForUser(parsedUserId)
	if err != nil {
		return c.Status(getMessagesErr.Code()).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(messages)
}
