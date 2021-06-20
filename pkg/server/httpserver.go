package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gigamono/gigamono-base-engine/internal/graphql"
)

func (server *BaseEngineServer) httpServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(":", server.Config.Services.BaseEngine.Ports.Public),
	)
	if err != nil {
		return err
	}

	server.setRoutes() // Set routes.

	// Use http server.
	httpServer := &http.Server{
		Handler: server.GinEngine,
	}

	return httpServer.Serve(listener)
}

func (server *BaseEngineServer) setRoutes() {
	graphqlHandler := graphql.Handler(&server.App)
	playgroundHandler := graphql.PlaygroundHandler()

	server.GinEngine.POST("/graphql", graphqlHandler)      // Handles all graphql requests.
	server.GinEngine.GET("/graphql", graphqlHandler)       // Handles query-only graphql requests.
	server.GinEngine.GET("/playground", playgroundHandler) // Shows playground UI.
}
