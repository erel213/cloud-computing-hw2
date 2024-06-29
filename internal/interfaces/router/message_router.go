package router

import (
	"log/slog"
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
		slog.Error("Error parsing request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	sendMessageErr := mr.messageService.SendMessage(*messageContract)
	if sendMessageErr != nil {
		slog.Error("Error sending message: %v", sendMessageErr)
		return c.Status(sendMessageErr.Code()).JSON(sendMessageErr)
	}

	slog.Info("Message sent:", "fromUser", messageContract.FromUser)
	return c.SendStatus(fiber.StatusCreated)
}

func (mr *MessageRouter) GetMessagesForUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if userId == "" {
		slog.Error("userId is required")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userId is required"})
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		slog.Error("userId is not valid")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userId is not valid"})
	}

	messages, getMessagesErr := mr.messageService.GetMessagesForUser(parsedUserId)
	if getMessagesErr != nil {
		slog.Error("Error getting messages for user: %v", getMessagesErr)
		return c.Status(getMessagesErr.Code()).JSON(getMessagesErr.Error())
	}

	slog.Info("Messages retrieved for user:", "userId", userId)
	return c.Status(fiber.StatusOK).JSON(messages)
}
