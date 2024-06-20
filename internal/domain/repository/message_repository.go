package repository

import (
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
)

type MessageRepository interface {
	CreateMessage(message *entity.Message) appError.AppError
}
