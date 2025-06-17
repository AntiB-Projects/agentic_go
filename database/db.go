package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxvector "github.com/pgvector/pgvector-go/pgx"
)

// Database holds the connection pool and repositories.
type Database struct {
	pool        *pgxpool.Pool
	Users       UserStorer
	Preferences PreferenceStorer
}

// New creates a new database connection pool and registers the pgvector type.
func New(ctx context.Context, databaseURL string) (*Database, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	// This is the key part for pgvector-go integration.
	// We register the custom 'vector' type with pgx.
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return pgxvector.RegisterTypes(ctx, conn)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	// Create the Database struct and attach repository implementations
	db := &Database{
		pool: pool,
	}
	db.Users = &UserRepository{conn: db.pool}
	db.Preferences = &PreferenceRepository{conn: db.pool}

	return db, nil
}

// Close gracefully closes the database connection pool.
func (db *Database) Close() {
	db.pool.Close()
}
