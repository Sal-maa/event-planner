package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	_generated "github.com/justjundana/event-planner/graph/generated"
	_model "github.com/justjundana/event-planner/graph/model"
	_middleware "github.com/justjundana/event-planner/middleware"
	_models "github.com/justjundana/event-planner/models"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) Register(ctx context.Context, input *_model.NewUser) (*_models.User, error) {
	userData := _models.User{}
	userData.Name = input.Name
	userData.Email = input.Email
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	userData.Password = string(passwordHash)
	userData.Address = input.Address
	userData.Occupation = input.Occupation

	responseData, err := r.userRepository.Create(userData)
	return &responseData, err
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*_model.LoginResponse, error) {
	var user _models.User

	user, err := r.userRepository.Login(email)
	if err != nil {
		return &_model.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	token, _ := _middleware.AuthService().GenerateToken(user.ID)

	return &_model.LoginResponse{
		ID:    strconv.Itoa(user.ID),
		Token: token,
	}, err
}

func (r *queryResolver) GetProfile(ctx context.Context) (*_models.User, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_models.User{}, errors.New("unauthorized")
	}

	responseData, err := r.userRepository.Profile(userId.ID)
	if err != nil {
		return &responseData, err
	}

	dataUser := _models.User{
		ID:         responseData.ID,
		Name:       responseData.Name,
		Email:      responseData.Email,
		Password:   responseData.Password,
		Address:    responseData.Address,
		Occupation: responseData.Occupation,
		Phone:      responseData.Phone,
	}

	return &dataUser, nil
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
	eventResponseData := []*_models.Event{}

	responseData, err := r.eventRepository.Get()
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, v := range responseData {
		eventResponseData = append(eventResponseData, &_models.Event{
			ID:          v.ID,
			UserID:      v.UserID,
			Image:       v.Image,
			Title:       v.Title,
			Description: v.Description,
			Location:    v.Location,
			Date:        v.Date,
			Quota:       v.Quota,
		})
	}

	return eventResponseData, nil
}

func (r *queryResolver) GetEvent(ctx context.Context, id int) (*_models.Event, error) {
	responseData, err := r.eventRepository.GetById(id)
	if err != nil {
		return nil, errors.New("not found")
	}

	return &responseData, nil
}

func (r *queryResolver) GetEventKeyword(ctx context.Context, keyword string) ([]*_models.Event, error) {
	// fmt.Println("jalan", keyword)
	eventResponseData := []*_models.Event{}
	responseData, err := r.eventRepository.GetByKey(keyword)
	if err != nil {
		return nil, errors.New("not found")
	}
	// fmt.Println("respondata", responseData)
	for _, v := range responseData {
		eventResponseData = append(eventResponseData, &_models.Event{
			ID:          v.ID,
			UserID:      v.UserID,
			Image:       v.Image,
			Title:       v.Title,
			Description: v.Description,
			Location:    v.Location,
			Date:        v.Date,
			Quota:       v.Quota,
		})
	}

	return eventResponseData, nil
}

func (r *queryResolver) GetEventLocation(ctx context.Context, location string) ([]*_models.Event, error) {
	eventResponseData := []*_models.Event{}
	responseData, err := r.eventRepository.GetByLocation(location)
	if err != nil {
		return nil, errors.New("not found")
	}

	if err != nil {
		return nil, errors.New("not found")
	}

	for _, v := range responseData {
		eventResponseData = append(eventResponseData, &_models.Event{
			ID:          v.ID,
			UserID:      v.UserID,
			Image:       v.Image,
			Title:       v.Title,
			Description: v.Description,
			Location:    v.Location,
			Date:        v.Date,
			Quota:       v.Quota,
		})
	}

	return eventResponseData, nil
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
