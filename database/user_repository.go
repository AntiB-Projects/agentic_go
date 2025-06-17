package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// UserStorer defines the database operations for users.
type UserStorer interface {
	CreateUser(ctx context.Context, email string) (*User, error)
}

// UserRepository implements the UserStorer interface.
type UserRepository struct {
	conn *pgxpool.Pool
}

// CreateUser inserts a new user into the database.
func (r *UserRepository) CreateUser(ctx context.Context, email string) (*User, error) {
	query := `INSERT INTO users (email) VALUES ($1) RETURNING id, created_at;`

	user := &User{Email: email}
	err := r.conn.QueryRow(ctx, query, email).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
