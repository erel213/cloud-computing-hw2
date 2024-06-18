package repository

import (
	"cmd/main.go/internal/appError"
	"cmd/main.go/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(*entity.User) (entity.User, appError.AppError)
}
