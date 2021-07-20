package main

import (
	"log"
	"myapp/graph"
	"myapp/graph/generated"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	// "myapp/config"
	"myapp/dataloader"

	"github.com/go-chi/chi"
)

const defaultPort = "8081"

func main() {
	// config.MigrateUp()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router := chi.NewRouter()
	router.Use(dataloader.DataloaderMiddleware)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
