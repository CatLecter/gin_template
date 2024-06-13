package repositories

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"src/internal/schemes"
)

type UserRepository struct{ db *sqlx.DB }

func NewUserRepository(db *sqlx.DB) User { return &UserRepository{db: db} }

func (repo *UserRepository) CreateUser(user schemes.UserRequest) (schemes.UserResponse, error) {
	userResp := schemes.UserResponse{}
	row := repo.db.QueryRow(
		"INSERT INTO users(full_name, phone) VALUES($1, $2) RETURNING *", user.FullName, user.Phone,
	)
	err := row.Scan(&userResp.UUID, &userResp.FullName, &userResp.Phone, &userResp.CreatedAt, &userResp.UpdatedAt)
	if err != nil {
		log.Errorf("Failed to insert user: %s", err.Error())
		return userResp, err
	}
	return userResp, nil
}
