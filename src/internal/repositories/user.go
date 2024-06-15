package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"src/internal/schemes"
)

type UserRepository struct{ db *pgxpool.Pool }

func NewUserRepository(db *pgxpool.Pool) User { return &UserRepository{db: db} }

func (repo *UserRepository) CreateUser(user *schemes.UserRequest) (*schemes.UserResponse, error) {
	ctx := context.Background()
	conn, err := repo.db.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		log.Errorf("Error acquiring connection: %v", err.Error())
		return nil, err
	}
	row := conn.QueryRow(
		ctx, "INSERT INTO users(full_name, phone) VALUES($1, $2) RETURNING *", &user.FullName, &user.Phone,
	)
	userResp := schemes.UserResponse{}
	err = row.Scan(&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt)
	if err != nil {
		log.Errorf("Failed to insert user: %s", err.Error())
		return nil, err
	}
	return &userResp, nil
}

func (repo *UserRepository) GetUserByUUID(userUUID *uuid.UUID) (*schemes.UserResponse, error) {
	ctx := context.Background()
	conn, err := repo.db.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		log.Errorf("Error acquiring connection: %v", err.Error())
		return nil, err
	}
	userResp := schemes.UserResponse{}
	row := conn.QueryRow(ctx, "SELECT * FROM users WHERE uuid = $1", &userUUID)
	err = row.Scan(&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt)
	if err != nil {
		log.Errorf("Failed to get user: %s", err.Error())
		return &userResp, err
	}
	return &userResp, nil
}
