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

func (service *GroupService) AddUserToGroup(request contracts.AddUserToGroupRequest) appError.AppError {
	//Check if user exists
	userExists, userExistsErr := service.userRepository.CheckIfUserExists(request.UserId)
	if userExistsErr != nil {
		return userExistsErr
	}

	if !userExists {
		return appError.AuthenticationError{Err: errors.New("user not exists")}
	}

	//Check if group exists
	groupExists, groupExistsErr := service.groupRepository.CheckIfGroupExists(request.GroupId)
	if groupExistsErr != nil {
		return groupExistsErr
	}

	//Check if user already exists in group
	userExistsInGroup, userExistsInGroupErr := service.groupRepository.CheckIfUserExistsInGroup(request.UserId, request.GroupId)
	if userExistsInGroupErr != nil {
		return userExistsInGroupErr
	}

	if userExistsInGroup {
		return appError.ValidationError{Err: errors.New("user already exists in group")}
	}

	if !groupExists {
		return appError.ValidationError{Err: errors.New("group not exists")}
	}

	//Add user to group
	repoErr := service.groupRepository.AddUserToGroup(request.UserId, request.GroupId)

	if repoErr != nil {
		return repoErr
	}

	return nil
}

func (service *GroupService) RemoveUserFromGroup(request contracts.RemoveUserFromGroupRequest) appError.AppError {
	//Check if group existw
	groupExists, groupExistsErr := service.groupRepository.CheckIfGroupExists(request.GroupId)
	if groupExistsErr != nil {
		return groupExistsErr
	}

	if !groupExists {
		return appError.ValidationError{Err: errors.New("group not exists")}
	}

	//Check if user exists in group
	userExistsInGroup, userExistsInGroupErr := service.groupRepository.CheckIfUserExistsInGroup(request.UserId, request.GroupId)
	if userExistsInGroupErr != nil {
		return userExistsInGroupErr
	}

	if !userExistsInGroup {
		return appError.ValidationError{Err: errors.New("user not exists in group")}
	}

	//Remove user from group
	repoErr := service.groupRepository.RemoveUserFromGroup(request.UserId, request.GroupId)

	if repoErr != nil {
		return repoErr
	}

	return nil
}
