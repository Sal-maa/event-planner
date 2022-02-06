package event

import _models "github.com/justjundana/event-planner/models"

type EventInterface interface {
	GetEvents() ([]_models.Event, error)
	Pagination(limit, offset *int) ([]_models.Event, error)
	GetEvent(id int) (_models.Event, error)
	SearchEvents(keyword string) ([]_models.Event, error)
	GetEventMostAttendant() ([]_models.Event, error)
	GetOwnEvent(userID int) ([]_models.Event, error)
	GetParticipateEvent(userID int) ([]_models.Event, error)
	CreateEvent(event _models.Event) error
	UpdateEvent(event _models.Event) error
	DeleteEvent(event _models.Event) error
}
