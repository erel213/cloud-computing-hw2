package postgres_user_repository

import (
	"cmd/main.go/internal/appError"
	"cmd/main.go/internal/domain/entity"
	"cmd/main.go/internal/domain/repository"
	"database/sql"
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
