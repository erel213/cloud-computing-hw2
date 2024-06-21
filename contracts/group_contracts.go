package contracts

import "github.com/google/uuid"

type CreateGroupRequest struct {
	GroupName string
	CreatedBy uuid.UUID
}

type CreateGroupResponse struct {
	GroupId   uuid.UUID
	GroupName string
}

type AddUserToGroupRequest struct {
	UserId  uuid.UUID
	GroupId uuid.UUID
}

type RemoveUserFromGroupRequest struct {
	UserId  uuid.UUID
	GroupId uuid.UUID
}
