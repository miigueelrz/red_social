package models

import "time"

type Post struct {
    ID         int       `json:"id" db:"id"`
    UserID     int       `json:"user_id" db:"user_id"`
    Content    string    `json:"content" db:"content"`
    CreatedAt  time.Time `json:"created_at" db:"created_at"`
    ImageURL   string    `json:"image_url" db:"image_url"`
    Author     string    `json:"author,omitempty" db:"author"`
    LikesCount int       `json:"likes_count" db:"likes_count"`
    UserLiked  bool      `json:"user_liked" db:"user_liked"`
}
