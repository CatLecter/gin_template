package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"src/internal/schemes"
)

type UserRepository struct{ db *pgxpool.Pool }

func NewUserRepository(db *pgxpool.Pool) UserRepositoryInterface { return &UserRepository{db: db} }

func (repo *UserRepository) CreateUser(user *schemes.UserRequest) (*schemes.UserResponse, error) {
	ctx := context.Background()
	conn, err := repo.db.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		log.Warnf("Error acquiring connection: %v", err.Error())
		return nil, err
	}
	row := conn.QueryRow(
		ctx, "INSERT INTO users(full_name, phone) VALUES($1, $2) RETURNING *", &user.FullName, &user.Phone,
	)
	userResp := schemes.UserResponse{}
	err = row.Scan(&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt)
	if err != nil {
		log.Warnf("Failed to insert user: %s", err.Error())
		return nil, err
	}
	return &userResp, nil
}

func (repo *UserRepository) GetUserByUUID(userUUID *uuid.UUID) (*schemes.UserResponse, error) {
	ctx := context.Background()
	conn, err := repo.db.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		log.Warnf("Error acquiring connection: %v", err.Error())
		return nil, err
	}
	row := conn.QueryRow(ctx, "SELECT * FROM users WHERE uuid = $1", &userUUID)
	userResp := schemes.UserResponse{}
	err = row.Scan(&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt)
	if err != nil {
		log.Warnf("Failed to get user: %s", err.Error())
		return nil, err
	}
	return &userResp, nil
}

func (repo *UserRepository) CheckUserByPhone(phone *string) (*bool, error) {
	ctx := context.Background()
	conn, err := repo.db.Acquire(ctx)
	defer conn.Release()

	if err != nil {
		log.Warnf("Error acquiring connection: %v", err.Error())
		return nil, err
	}
	row := conn.QueryRow(
		ctx, "SELECT CASE WHEN EXISTS (SELECT uuid FROM users WHERE phone = $1) THEN true ELSE false END", &phone,
	)
	var result bool
	err = row.Scan(&result)
	if err != nil {
		log.Warnf("Failed to check user: %s", err.Error())
		return nil, err
	}
	return &result, nil
}
