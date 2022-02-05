package models

import (
	"time"
)

type Participant struct {
	ID        int       `json:"id" form:"id"`
	EventID   int       `json:"event_id" form:"event_id"`
	UserID    int       `json:"user_id" form:"user_id"`
	Status    bool      `json:"status" form:"status"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}
