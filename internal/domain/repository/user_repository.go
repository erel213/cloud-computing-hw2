package repository

import (
	"cmd/main.go/internal/appError"
	"cmd/main.go/internal/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(*entity.User) (entity.User, appError.AppError)
	CheckIfUserExists(userId uuid.UUID) (bool, appError.AppError)
}
