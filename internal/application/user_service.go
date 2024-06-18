package application

import (
	"cmd/main.go/contracts"
	"cmd/main.go/internal/appError"
	"cmd/main.go/internal/domain/entity"
	"cmd/main.go/internal/domain/repository"
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
