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

func (r *CommentRepository) GetComment(id int) (_models.Comment, error) {
	var comment _models.Comment

	row := r.db.QueryRow(`SELECT id, event_id, user_id, content FROM comments WHERE id = ?`, id)

	err := row.Scan(&comment.ID, &comment.EventID, &comment.UserID, &comment.Content)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *CommentRepository) CreateComment(comment _models.Comment) error {
	query := `INSERT INTO comments (event_id, user_id, content) VALUES (?, ?, ?)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(comment.EventID, comment.UserID, comment.Content)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) UpdateComment(comment _models.Comment) error {
	query := `UPDATE comments SET event_id = ?, user_id = ?, content = ? WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(comment.EventID, comment.UserID, comment.Content, comment.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) DeleteComment(comment _models.Comment) error {
	query := `DELETE FROM comments WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(comment.ID)
	if err != nil {
		return err
	}

	return nil
}
