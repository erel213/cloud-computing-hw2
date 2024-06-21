package application

import (
	"errors"
	"whatsapp-like/contracts"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"
)

type GroupService struct {
	groupRepository repository.GroupRepository
	userRepository  repository.UserRepository
}

func NewGroupService(groupRepo repository.GroupRepository, userRepo repository.UserRepository) *GroupService {
	return &GroupService{
		groupRepository: groupRepo,
		userRepository:  userRepo,
	}
}

func (service *GroupService) CreateNewGroup(request contracts.CreateGroupRequest) (*contracts.CreateGroupResponse, appError.AppError) {
	//Check if created by user exists
	createdByExists, userExistsErr := service.userRepository.CheckIfUserExists(request.CreatedBy)
	if userExistsErr != nil {
		return nil, userExistsErr
	}

	if !createdByExists {
		return nil, appError.AuthenticationError{Err: errors.New("user not exists")}
	}

	//Create group
	group, entityErr := entity.NewGroup(request.GroupName, request.CreatedBy)

	if entityErr != nil {
		return nil, entityErr
	}

	repoErr := service.groupRepository.CreateGroup(group)

	if repoErr != nil {
		return nil, repoErr
	}

	groupResponse := contracts.CreateGroupResponse{
		GroupId:   group.GroupId,
		GroupName: group.GroupName,
	}

	return &groupResponse, nil
}
