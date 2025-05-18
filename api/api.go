package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-lib/server"
)

/*
API - Provides a core abstraction of the API itself. Primarily handles dependency injection and graceful termination
of the API
*/
type API struct {
	// server - The server structure that provides dependencies for the API
	server *server.Server

	// app - The primary fiber application that handles API level routing
	app *fiber.App
}

/*
New - A constructor for the API. This will use the server.Server structure passed in its parameter, and initialize a
new fiber.App for you. Note: The server structure provided here should not have a pre-connected database, this structure
will handle server.Database lifecycle for you
*/
func New(server *server.Server) *API {
	return &API{
		server: server,
		app:    fiber.New(),
	}
}

/*
FromConfig - Initializes the server.Server structure from configuration values provided by viper. Note: The server
structure provided here should not have a pre-connected database, this structure will handle server.Database lifecycle
for you
*/
func FromConfig() *API {
	return &API{
		server: server.FromConfig(),
		app:    fiber.New(),
	}
}
