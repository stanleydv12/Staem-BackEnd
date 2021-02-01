package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/stanleydv12/gqlgen-todos/graph"
	"github.com/stanleydv12/gqlgen-todos/graph/generated"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
