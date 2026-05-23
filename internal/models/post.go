package models

import "time"

type Post struct {
    ID        int       `json:"id" db:"id"`
    UserID    int       `json:"user_id" db:"user_id"`
    Content   string    `json:"content" db:"content"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    Author    string    `json:"author,omitempty" db:"author"`
}
