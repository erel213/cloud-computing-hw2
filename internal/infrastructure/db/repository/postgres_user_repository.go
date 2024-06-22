package postgresRepository

import (
	"database/sql"
	"whatsapp-like/internal/appError"
	"whatsapp-like/internal/domain/entity"
	"whatsapp-like/internal/domain/repository"

	"github.com/google/uuid"
	"github.com/lib/pq"
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

func (p *PostgresUserRepository) BlockUser(userId uuid.UUID, blockUserId uuid.UUID) appError.AppError {
	generatedId, err := uuid.NewUUID()
	if err != nil {
		return appError.InternalError{Err: err}
	}

	query := `
	INSERT INTO blocked_users (id, user_block, user_blocked)
	VALUES ($1, $2, $3)
	`

	_, err = p.db.Exec(query, generatedId, userId, blockUserId)

	if err != nil {
		return appError.InternalError{Err: err}
	}

	return nil
}

func (p *PostgresUserRepository) GetUserById(userId uuid.UUID) (*entity.User, appError.AppError) {
	query := `
	select 
		u.user_id,  array_agg(bu.user_blocked)  
	from
		 users u 
	left 
		join blocked_users bu
	on
		bu.user_block = u.user_id 
	where u.user_id = $1
	group by u.user_id 
	`

	var user entity.User
	err := p.db.QueryRow(query, userId).Scan(&user.UserId, pq.Array(&user.BlockedUsers))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.NotFoundError{Err: err}
		}
		return nil, appError.InternalError{Err: err}
	}

	return &user, nil
}
