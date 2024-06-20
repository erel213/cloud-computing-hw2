package router

import (
	"cmd/main.go/contracts"
	"cmd/main.go/internal/application"

	"github.com/gofiber/fiber/v2"
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
