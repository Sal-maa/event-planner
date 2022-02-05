package models

import (
	"time"
)

type Comment struct {
	ID        int       `json:"id" form:"id"`
	UserID    int       `json:"user_id" form:"user_id"`
	EventID   int       `json:"event_id" form:"event_id"`
	Content   string    `json:"content" form:"content"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}
