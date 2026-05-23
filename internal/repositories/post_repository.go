package repositories

import (
	"database/sql"
	"fmt"

	"red_social/internal/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	query := `INSERT INTO posts (user_id, content)
		VALUES ($1, $2)
		RETURNING id, created_at`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing create post statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.UserID, post.Content).Scan(&post.ID, &post.CreatedAt)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}

	return nil
}

func (r *PostRepository) GetRecentPosts() ([]models.Post, error) {
	query := `SELECT p.id, p.user_id, p.content, p.created_at, u.username AS author
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY p.created_at DESC`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing get recent posts statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("error querying recent posts: %w", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.Author)
		if err != nil {
			return nil, fmt.Errorf("error scanning post row: %w", err)
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating post rows: %w", err)
	}

	return posts, nil
}
