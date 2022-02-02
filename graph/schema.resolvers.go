package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	_generated "github.com/justjundana/event-planner/graph/generated"
	_model "github.com/justjundana/event-planner/graph/model"
	_models "github.com/justjundana/event-planner/models"
)

func (r *mutationResolver) Register(ctx context.Context, input *_model.NewUser) (*_models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*_model.LoginResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfile(ctx context.Context) (*_models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfileEvent(ctx context.Context) ([]*_models.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*_models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (*_models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetEvents(ctx context.Context) ([]*_models.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetEvent(ctx context.Context, id int) (*_models.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetEventKeyword(ctx context.Context, keyword string) (*_models.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetEventLocation(ctx context.Context, location string) (*_models.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetComments(ctx context.Context, eventID int) ([]*_models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetParticipants(ctx context.Context, eventID int) ([]*_models.Participant, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns _generated.MutationResolver implementation.
func (r *Resolver) Mutation() _generated.MutationResolver { return &mutationResolver{r} }

// Query returns _generated.QueryResolver implementation.
func (r *Resolver) Query() _generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
