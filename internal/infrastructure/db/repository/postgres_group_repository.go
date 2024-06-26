package postgresRepository

import (
	"database/sql"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"

	"github.com/google/uuid"
	"github.com/lib/pq"
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

func (repo *PostgresGroupRepository) CheckIfGroupExists(groupId uuid.UUID) (bool, appError.AppError) {
	query := `
	SELECT EXISTS (
		SELECT 1
		FROM groups
		WHERE group_id = $1
	)`

	var exists bool
	err := repo.DB.QueryRow(query, groupId).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, appError.InternalError{Err: err}
	}

	return exists, nil
}

func (repo *PostgresGroupRepository) CheckIfUserExistsInGroup(userId uuid.UUID, groupId uuid.UUID) (bool, appError.AppError) {
	query := `
	SELECT EXISTS (
		SELECT 1
		FROM user_group
		WHERE user_id = $1 AND group_id = $2
	)`
	var exists bool
	err := repo.DB.QueryRow(query, userId, groupId).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, appError.InternalError{Err: err}
	}

	return exists, nil
}

func (repo *PostgresGroupRepository) AddUserToGroup(userId uuid.UUID, groupId uuid.UUID) appError.AppError {
	userGroupId := uuid.New()

	query := `
	INSERT INTO user_group(user_group_id, user_id, group_id)
	VALUES ($1, $2, $3)`

	_, err := repo.DB.Exec(query, userGroupId, userId, groupId)

	if err != nil {
		return appError.InternalError{Err: err}
	}

	return nil
}

func (repo *PostgresGroupRepository) RemoveUserFromGroup(userId uuid.UUID, groupId uuid.UUID) appError.AppError {
	query := `
	DELETE FROM user_group
	WHERE user_id = $1 AND group_id = $2`

	_, err := repo.DB.Exec(query, userId, groupId)

	if err != nil {
		return appError.InternalError{Err: err}
	}

	return nil
}

func (repo *PostgresGroupRepository) GetGroupById(groupId uuid.UUID) (*entity.Group, appError.AppError) {
	query := `
	SELECT groups.group_id, groups.created_by, groups.group_name, array_agg(user_group.user_id)
	FROM groups
	JOIN user_group ON groups.group_id = user_group.group_id
	WHERE groups.group_id = $1
	GROUP BY groups.group_id`

	var group entity.Group
	err := repo.DB.QueryRow(query, groupId).Scan(&group.GroupId, &group.CreatedBy, &group.GroupName, pq.Array(&group.Users))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.NotFoundError{Err: err}
		}
		return nil, appError.InternalError{Err: err}
	}

	return &group, nil
}
