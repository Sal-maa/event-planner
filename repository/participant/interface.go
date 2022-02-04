package participant

import (
	_models "github.com/justjundana/event-planner/models"
)

type ParticipantInterface interface {
	GetParticipants(eventID int) ([]_models.Participant, error)
}
