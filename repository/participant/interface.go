package participant

import (
	_models "github.com/justjundana/event-planner/models"
)

type ParticipantInterface interface {
	GetParticipants(eventID int) ([]_models.Participant, error)
	GetParticipant(id int) (_models.Participant, error)
	CheckParticipant(userID int, eventID int) (_models.Participant, error)
	CreateParticipant(participant _models.Participant) error
	DeleteParticipant(participant _models.Participant) error
	UpdateParticipant(participant _models.Participant) error
}
