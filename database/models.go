package database

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

// PreferenceType mirrors the 'preference_enum' in the database.
type PreferenceType string

const (
	Like    PreferenceType = "like"
	Dislike PreferenceType = "dislike"
	Neutral PreferenceType = "neutral"
)

// User corresponds to the 'users' table.
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// UserPreference corresponds to the 'user_preferences' table.
type UserPreference struct {
	ID        int64           `json:"id"`
	UserID    int64           `json:"userId"`
	Content   string          `json:"content"`
	Type      PreferenceType  `json:"type"`
	Embedding pgvector.Vector `json:"-"` // Omit vector from JSON for brevity
	CreatedAt time.Time       `json:"createdAt"`
}

// SimilarPreference is a struct to hold search results
// which include the distance metric.
type SimilarPreference struct {
	UserPreference
	Distance float64 `json:"distance"`
}
