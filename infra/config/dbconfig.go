package config

import (
	"database/sql"
)

type Env struct {
	DB *sql.DB
}

func NewDB(s string, dt string) (*sql.DB, error) {
	db, err := sql.Open(s, dt)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
