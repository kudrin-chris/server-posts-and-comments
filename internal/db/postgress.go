package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresOptions struct {
	Host, Port, User, Password, Name string
}

func NewPostgresDB(p PostgresOptions) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.Host, p.Port, p.User, p.Password, p.Name))
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	return db, nil
}
