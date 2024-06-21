package postgresRepository

import (
	"database/sql"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"
)

type PostgresGroupRepository struct {
	DB *sql.DB
}

func NewPostgresGroupRepostiroy(db *sql.DB) repository.GroupRepository {
	return &PostgresGroupRepository{
		DB: db,
	}
}

func (repo *PostgresGroupRepository) CreateGroup(group *entity.Group) appError.AppError {
	query := `INSERT INTO groups(group_id, created_by, group_name)
	VALUES ($1, $2, $3)`

	_, err := repo.DB.Exec(query, group.GroupId, group.CreatedBy, group.GroupName)

	if err != nil {
		return appError.InternalError{Err: err}
	}
	return nil
}
