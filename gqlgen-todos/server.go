package main

import (
	"andrew.com/bff/BackRoundResolver"
	guuid "github.com/google/uuid"
	"log"
	"net/http"
	"os"

	"andrew.com/bff/graph"
	"andrew.com/bff/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8082"

/*to generate regenerate run*/
/*go run github.com/99designs/gqlgen*/
func main() {

	graph.BResolver = BackRoundResolver.NewBackroundResolver()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	go graph.BResolver.BroadcastEvent()
	var subs = BackRoundResolver.Subscriber{
		Id:     guuid.New().String(),
		Stop:   make(chan struct{}),
		Events: make(chan BackRoundResolver.IEvent),
	}
	graph.BResolver.MySubcriber <- subs
	go func(subscriber BackRoundResolver.ISubscriber) {
		subs.HandleMessage()
	}(subs)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
