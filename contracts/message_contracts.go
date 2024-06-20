package contracts

import "github.com/google/uuid"

type SendMessageRequest struct {
	FromUser    uuid.UUID `json:"from_user"`
	To          uuid.UUID `json:"to"`
	MessageBody string    `json:"message_body"`
	ToGroup     bool      `json:"to_group"`
}
