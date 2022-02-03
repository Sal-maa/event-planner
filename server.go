package main

import (
	"log"
	"net/http"

	handler "github.com/99designs/gqlgen/graphql/handler"
	playground "github.com/99designs/gqlgen/graphql/playground"
	chi "github.com/go-chi/chi"

	_config "github.com/justjundana/event-planner/config"
	_graph "github.com/justjundana/event-planner/graph"
	_generated "github.com/justjundana/event-planner/graph/generated"
	_commentRepository "github.com/justjundana/event-planner/repository/comment"
	_eventRepository "github.com/justjundana/event-planner/repository/event"
	_participantRepository "github.com/justjundana/event-planner/repository/participant"
	_userRepository "github.com/justjundana/event-planner/repository/user"
	cors "github.com/rs/cors"
)

func main() {
	db := _config.FetchConnection()

	router := chi.NewRouter()
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	router.Use(corsOptions.Handler)
	userRepo := _userRepository.New(db)
	evenRepo := _eventRepository.New(db)
	participantRepo := _participantRepository.New(db)
	commentRepo := _commentRepository.New(db)

	client := _graph.NewResolver(userRepo, evenRepo, participantRepo, commentRepo)
	srv := handler.NewDefaultServer(_generated.NewExecutableSchema(_generated.Config{Resolvers: client}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
