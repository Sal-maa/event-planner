package event

import (
	"database/sql"
	"log"

	_models "github.com/justjundana/event-planner/models"
)

type EventRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) GetEvents() ([]_models.Event, error) {
	var events []_models.Event
	rows, err := r.db.Query(`SELECT id, user_id, image, title, description, location, date, quota FROM events ORDER BY id ASC`)
	if err != nil {
		log.Fatalf("Error")
	}

	defer rows.Close()

	for rows.Next() {
		var event _models.Event

		err = rows.Scan(&event.ID, &event.UserID, &event.Image, &event.Title, &event.Description, &event.Location, &event.Date, &event.Quota)
		if err != nil {
			log.Fatalf("Error")
		}

		events = append(events, event)
	}

	return events, nil
}

func (r *EventRepository) GetEvent(id int) (_models.Event, error) {
	var event _models.Event

	row := r.db.QueryRow(`SELECT id, user_id, image, title, description, location, date, quota FROM events WHERE id = ?`, id)

	err := row.Scan(&event.ID, &event.UserID, &event.Image, &event.Title, &event.Description, &event.Location, &event.Date, &event.Quota)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *EventRepository) GetEventKeyword(keyword string) ([]_models.Event, error) {
	var events []_models.Event
	rows, err := r.db.Query(`SELECT id, user_id, image, title, description, location, date, quota FROM events WHERE title LIKE ? `, "%"+keyword+"%")
	if err != nil {
		log.Fatalf("Error")
	}
	defer rows.Close()

	for rows.Next() {
		var event _models.Event

		err := rows.Scan(&event.ID, &event.UserID, &event.Image, &event.Title, &event.Description, &event.Location, &event.Date, &event.Quota)
		if err != nil {
			log.Fatalf("Error")
		}

		events = append(events, event)
	}

	return events, nil
}

func (r *EventRepository) GetEventLocation(location string) ([]_models.Event, error) {
	var events []_models.Event
	rows, err := r.db.Query(`SELECT id, user_id, image, title, description, location, date, quota FROM events WHERE location LIKE ?`, "%"+location+"%")
	if err != nil {
		log.Fatalf("Error")
	}
	defer rows.Close()

	for rows.Next() {
		var event _models.Event

		err := rows.Scan(&event.ID, &event.UserID, &event.Image, &event.Title, &event.Description, &event.Location, &event.Date, &event.Quota)
		if err != nil {
			log.Fatalf("Error")
		}

		events = append(events, event)
	}

	return events, nil
}

func (r *EventRepository) GetOwnEvent(userID int) ([]_models.Event, error) {
	var events []_models.Event
	rows, err := r.db.Query(`SELECT id, user_id, image, title, description, location, date, quota FROM events WHERE user_id = ?`, userID)
	if err != nil {
		log.Fatalf("Error")
	}

	defer rows.Close()

	for rows.Next() {
		var event _models.Event

		err := rows.Scan(&event.ID, &event.UserID, &event.Image, &event.Title, &event.Description, &event.Location, &event.Date, &event.Quota)
		if err != nil {
			log.Fatalf("Error")
		}

		events = append(events, event)
	}

	return events, nil
}
