package postgresRepository

import (
	"database/sql"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"

	"github.com/google/uuid"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repository.UserRepository {
	return &PostgresUserRepository{db}
}

func (p *PostgresUserRepository) CreateUser(user *entity.User) (entity.User, appError.AppError) {
	query := `
	INSERT into users (user_id)
	VALUES ($1)
	`

	_, err := p.db.Exec(query, user.UserId)

	if err != nil {
		return entity.User{}, appError.InternalError{Err: err}
	}

	return *user, nil
}

func (p *PostgresUserRepository) CheckIfUserExists(userId uuid.UUID) (bool, appError.AppError) {
	query := `
	SELECT EXISTS (
		SELECT 1
		FROM users
		WHERE user_id = $1
	)`

	var exists bool
	err := p.db.QueryRow(query, userId).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, appError.InternalError{Err: err}
	}

	return exists, nil
}
