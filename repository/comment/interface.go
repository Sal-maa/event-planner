package comment

import (
	_models "github.com/justjundana/event-planner/models"
)

type CommentInterface interface {
	GetComments(eventID int) ([]_models.Comment, error)
	GetComment(id int) (_models.Comment, error)
	CreateComment(comment _models.Comment) error
}
