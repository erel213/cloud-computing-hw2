package application

import (
	"whatsapp-like/contracts"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (service *UserService) CreateUser() (contracts.CreateUserResponse, appError.AppError) {
	userEntity, createUserErr := entity.NewUser()
	if createUserErr != nil {
		return contracts.CreateUserResponse{}, createUserErr
	}

	// Fix: Assign the result of the function call to a variable
	_, err := service.UserRepo.CreateUser(userEntity)
	if err != nil {
		return contracts.CreateUserResponse{}, err
	}

	return contracts.CreateUserResponse{
		UserId: userEntity.UserId,
	}, nil
}
