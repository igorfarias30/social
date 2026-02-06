package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Author    int64    `json:"author"`
	UserId    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts (title, content, author, user_id, tags)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`

	error := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.Author,
		post.UserId,
		pq.Array(post.Tags),
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	if error != nil {
		return error
	}

	return nil
}
