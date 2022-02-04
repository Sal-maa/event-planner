package participant

import (
	"database/sql"
	"log"

	_models "github.com/justjundana/event-planner/models"
)

type ParticipantRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ParticipantRepository {
	return &ParticipantRepository{
		db: db,
	}
}

func (r *ParticipantRepository) GetParticipants(eventID int) ([]_models.Participant, error) {
	var participants []_models.Participant
	rows, err := r.db.Query(`SELECT id, user_id, event_id, status FROM participants WHERE status = 1 AND event_id = ?`, eventID)
	if err != nil {
		log.Fatalf("Error")
	}

	defer rows.Close()

	for rows.Next() {
		var participant _models.Participant

		err := rows.Scan(&participant.ID, &participant.UserID, &participant.EventID, &participant.Status)
		if err != nil {
			log.Fatalf("Error")
		}

		participants = append(participants, participant)
	}

	return participants, nil
}
