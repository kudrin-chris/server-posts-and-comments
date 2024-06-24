package main

import (
	"log"
	"net/http"
	"os"
	"server/graph"
	"server/internal/db"
	"server/internal/db/postgres"
	"server/internal/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	optionsDb := db.PostgresOptions{
		Name:     os.Getenv("POSTGRES_DBNAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	postgresDb, _ := db.NewPostgresDB(optionsDb)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	posts := postgres.NewPostsPostgres(postgresDb)
	comments := postgres.NewCommentsPostgres(postgresDb)
	services := service.NewServices(posts, comments)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		PostsService:    services.Posts,
		CommentsService: services.Comments,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
