package entity

import (
	"cmd/main.go/internal/appError"
	"errors"

	"github.com/google/uuid"
)

type User struct {
	UserId uuid.UUID
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
