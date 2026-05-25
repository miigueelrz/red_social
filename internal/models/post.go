package models

import "time"

type Post struct {
	ID           int       `json:"id" db:"id"`
	UserID       int       `json:"user_id" db:"user_id"`
	Author       string    `json:"author,omitempty" db:"author"`
	Content      string    `json:"content" db:"content"`
	ImageURL     string    `json:"image_url" db:"image_url"`
	ParentID     *int      `json:"parent_id" db:"parent_id"`
	LikesCount   int       `json:"likes_count" db:"likes_count"`
	RepliesCount int       `json:"replies_count" db:"replies_count"`
	UserLiked    bool      `json:"user_liked" db:"user_liked"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	Replies      []Post    `json:"replies,omitempty" db:"-"`
}
