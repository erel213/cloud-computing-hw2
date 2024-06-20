package postgresRepository

import (
	"cmd/main.go/internal/appError"
	"cmd/main.go/internal/domain/entity"
	"cmd/main.go/internal/domain/repository"
	"database/sql"

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
	INSERT into User (
		user_id,
	VALUES ($1)
	)`

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
		FROM User
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
