package repository

import (
	"cmd/main.go/internal/appError"
	"cmd/main.go/internal/domain/entity"
)

type MessageRepository interface {
	CreateMessage(message *entity.Message) appError.AppError
}
