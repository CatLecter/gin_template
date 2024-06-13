package database

import (
	"github.com/jmoiron/sqlx"
)

func NewDB(uri string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
