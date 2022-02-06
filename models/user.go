package models

import (
	"time"
)

type User struct {
	ID         int       `json:"id" form:"id"`
	Avatar     string    `json:"avatar" form:"avatar"`
	Name       string    `json:"name" form:"name"`
	Email      string    `json:"email" form:"email"`
	Password   string    `json:"password" form:"password"`
	Address    string    `json:"address" form:"address"`
	Occupation string    `json:"occupation" form:"occupation"`
	Phone      string    `json:"phone" form:"phone"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at" form:"deleted_at"`
}
