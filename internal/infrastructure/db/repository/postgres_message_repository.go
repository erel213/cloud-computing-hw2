package postgresRepository

import (
	"cmd/main.go/internal/appError"
	"cmd/main.go/internal/domain/entity"
	"cmd/main.go/internal/domain/repository"
	"database/sql"
)

type PostgresMessageRepository struct {
	db *sql.DB
}

func NewPostgresMessageRepository(db *sql.DB) repository.MessageRepository {
	return &PostgresMessageRepository{db}
}

func (repo *PostgresMessageRepository) CreateMessage(message *entity.Message) appError.AppError {
	query := `
	INSERT into Message (
		message_id,
		from_user,
		to,
		message_body,
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
