-- Filename: 02_create_schema.sql
-- Description: Creates the application schema (tables, types, etc.).
-- This script should be run by the application user (e.g., ago_user_data).

-- Exit on error
\set ON_ERROR_STOP on

-- Create a custom type for preference categories for better data integrity
CREATE TYPE preference_enum AS ENUM ('like', 'dislike', 'neutral');

-- Table to store basic user information
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Table to store individual user preferences and their vector embeddings
CREATE TABLE user_preferences (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    preference_type preference_enum NOT NULL,
    embedding vector(384),
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Create an HNSW index for fast vector similarity searches
CREATE INDEX ON user_preferences USING hnsw (embedding vector_cosine_ops);