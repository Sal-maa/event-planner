package models

import (
	"time"
)

type Event struct {
	ID          int       `json:"id" form:"id"`
	UserID      int       `json:"user_id" form:"user_id"`
	Image       string    `json:"image" form:"image"`
	Title       string    `json:"title" form:"title"`
	CategoryId  int       `json:"category_id" form:"category_id"`
	Description string    `json:"description" form:"description"`
	Location    string    `json:"location" form:"location"`
	Date        time.Time `json:"date" form:"date"`
	Quota       int       `json:"quota" form:"quota"`
}
