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
	query := `INSERT INTO posts (user_id, content, image_url, parent_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing create post statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.UserID, post.Content, post.ImageURL, post.ParentID).Scan(&post.ID, &post.CreatedAt)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}

	return nil
}

func (r *PostRepository) GetRecentPosts(currentUserID int) ([]models.Post, error) {
	query := `SELECT p.id, p.user_id, u.username, u.name, u.avatar_url, p.content, p.created_at, p.image_url, p.parent_id,
		(SELECT COUNT(*) FROM likes WHERE post_id = p.id) AS likes_count,
		(SELECT COUNT(*) FROM posts WHERE parent_id = p.id) AS replies_count,
		EXISTS(SELECT 1 FROM likes WHERE post_id = p.id AND user_id = $1) AS user_liked
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.parent_id IS NULL
		ORDER BY p.created_at DESC`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing get recent posts statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(currentUserID)
	if err != nil {
		return nil, fmt.Errorf("error querying recent posts: %w", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Author, &post.AuthorName, &post.AuthorAvatarURL, &post.Content, &post.CreatedAt, &post.ImageURL, &post.ParentID, &post.LikesCount, &post.RepliesCount, &post.UserLiked)
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

func (r *PostRepository) GetReplies(postID, currentUserID int) ([]models.Post, error) {
	query := `SELECT p.id, p.user_id, u.username, u.name, u.avatar_url, p.content, p.created_at, p.image_url, p.parent_id,
		(SELECT COUNT(*) FROM likes WHERE post_id = p.id) AS likes_count,
		(SELECT COUNT(*) FROM posts WHERE parent_id = p.id) AS replies_count,
		EXISTS(SELECT 1 FROM likes WHERE post_id = p.id AND user_id = $2) AS user_liked
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.parent_id = $1
		ORDER BY p.created_at ASC`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing get replies statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(postID, currentUserID)
	if err != nil {
		return nil, fmt.Errorf("error querying replies: %w", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Author, &post.AuthorName, &post.AuthorAvatarURL, &post.Content, &post.CreatedAt, &post.ImageURL, &post.ParentID, &post.LikesCount, &post.RepliesCount, &post.UserLiked)
		if err != nil {
			return nil, fmt.Errorf("error scanning reply row: %w", err)
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating reply rows: %w", err)
	}

	return posts, nil
}

func (r *PostRepository) ToggleLike(userID, postID int) (bool, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return false, fmt.Errorf("error starting like transaction: %w", err)
	}
	defer tx.Rollback()

	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = $1 AND post_id = $2)", userID, postID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if like exists: %w", err)
	}

	if exists {
		_, err = tx.Exec("DELETE FROM likes WHERE user_id = $1 AND post_id = $2", userID, postID)
		if err != nil {
			return false, fmt.Errorf("error deleting like: %w", err)
		}

		if err := tx.Commit(); err != nil {
			return false, fmt.Errorf("error committing unlike: %w", err)
		}
		return false, nil
	}

	_, err = tx.Exec("INSERT INTO likes (user_id, post_id) VALUES ($1, $2)", userID, postID)
	if err != nil {
		return false, fmt.Errorf("error inserting like: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return false, fmt.Errorf("error committing like: %w", err)
	}
	return true, nil
}

func (r *PostRepository) GetLikesCount(postID int) (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = $1", postID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error getting likes count: %w", err)
	}
	return count, nil
}
