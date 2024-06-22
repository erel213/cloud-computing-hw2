package repository

import (
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"

	"github.com/google/uuid"
)

type GroupRepository interface {
	CreateGroup(group *entity.Group) appError.AppError
	CheckIfGroupExists(uuid.UUID) (bool, appError.AppError)
	CheckIfUserExistsInGroup(userId uuid.UUID, groupId uuid.UUID) (bool, appError.AppError)
	AddUserToGroup(userId uuid.UUID, groupId uuid.UUID) appError.AppError
	RemoveUserFromGroup(userId uuid.UUID, groupId uuid.UUID) appError.AppError
	GetGroupById(groupId uuid.UUID) (*entity.Group, appError.AppError)
}
