package application

import (
	"errors"
	"fmt"
	"whatsapp-like/contracts"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"

	"github.com/google/uuid"
)

type MessageService struct {
	messageRepository repository.MessageRepository
	userRepository    repository.UserRepository
	groupRepostitory  repository.GroupRepository
}

func NewMessageService(messageRepository repository.MessageRepository, userRepository repository.UserRepository, groupRepository repository.GroupRepository) *MessageService {
	return &MessageService{
		messageRepository: messageRepository,
		userRepository:    userRepository,
		groupRepostitory:  groupRepository,
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

	//Validate To prop
	if request.ToGroup {
		group, userInGroupErr := ms.groupRepostitory.GetGroupById(request.To)
		if userInGroupErr != nil {
			return userInGroupErr
		}

		if group == nil {
			return appError.NotFoundError{Err: fmt.Errorf(fmt.Sprintf("group %s not found", request.To))}
		}

		//Check if from user is in group
		isUserInGroup, userInGroupErr := group.IsUserInGroup(request.FromUser)
		if userInGroupErr != nil {
			return userInGroupErr
		}

		if !isUserInGroup {
			return appError.ValidationError{Err: fmt.Errorf(fmt.Sprintf("user %s not in group %s", request.FromUser, request.To))}
		}

	} else {
		//Check if to user exists
		toUserExists, err := ms.userRepository.CheckIfUserExists(request.To)

		if err != nil {
			return err
		}

		if !toUserExists {
			return appError.NotFoundError{Err: fmt.Errorf(fmt.Sprintf("to user %s not exists", request.To))}
		}
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

func (ms *MessageService) GetMessagesForUser(userId uuid.UUID) (*[]entity.Message, appError.AppError) {
	// Check if user exists
	user, err := ms.userRepository.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, appError.NotFoundError{Err: errors.New("user does not exist")}
	}

	messages, messageRepoErr := ms.messageRepository.GetMessagesForUser(userId)

	if messageRepoErr != nil {
		return nil, messageRepoErr
	}

	filteredMessages, filterMessageErr := user.FilterBlockedMessages(messages)
	if filterMessageErr != nil {
		return nil, filterMessageErr
	}

	return filteredMessages, nil
}
