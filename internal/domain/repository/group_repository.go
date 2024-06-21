package repository

import (
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
)

type GroupRepository interface {
	CreateGroup(group *entity.Group) appError.AppError
}
