// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type NewUser struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	Occupation string `json:"occupation"`
}
