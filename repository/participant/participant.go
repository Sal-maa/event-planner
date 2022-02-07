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
	rows, err := r.db.Query(`
	SELECT 
		participants.id, participants.user_id, participants.event_id, participants.status,
		users.id, users.avatar, users.name, users.email, users.address, users.occupation, users.phone
	FROM 
		participants
	JOIN
		users ON users.id = participants.user_id
	WHERE status = 1 AND event_id = ?`, eventID)

	if err != nil {
		log.Fatalf("Error")
	}

	defer rows.Close()

	for rows.Next() {
		var participant _models.Participant

		err := rows.Scan(&participant.ID, &participant.UserID, &participant.EventID, &participant.Status, &participant.User.ID, &participant.User.Avatar, &participant.User.Name, &participant.User.Email, &participant.User.Address, &participant.User.Occupation, &participant.User.Phone)
		if err != nil {
			log.Fatalf("Error")
		}

		participants = append(participants, participant)
	}

	return participants, nil
}

func (r *ParticipantRepository) GetParticipant(id int) (_models.Participant, error) {
	var participant _models.Participant
	row, err := r.db.Query(`SELECT id, user_id, event_id, status FROM participants WHERE status = 1 AND id = ?`, id)
	if err != nil {
		log.Fatalf("Error")
	}

	defer row.Close()

	for row.Next() {
		err := row.Scan(&participant.ID, &participant.UserID, &participant.EventID, &participant.Status)
		if err != nil {
			log.Fatalf("Error")
		}
	}

	return participant, nil
}

func (r *ParticipantRepository) CheckParticipant(userID int, eventID int) (_models.Participant, error) {
	var participant _models.Participant

	row := r.db.QueryRow(`SELECT COUNT(user_id) as data FROM participants WHERE user_id = ? AND event_id = ? GROUP BY user_id`, userID, eventID)

	err := row.Scan(&participant.UserID)
	if err != nil {
		return participant, err
	}

	return participant, nil
}

func (r *ParticipantRepository) CreateParticipant(participant _models.Participant) error {
	query := `INSERT INTO participants (event_id, user_id, status) VALUES (?, ?, ?)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(participant.EventID, participant.UserID, participant.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *ParticipantRepository) DeleteParticipant(participant _models.Participant) error {
	query := `DELETE FROM participants WHERE event_id = ? AND user_id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(participant.EventID, participant.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ParticipantRepository) UpdateParticipant(participant _models.Participant) error {
	query := `UPDATE participants SET status = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(participant.Status)
	if err != nil {
		return err
	}

	return nil
}
