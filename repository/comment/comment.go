package comment

import (
	"database/sql"
	"log"

	_models "github.com/justjundana/event-planner/models"
)

type CommentRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) GetComments(eventID int) ([]_models.Comment, error) {
	var comments []_models.Comment
	rows, err := r.db.Query(`SELECT id, user_id, event_id, content FROM comments WHERE event_id = ?`, eventID)
	if err != nil {
		log.Fatalf("Error")
	}

	defer rows.Close()

	for rows.Next() {
		var comment _models.Comment

		err := rows.Scan(&comment.ID, &comment.UserID, &comment.EventID, &comment.Content)
		if err != nil {
			log.Fatalf("Error")
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
