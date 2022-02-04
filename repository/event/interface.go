package event

import _models "github.com/justjundana/event-planner/models"

type EventInterface interface {
	GetEvents() ([]_models.Event, error)
	GetEvent(id int) (_models.Event, error)
	GetEventKeyword(keyword string) ([]_models.Event, error)
	GetEventLocation(location string) ([]_models.Event, error)
	GetOwnEvent(userID int) ([]_models.Event, error)
	CreateEvent(event _models.Event) error
}
