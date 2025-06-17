package database

import (
	"context"
	"errors"

	"github.com/AntiB-Projects/agentic_go/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB is the global database connection pool
var DB *pgxpool.Pool

func Init(cfg *config.Config) error {

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return errors.New("failed to connect to db: " + err.Error())
	}

	// test connection
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return errors.New("failed to ping db: " + err.Error())
	}

	DB = pool
	return nil
}
