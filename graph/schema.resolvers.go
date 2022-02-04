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
	userData.Phone = input.Phone

	responseData, err := r.userRepository.Register(userData)
	return &responseData, err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *_model.EditUser) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}
	user, err := r.userRepository.Profile(userId.ID)
	if err != nil {
		return nil, errors.New("not found")
	}

	user.Name = *input.Name
	user.Email = *input.Email
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.MinCost)
	user.Password = string(passwordHash)
	user.Address = *input.Address
	user.Occupation = *input.Occupation
	user.Phone = *input.Phone

	updateErr := r.userRepository.UpdateUser(user)
	return &_model.Response{Code: 200, Message: "Update data Success", Success: true}, updateErr
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	user, err := r.userRepository.Profile(userId.ID)
	if err != nil {
		return nil, errors.New("not found")
	}

	deleteErr := r.userRepository.DeleteUser(user)
	return &_model.Response{Code: 200, Message: "Delete data Success", Success: true}, deleteErr
}

func (r *mutationResolver) CreateEvent(ctx context.Context, input *_model.NewEvent) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	fmt.Println("date", input.Date)
	eventData := _models.Event{}
	eventData.UserID = userId.ID
	eventData.Title = input.Title
	eventData.Image = input.Image
	eventData.Description = input.Description
	eventData.Location = input.Location
	eventData.Date = input.Date
	eventData.Quota = input.Quota

	createErr := r.eventRepository.CreateEvent(eventData)
	return &_model.Response{Code: 200, Message: "Create event Success", Success: true}, createErr
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id int, input *_model.EditEvent) (*_model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEvent(ctx context.Context, id int) (*_model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateParticipant(ctx context.Context, input *_model.NewParticipant) (*_model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateParticipant(ctx context.Context, id int, input *_model.EditParticipant) (*_model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteParticipant(ctx context.Context, id int) (*_model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *_model.NewComment) (*_model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id int, input *_model.EditComment) (*_model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*_model.Response, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*_model.LoginResponse, error) {
	user, err := r.userRepository.Login(email)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("not found")
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
		return nil, err
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

func (r *queryResolver) GetUsers(ctx context.Context) ([]*_models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (*_models.User, error) {
	responseData, err := r.userRepository.Profile(id)
	if err != nil {
		return nil, errors.New("not found")
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

func (r *queryResolver) GetOwnEvent(ctx context.Context) ([]*_models.Event, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return []*_models.Event{}, errors.New("unauthorized")
	}

	events := []*_models.Event{}

	responseData, err := r.eventRepository.GetOwnEvent(userId.ID)
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		events = append(events, &_models.Event{
			ID:          data.ID,
			UserID:      data.UserID,
			Image:       data.Image,
			Title:       data.Title,
			Description: data.Description,
			Location:    data.Location,
			Date:        data.Date,
			Quota:       data.Quota,
		})
	}

	return events, nil
}

func (r *queryResolver) GetEvents(ctx context.Context) ([]*_models.Event, error) {
	events := []*_models.Event{}

	responseData, err := r.eventRepository.GetEvents()
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		events = append(events, &_models.Event{
			ID:          data.ID,
			UserID:      data.UserID,
			Image:       data.Image,
			Title:       data.Title,
			Description: data.Description,
			Location:    data.Location,
			Date:        data.Date,
			Quota:       data.Quota,
		})
	}

	return events, nil
}

func (r *queryResolver) GetEvent(ctx context.Context, id int) (*_models.Event, error) {
	responseData, err := r.eventRepository.GetEvent(id)
	if err != nil {
		return nil, errors.New("not found")
	}

	return &responseData, nil
}

func (r *queryResolver) GetEventKeyword(ctx context.Context, search string) ([]*_models.Event, error) {
	events := []*_models.Event{}

	responseData, err := r.eventRepository.GetEventKeyword(search)
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		events = append(events, &_models.Event{
			ID:          data.ID,
			UserID:      data.UserID,
			Image:       data.Image,
			Title:       data.Title,
			Description: data.Description,
			Location:    data.Location,
			Date:        data.Date,
			Quota:       data.Quota,
		})
	}

	return events, nil
}

func (r *queryResolver) GetEventLocation(ctx context.Context, search string) ([]*_models.Event, error) {
	events := []*_models.Event{}

	responseData, err := r.eventRepository.GetEventLocation(search)
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		events = append(events, &_models.Event{
			ID:          data.ID,
			UserID:      data.UserID,
			Image:       data.Image,
			Title:       data.Title,
			Description: data.Description,
			Location:    data.Location,
			Date:        data.Date,
			Quota:       data.Quota,
		})
	}

	return events, nil
}

func (r *queryResolver) GetComments(ctx context.Context, eventID int) ([]*_models.Comment, error) {
	comments := []*_models.Comment{}

	responseData, err := r.commentRepository.GetComments(eventID)
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		comments = append(comments, &_models.Comment{
			ID:      data.ID,
			EventID: data.EventID,
			UserID:  data.UserID,
			Content: data.Content,
		})
	}

	return comments, nil
}

func (r *queryResolver) GetParticipants(ctx context.Context, eventID int) ([]*_models.Participant, error) {
	participants := []*_models.Participant{}

	responseData, err := r.participantRepository.GetParticipants(eventID)
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		participants = append(participants, &_models.Participant{
			ID:      data.ID,
			EventID: data.EventID,
			UserID:  data.UserID,
			Status:  data.Status,
		})
	}

	return participants, nil
}

// Mutation returns _generated.MutationResolver implementation.
func (r *Resolver) Mutation() _generated.MutationResolver { return &mutationResolver{r} }

// Query returns _generated.QueryResolver implementation.
func (r *Resolver) Query() _generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
