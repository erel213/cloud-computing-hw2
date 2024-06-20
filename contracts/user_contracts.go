package contracts

import "github.com/google/uuid"

type CreateUserResponse struct {
	UserId uuid.UUID `json:"user_id"`
}
