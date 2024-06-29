package application

import (
	"errors"
	"fmt"
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

	addUserToGroupErr := service.groupRepository.AddUserToGroup(group.CreatedBy, group.GroupId)
	if addUserToGroupErr != nil {
		return nil, addUserToGroupErr
	}

	//Add created user to the group

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
		return appError.NotFoundError{Err: fmt.Errorf("user %s not exists", request.UserId)}
	}

	//Check if group exists
	group, groupExistsErr := service.groupRepository.GetGroupById(request.GroupId)
	if groupExistsErr != nil {
		return groupExistsErr
	}

	if group == nil {
		return appError.NotFoundError{Err: fmt.Errorf("group %s not found", request.GroupId)}
	}

	//Add user to a group
	addUserErr := group.AddUser(request.UserId)
	if addUserErr != nil {
		return addUserErr
	}

	repoErr := service.groupRepository.AddUserToGroup(request.UserId, request.GroupId)
	if repoErr != nil {
		return repoErr
	}

	return nil
}

func (service *GroupService) RemoveUserFromGroup(request contracts.RemoveUserFromGroupRequest) appError.AppError {
	group, getGroupErr := service.groupRepository.GetGroupById(request.GroupId)

	if getGroupErr != nil {
		return getGroupErr
	}

	if group == nil {
		return appError.NotFoundError{Err: fmt.Errorf("group %s not exists", request.GroupId)}
	}

	removeUserErr := group.RemoveUser(request.UserId)
	if removeUserErr != nil {
		return removeUserErr
	}

	repoErr := service.groupRepository.RemoveUserFromGroup(request.UserId, request.GroupId)
	if repoErr != nil {
		return repoErr
	}

	return nil
}
