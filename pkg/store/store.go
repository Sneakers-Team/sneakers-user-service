package store

import (
	"os"

	"github.com/jackc/pgx"
)

type Store struct {
	Config *Config
	Db     *pgx.Conn
}

func NewStore(config *Config) (*Store, error) {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     5432,
		Database: os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})

	if err != nil {
		return nil, err
	}

	return &Store{
		Config: config,
		Db:     conn,
	}, nil
}
