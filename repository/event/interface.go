package event

import _models "github.com/justjundana/event-planner/models"

type EventInterface interface {
	Get() ([]_models.Event, error)
	GetById(id int) (_models.Event, error)
	GetByKey(keyword string) ([]_models.Event, error)
	GetByLocation(location string) ([]_models.Event, error)
}
