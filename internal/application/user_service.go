package application

import (
	"fmt"
	"whatsapp-like/contracts"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (service *UserService) CreateUser() (contracts.CreateUserResponse, appError.AppError) {
	userEntity, createUserErr := entity.NewUser()
	if createUserErr != nil {
		return contracts.CreateUserResponse{}, createUserErr
	}

	// Fix: Assign the result of the function call to a variable
	_, err := service.userRepo.CreateUser(userEntity)
	if err != nil {
		return contracts.CreateUserResponse{}, err
	}

	return contracts.CreateUserResponse{
		UserId: userEntity.UserId,
	}, nil
}

func (service *UserService) BlockUser(userId uuid.UUID, blockedUserId uuid.UUID) (*contracts.UserResponse, appError.AppError) {
	//Get user
	user, getUserErr := service.userRepo.GetUserById(userId)

	if getUserErr != nil {
		return nil, getUserErr
	}

	if user == nil {
		return nil, appError.NotFoundError{Err: fmt.Errorf("user %s not found", userId)}
	}

	blockedUser, blockedUserErr := service.userRepo.GetUserById(blockedUserId)

	if blockedUserErr != nil {
		return nil, blockedUserErr
	}

	if blockedUser == nil {
		return nil, appError.NotFoundError{Err: fmt.Errorf("blocked user %s not exist", blockedUserId)}
	}

	blockingUserErr := user.BlockUser(blockedUser.UserId)

	if blockingUserErr != nil {
		return nil, blockingUserErr
	}

	repoErr := service.userRepo.BlockUser(user.UserId, blockedUser.UserId)

	if repoErr != nil {
		return nil, repoErr
	}

	respone := contracts.UserResponse{
		UserId:       user.UserId,
		BlockedUsers: user.BlockedUsers,
	}

	return &respone, nil
}

func (service *UserService) GetUserById(userId uuid.UUID) (*contracts.UserResponse, appError.AppError) {
	user, getUserErr := service.userRepo.GetUserById(userId)

	if getUserErr != nil {
		return nil, getUserErr
	}

	if user == nil {
		return nil, appError.NotFoundError{Err: fmt.Errorf("user %s not found", userId)}
	}

	response := contracts.UserResponse{
		UserId:       user.UserId,
		BlockedUsers: user.BlockedUsers,
	}

	return &response, nil
}
