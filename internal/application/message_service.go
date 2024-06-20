package application

import (
	"cmd/main.go/contracts"
	"cmd/main.go/internal/appError"
	"cmd/main.go/internal/domain/entity"
	"cmd/main.go/internal/domain/repository"
	"errors"
)

type MessageService struct {
	messageRepository repository.MessageRepository
	userRepository    repository.UserRepository
}

func NewMessageService(messageRepository repository.MessageRepository, userRepository repository.UserRepository) *MessageService {
	return &MessageService{
		messageRepository: messageRepository,
		userRepository:    userRepository,
	}
}

func (ms *MessageService) SendMessage(request contracts.SendMessageRequest) appError.AppError {
	// Check if from user exists
	fromUserExists, err := ms.userRepository.CheckIfUserExists(request.FromUser)

	if err != nil {
		return err
	}

	if !fromUserExists {
		return appError.ValidationError{Err: errors.New("from user does not exist")}
	}

	//Create the message
	message, createMessageErr := entity.NewMessage(request.FromUser, request.To, request.MessageBody, request.ToGroup)

	if createMessageErr != nil {
		return createMessageErr
	}

	// Save the message
	saveMessageErr := ms.messageRepository.CreateMessage(message)
	if saveMessageErr != nil {
		return saveMessageErr
	}

	return nil
}
