package entity

import (
	"errors"
	"whatsapp-like/internal/appError"

	"github.com/google/uuid"
)

type Message struct {
	MessageId   uuid.UUID
	FromUser    uuid.UUID
	To          uuid.UUID
	MessageBody string
	ToGroup     bool
}

func NewMessage(fromUser, to uuid.UUID, messageBody string, toGroup bool) (*Message, appError.AppError) {
	messageValidationErr := validateNewMessage(fromUser, to, messageBody)

	if messageValidationErr != nil {
		return nil, messageValidationErr
	}
	message := Message{
		MessageId:   uuid.New(),
		FromUser:    fromUser,
		To:          to,
		MessageBody: messageBody,
		ToGroup:     toGroup,
	}

	return &message, nil
}

func validateNewMessage(fromUser uuid.UUID, to uuid.UUID, messageBody string) appError.AppError {
	if fromUser == uuid.Nil {
		return appError.ValidationError{Err: errors.New("from user cannot be empty")}
	}

	if to == uuid.Nil {
		return appError.ValidationError{Err: errors.New("to cannot be empty")}
	}

	if messageBody == "" {
		return appError.ValidationError{Err: errors.New("message body cannot be empty")}
	}

	return nil
}
