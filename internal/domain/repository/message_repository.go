package repository

import (
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"

	"github.com/google/uuid"
)

type MessageRepository interface {
	CreateMessage(message *entity.Message) appError.AppError
	GetMessagesForUser(userId uuid.UUID) ([]*entity.Message, appError.AppError)
}
