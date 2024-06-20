package repository

import (
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(*entity.User) (entity.User, appError.AppError)
	CheckIfUserExists(userId uuid.UUID) (bool, appError.AppError)
}
