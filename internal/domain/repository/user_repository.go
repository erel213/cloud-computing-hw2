package repository

import (
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(*entity.User) (entity.User, appError.AppError)
	CheckIfUserExists(userId uuid.UUID) (bool, appError.AppError)
	BlockUser(userId uuid.UUID, blockedUserId uuid.UUID) appError.AppError
	GetUserById(userId uuid.UUID) (*entity.User, appError.AppError)
}
