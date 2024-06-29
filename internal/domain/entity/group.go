package entity

import (
	"errors"
	"fmt"
	"whatsapp-like/internal/appError"

	"github.com/google/uuid"
)

type Group struct {
	GroupId   uuid.UUID
	CreatedBy uuid.UUID
	GroupName string
	Users     []uuid.UUID
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

func (g *Group) AddUser(userId uuid.UUID) appError.AppError {
	//Check if user in group
	for _, user := range g.Users {
		if user == userId {
			return appError.ValidationError{Err: fmt.Errorf("user %s exists in group %s", userId, g.GroupId)}
		}
	}

	g.Users = append(g.Users, userId)
	return nil
}

func (g *Group) RemoveUser(userId uuid.UUID) appError.AppError {
	for i, user := range g.Users {
		if user == userId {
			g.Users = append(g.Users[:i], g.Users[i+1:]...)
			return nil
		}
	}

	return appError.NotFoundError{Err: fmt.Errorf("user %s not exist in group %s", userId, g.GroupId)}
}

func (g *Group) IsUserInGroup(userId uuid.UUID) (bool, appError.AppError) {
	for _, user := range g.Users {
		if user == userId {
			return true, nil
		}
	}

	return false, nil
}

func validateNewGroup(groupName string) appError.AppError {
	if groupName == "" {
		return appError.ValidationError{Err: errors.New("group name cannot be null")}
	}

	return nil
}
