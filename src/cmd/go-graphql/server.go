package main

import (
	"log"
	"net/http"
	"os"

	"github.com/raghanag/my-project/pkg/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/raghanag/my-project/cmd/go-graphql/graph"
	"github.com/raghanag/my-project/cmd/go-graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	dbUserName := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbURL := os.Getenv("POSTGRES_URL")
	dbName := os.Getenv("POSTGRES_DB")

	if port == "" {
		port = defaultPort
	}

	ingestService := &postgres.IngestService{
		DbUserName: dbUserName,
		DbPassword: dbPassword,
		DbURL:      dbURL,
		DbName:     dbName,
	}
	err := ingestService.Initialise()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{Ingest: ingestService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
