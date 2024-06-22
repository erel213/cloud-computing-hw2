package entity

import (
	"errors"
	"fmt"
	"slices"
	"whatsapp-like/internal/appError"

	"github.com/google/uuid"
)

type User struct {
	UserId       uuid.UUID
	BlockedUsers []uuid.UUID
}

func NewUser() (*User, appError.AppError) {
	userId, err := uuid.NewUUID()

	if err != nil {
		return nil, &appError.InternalError{Err: errors.New("failed to generate UUID")}
	}

	return &User{
		UserId: userId,
	}, nil
}

func (u *User) BlockUser(blockedUserId uuid.UUID) appError.AppError {

	if slices.Contains(u.BlockedUsers, blockedUserId) {
		return appError.ValidationError{Err: fmt.Errorf("user %s already block user %s", u.UserId, blockedUserId)}
	}

	u.BlockedUsers = append(u.BlockedUsers, blockedUserId)
	return nil
}

func (u *User) FilterBlockedMessages(messages *[]Message) (*[]Message, appError.AppError) {
	filteredMessage := make([]Message, 0)

	for _, message := range *messages {
		if slices.Contains(u.BlockedUsers, message.FromUser) {
			continue
		}

		filteredMessage = append(filteredMessage, message)
	}

	return &filteredMessage, nil
}
