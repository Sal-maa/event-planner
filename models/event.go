package models

import (
	"time"
)

type Event struct {
	ID          int       `json:"id" form:"id"`
	UserID      int       `json:"user_id" form:"user_id"`
	Image       string    `json:"image" form:"image"`
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	Location    string    `json:"location" form:"location"`
	Date        time.Time `json:"date" form:"date"`
	Quota       int       `json:"quota" form:"quota"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
}
