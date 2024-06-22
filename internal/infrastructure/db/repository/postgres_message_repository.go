package postgresRepository

import (
	"database/sql"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"

	"github.com/google/uuid"
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

func (repo *PostgresMessageRepository) GetMessagesForUser(userId uuid.UUID) ([]*entity.Message, appError.AppError) {
	query := `
	SELECT 
		message_id,
		from_user,
		recpient,
		message_content 
	FROM (
		SELECT 
			message_id,
			from_user,
			CASE 
				WHEN to_group = false THEN send_to 
				ELSE ug.user_id 
			END AS recpient,
			m.message_content 
		FROM 
			messages m 
		LEFT JOIN 
			user_group ug 
		ON 
			m.send_to = ug.group_id
	) AS subquery
	WHERE 
		recpient = $1;`

	rows, err := repo.db.Query(query, userId)

	if err != nil {
		return nil, appError.InternalError{Err: err}
	}

	defer rows.Close()

	messages := make([]*entity.Message, 0)

	for rows.Next() {
		var message entity.Message
		err := rows.Scan(&message.MessageId, &message.FromUser, &message.To, &message.MessageBody)

		if err != nil {
			return nil, appError.InternalError{Err: err}
		}

		messages = append(messages, &message)
	}

	return messages, nil

}
