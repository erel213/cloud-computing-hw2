package entity

import (
	"errors"
	"whatsapp-like/internal/appError"

	"github.com/google/uuid"
)

type Group struct {
	GroupId   uuid.UUID
	CreatedBy uuid.UUID
	GroupName string
}

func NewGroup(groupName string, createdBy uuid.UUID) (*Group, appError.AppError) {
	validationErr := validateNewGroup(groupName)
	if validationErr != nil {
		return nil, validationErr
	}
	groupId, err := uuid.NewUUID()
	if err != nil {
		return nil, &appError.InternalError{Err: err}
	}

	return &Group{
		GroupId:   groupId,
		GroupName: groupName,
		CreatedBy: createdBy,
	}, nil
}

func validateNewGroup(groupName string) appError.AppError {
	if groupName == "" {
		return appError.ValidationError{Err: errors.New("group name cannot be null")}
	}

	return nil
}
