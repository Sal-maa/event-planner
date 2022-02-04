package user

import (
	_models "github.com/justjundana/event-planner/models"
)

type UserInterface interface {
	Register(user _models.User) (_models.User, error)
	Login(email string) (_models.User, error)
	Profile(id int) (_models.User, error)
	UpdateUser(user _models.User) error
	DeleteUser(user _models.User) error
}
