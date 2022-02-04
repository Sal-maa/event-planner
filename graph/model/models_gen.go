// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type EditComment struct {
	Content *string `json:"content"`
}

type EditEvent struct {
	Image       *string    `json:"image"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Location    *string    `json:"location"`
	Date        *time.Time `json:"date"`
	Quota       *int       `json:"quota"`
}

type EditParticipant struct {
	Status *bool `json:"status"`
}

type EditUser struct {
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
	Address    *string `json:"address"`
	Occupation *string `json:"occupation"`
	Phone      *string `json:"phone"`
}

type LoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type NewComment struct {
	EventID int    `json:"eventID"`
	Content string `json:"content"`
}

type NewEvent struct {
	UserID      *int      `json:"userID"`
	Image       string    `json:"image"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	Quota       int       `json:"quota"`
}

type NewParticipant struct {
	EventID int  `json:"eventID"`
	Status  bool `json:"status"`
}

type NewUser struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	Occupation string `json:"occupation"`
	Phone      string `json:"phone"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
