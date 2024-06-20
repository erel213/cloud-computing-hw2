package postgresRepository

import (
	"database/sql"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"
)

type PostgresMessageRepository struct {
	db *sql.DB
}

func NewPostgresMessageRepository(db *sql.DB) repository.MessageRepository {
	return &PostgresMessageRepository{db}
}

func (repo *PostgresMessageRepository) CreateMessage(message *entity.Message) appError.AppError {
	query := `
	INSERT into messages (
		message_id,
		from_user,
		send_to,
		message_content,
		to_group
	)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := repo.db.Exec(query, message.MessageId, message.FromUser, message.To, message.MessageBody, message.ToGroup)

	if err != nil {
		return appError.InternalError{Err: err}
	}

	return nil
}
