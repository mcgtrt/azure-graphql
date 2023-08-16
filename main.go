package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mcgtrt/azure-graphql/api"
	"github.com/mcgtrt/azure-graphql/graph"
	"github.com/mcgtrt/azure-graphql/store"
)

var httpListenAddr = "3000"

func main() {
	resolver, err := graph.NewResolver()
	if err != nil {
		panic(err)
	}

	store, err := store.NewAzureEmployeeStore()
	if err != nil {
		panic(err)
	}

	var (
		authHandler = api.NewAuthHandler(store)
		srv         = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/employee"))
	http.Handle("/employee", api.JWTAuthenticate(context.Background(), store, srv))
	http.HandleFunc("/login", authHandler.HandleAuth)

	log.Printf("Starting HTTP server at port: %s", httpListenAddr)
	log.Fatal(http.ListenAndServe(":"+httpListenAddr, nil))
}

func init() {
	if port := os.Getenv("HTTP_LISTEN_ADDR"); port != "" {
		httpListenAddr = port
	}
}
