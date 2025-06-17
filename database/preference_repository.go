package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pgvector/pgvector-go"
)

// PreferenceStorer defines the database operations for user preferences.
type PreferenceStorer interface {
	CreatePreference(ctx context.Context, p *UserPreference) error
	FindSimilarPreferences(ctx context.Context, userID int64, queryVector pgvector.Vector, limit int) ([]SimilarPreference, error)
}

// PreferenceRepository implements the PreferenceStorer interface.
type PreferenceRepository struct {
	conn *pgxpool.Pool
}

// CreatePreference inserts a new user preference into the database.
func (r *PreferenceRepository) CreatePreference(ctx context.Context, p *UserPreference) error {
	query := `
        INSERT INTO user_preferences (user_id, content, preference_type, embedding)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at;
    `
	err := r.conn.QueryRow(ctx, query, p.UserID, p.Content, p.Type, p.Embedding).Scan(&p.ID, &p.CreatedAt)
	if err != nil {
		return errors.New("failed to create user preference: %w" + err.Error())
	}
	return nil
}

// FindSimilarPreferences performs a semantic search using cosine distance.
func (r *PreferenceRepository) FindSimilarPreferences(ctx context.Context, userID int64, queryVector pgvector.Vector, limit int) ([]SimilarPreference, error) {
	query := `
        SELECT id, user_id, content, preference_type, embedding, created_at, embedding <=> $2 AS distance
        FROM user_preferences
        WHERE user_id = $1
        ORDER BY distance ASC
        LIMIT $3;
    `
	rows, err := r.conn.Query(ctx, query, userID, queryVector, limit)
	if err != nil {
		return nil, fmt.Errorf("query for similar preferences failed: %w", err)
	}
	defer rows.Close()

	var results []SimilarPreference
	for rows.Next() {
		var p SimilarPreference
		if err := rows.Scan(&p.ID, &p.UserID, &p.Content, &p.Type, &p.Embedding, &p.CreatedAt, &p.Distance); err != nil {
			return nil, fmt.Errorf("failed to scan preference row: %w", err)
		}
		results = append(results, p)
	}

	return results, nil
}
