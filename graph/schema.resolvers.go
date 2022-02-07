package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

	checkEmail, errCheck := r.userRepository.CheckEmail(userData)
	if errCheck != nil {
		fmt.Println("errCheck", errCheck)
		return nil, errCheck
	}

	if checkEmail.Email == userData.Email {
		return nil, errors.New("email is already exist")
	}

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

	if input.Avatar != nil {
		user.Avatar = *input.Avatar
	}
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Password != nil {
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.MinCost)
		user.Password = string(passwordHash)
	}
	if input.Address != nil {
		user.Address = *input.Address
	}
	if input.Occupation != nil {
		user.Occupation = *input.Occupation
	}
	if input.Phone != nil {
		user.Phone = *input.Phone
	}

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

	eventData := _models.Event{}
	eventData.UserID = userId.ID
	eventData.Title = input.Title
	eventData.CategoryId = input.CategoryID
	eventData.Image = input.Image
	eventData.Description = input.Description
	eventData.Location = input.Location
	eventData.Date = input.Date
	eventData.Quota = input.Quota

	createErr := r.eventRepository.CreateEvent(eventData)
	return &_model.Response{Code: 200, Message: "Create event Success", Success: true}, createErr
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id int, input *_model.EditEvent) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	event, err := r.eventRepository.GetEvent(id)
	if err != nil {
		return nil, errors.New("not found")
	}

	if event.UserID != userId.ID {
		return &_model.Response{Code: http.StatusForbidden, Message: "you don't have permission to update this event", Success: false}, nil
	}

	if input.Image != nil {
		event.Image = *input.Image
	}
	if input.Title != nil {
		event.Title = *input.Title
	}
	if input.CategoryID != nil {
		event.CategoryId = *input.CategoryID
	}
	if input.Description != nil {
		event.Description = *input.Description
	}
	if input.Location != nil {
		event.Location = *input.Location
	}
	if input.Date != nil {
		event.Date = *input.Date
	}
	if input.Quota != nil {
		event.Quota = *input.Quota
	}

	updateErr := r.eventRepository.UpdateEvent(event)
	return &_model.Response{Code: http.StatusOK, Message: "Update event Success", Success: true}, updateErr
}

func (r *mutationResolver) DeleteEvent(ctx context.Context, id int) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	event, err := r.eventRepository.GetEvent(id)
	if err != nil {
		return nil, errors.New("not found")
	}

	if event.UserID != userId.ID {
		return &_model.Response{Code: http.StatusForbidden, Message: "you don't have permission to delete this event", Success: false}, nil
	}

	deleteErr := r.eventRepository.DeleteEvent(event)
	return &_model.Response{Code: http.StatusOK, Message: "Delete event Success", Success: true}, deleteErr
}

func (r *mutationResolver) CreateParticipant(ctx context.Context, input *_model.NewParticipant) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	checkParticipant, _ := r.participantRepository.CheckParticipant(userId.ID, input.EventID)
	if checkParticipant.UserID > 0 {
		return &_model.Response{Code: http.StatusForbidden, Message: "you have been registered in this event", Success: false}, nil
	}

	participant := _models.Participant{}
	participant.EventID = input.EventID
	participant.UserID = userId.ID
	participant.Status = input.Status

	createErr := r.participantRepository.CreateParticipant(participant)
	return &_model.Response{Code: http.StatusOK, Message: "You have been successfully registered in this event", Success: true}, createErr
}

func (r *mutationResolver) UpdateParticipant(ctx context.Context, id int, input *_model.EditParticipant) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	participant, err := r.participantRepository.GetParticipant(id)
	if err != nil {
		return nil, errors.New("not found")
	}

	if participant.UserID != userId.ID {
		return &_model.Response{Code: http.StatusForbidden, Message: "you don't have permission to update this event", Success: false}, nil
	}

	participant.Status = *input.Status

	updateErr := r.participantRepository.UpdateParticipant(participant)
	return &_model.Response{Code: http.StatusOK, Message: "Update participant status success", Success: true}, updateErr
}

