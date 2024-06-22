package contracts

import "github.com/google/uuid"

type CreateUserResponse struct {
	UserId uuid.UUID `json:"user_id"`
}

type UserResponse struct {
	UserId       uuid.UUID
	BlockedUsers []uuid.UUID
}

type BlockUserRequest struct {
	UserId        uuid.UUID
	BlockedUserId uuid.UUID
}
