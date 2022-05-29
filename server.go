package main

import (
	"log"
	"net/http"
	"os"

	"acy.com/gqlgendemo/database"
	"acy.com/gqlgendemo/graph"
	"acy.com/gqlgendemo/graph/generated"
	"acy.com/gqlgendemo/repository"
	"acy.com/gqlgendemo/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const defaultPort = "7000"

func main() {
	godotenv.Load()

	config := &database.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        Password: os.Getenv("DB_PASS"),
        User:     os.Getenv("DB_USER"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
        DBName:   os.Getenv("DB_NAME"),
		Schema:   os.Getenv("DB_SCHEMA"),
    }

	db, err := database.NewConnection(config)
    if err != nil {
        panic(err)
    }

	database.Migrate(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		BookService: service.NewBookService(repository.NewBookRepository(db)),
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)
	router.Handle("/graphql", server)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