func (r *mutationResolver) DeleteParticipant(ctx context.Context, eventID int) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	participant := _models.Participant{}
	participant.EventID = eventID
	participant.UserID = userId.ID

	deleteErr := r.participantRepository.DeleteParticipant(participant)
	return &_model.Response{Code: http.StatusOK, Message: "Delete participant Success", Success: true}, deleteErr
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *_model.NewComment) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	comment := _models.Comment{}
	comment.EventID = input.EventID
	comment.UserID = userId.ID
	comment.Content = input.Content

	createErr := r.commentRepository.CreateComment(comment)
	return &_model.Response{Code: http.StatusOK, Message: "Create comment Success", Success: true}, createErr
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id int, input *_model.EditComment) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	comment, err := r.commentRepository.GetComment(id)
	if err != nil {
		return nil, errors.New("not found")
	}

	if comment.UserID != userId.ID {
		return &_model.Response{Code: http.StatusForbidden, Message: "you don't have permission to update this comment", Success: false}, nil
	}

	if input.Content != nil {
		comment.Content = *input.Content
	}

	updateErr := r.commentRepository.UpdateComment(comment)
	return &_model.Response{Code: http.StatusOK, Message: "Update comment Success", Success: true}, updateErr
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*_model.Response, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return &_model.Response{}, errors.New("unauthorized")
	}

	comment, err := r.commentRepository.GetComment(id)
	if err != nil {
		return nil, errors.New("not found")
	}

	if comment.UserID != userId.ID {
		return &_model.Response{Code: http.StatusForbidden, Message: "you don't have permission to update this comment", Success: false}, nil
	}

	deleteErr := r.commentRepository.DeleteComment(comment)
	return &_model.Response{Code: http.StatusOK, Message: "Delete comment Success", Success: true}, deleteErr
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*_model.LoginResponse, error) {
	user, err := r.userRepository.Login(email)
	if err != nil {
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
	users := []*_models.User{}

	responseData, err := r.userRepository.GetUsers()
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		users = append(users, &_models.User{
			ID:         data.ID,
			Name:       data.Name,
			Email:      data.Email,
			Password:   data.Password,
			Address:    data.Address,
			Occupation: data.Occupation,
			Phone:      data.Phone,
		})
	}

	return users, nil
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
			CategoryId:  data.CategoryId,
			Description: data.Description,
			Location:    data.Location,
			Date:        data.Date,
			Quota:       data.Quota,
		})
	}

	return events, nil
}

func (r *queryResolver) GetParticipateEvent(ctx context.Context) ([]*_models.Event, error) {
	userId := _middleware.ForContext(ctx)
	if userId == nil {
		return []*_models.Event{}, errors.New("unauthorized")
	}

	events := []*_models.Event{}

	responseData, err := r.eventRepository.GetParticipateEvent(userId.ID)
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		events = append(events, &_models.Event{
			ID:          data.ID,
			Image:       data.Image,
			Title:       data.Title,
			CategoryId:  data.CategoryId,
			Description: data.Description,
			Location:    data.Location,
			Date:        data.Date,
			Quota:       data.Quota,
		})
	}

	return events, nil
}

func (r *queryResolver) GetPaginationEvents(ctx context.Context, limit *int, offset *int) ([]*_models.Event, error) {
	events := []*_models.Event{}

	responseData, err := r.eventRepository.Pagination(limit, offset)
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		events = append(events, &_models.Event{
			ID:          data.ID,
			UserID:      data.UserID,
			Image:       data.Image,
			Title:       data.Title,
			CategoryId:  data.CategoryId,
			Description: data.Description,
			Location:    data.Location,
			Date:        data.Date,
			Quota:       data.Quota,
		})
	}

	return events, nil
}

func (r *queryResolver) GetJoinableEvents(ctx context.Context) ([]*_models.Event, error) {
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
			CategoryId:  data.CategoryId,
			Description: data.Description,
			Location:    data.Location,
			Date:        data.Date,
			Quota:       data.Quota,
		})
	}

	return events, nil
}

func (r *queryResolver) GetEvent(ctx context.Context, id int) (*_models.Event, error) {
	event, err := r.eventRepository.GetEvent(id)
	if err != nil {
		return nil, errors.New("not found")
	}

	responseData := _models.Event{
		ID:          event.ID,
		UserID:      event.UserID,
		Image:       event.Image,
		Title:       event.Title,
		CategoryId:  event.CategoryId,
		Description: event.Description,
		Location:    event.Location,
		Date:        event.Date,
		Quota:       event.Quota,
		User: _models.User{
			ID:         event.User.ID,
			Avatar:     event.User.Avatar,
			Name:       event.User.Name,
			Email:      event.User.Email,
			Address:    event.User.Address,
			Occupation: event.User.Occupation,
			Phone:      event.User.Phone,
		},
	}

	comments, err := r.commentRepository.GetComments(id)
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, comment := range comments {
		comment := _models.Comment{
			ID:      comment.ID,
			EventID: comment.EventID,
			UserID:  comment.UserID,
			User:    comment.User,
			Content: comment.Content,
		}

		responseData.Comments = append(responseData.Comments, comment)
	}

	return &responseData, nil
}

func (r *queryResolver) GetEventsBySearch(ctx context.Context, search string) ([]*_models.Event, error) {
	events := []*_models.Event{}

	responseData, err := r.eventRepository.SearchEvents(search)
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		events = append(events, &_models.Event{
			ID:          data.ID,
			UserID:      data.UserID,
			Image:       data.Image,
			Title:       data.Title,
			CategoryId:  data.CategoryId,
			Description: data.Description,
			Location:    data.Location,
			Date:        data.Date,
			Quota:       data.Quota,
		})
	}

	return events, nil
}

func (r *queryResolver) GetEventMostAttendant(ctx context.Context) ([]*_models.Event, error) {
	events := []*_models.Event{}

	responseData, err := r.eventRepository.GetEventMostAttendant()
	if err != nil {
		return nil, errors.New("not found")
	}

	for _, data := range responseData {
		events = append(events, &_models.Event{
			ID:          data.ID,
			UserID:      data.UserID,
			Image:       data.Image,
			Title:       data.Title,
			CategoryId:  data.CategoryId,
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
			User:    data.User,
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
			User:    data.User,
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
