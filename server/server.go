package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/chadbohannan/go-gqlgen-sample"
)

const defaultPort = "8080"

var stringList = []string{
	"first string",
	"second string",
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground(
		"GraphQL playground",
		"/query"))

	http.Handle("/query",
		handler.GraphQL(
			go_gqlgen_sample.NewExecutableSchema(
				go_gqlgen_sample.Config{
					Resolvers: &go_gqlgen_sample.Resolver{
						Data: stringList,
					},
				},
			),
		),
	)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}