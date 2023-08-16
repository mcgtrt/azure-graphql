package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mcgtrt/azure-graphql/graph"
)

var httpListenAddr = "3000"

func main() {
	resolver, err := graph.NewResolver()
	if err != nil {
		panic(err)
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Comment out to turn off the playground
	http.Handle("/", playground.Handler("GraphQL playground", "/employee"))
	http.Handle("/employee", srv)

	log.Printf("Starting HTTP server at port: %s", httpListenAddr)
	log.Fatal(http.ListenAndServe(":"+httpListenAddr, nil))
}

func init() {
	if port := os.Getenv("HTTP_LISTEN_ADDR"); port != "" {
		httpListenAddr = port
	}
}
