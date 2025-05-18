package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-lib/server"
	"strconv"
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
Start - Connects to MongoDB and starts the API
*/
func (api *API) Start(port int) error {
	/*
		Realistically, these should probably be exposed to the user for them to modify,
		however they are hardcoded for now to ensure that these will ensure the most performance
	*/
	config := fiber.ListenConfig{
		DisableStartupMessage: true,
		EnablePrefork:         false, // this makes log entries duplicate
		ListenerNetwork:       "tcp4",
	}

	api.server.Log().LogDatabaseEvent("DatabaseConnect",
		api.server.Database().Options().Hostname,
		int(api.server.Database().Options().Port),
	)

	/*
		We still need to connect to our database as the constructors for Server do not
		provide this functionality by default.
	*/
	err := api.server.Database().Connect()
	if err != nil {
		api.server.Log().LogErrorEvent("Failed to connect to database", err)
		return err
	}

	/*
		Once our database is connected we can properly start our API
	*/
	api.server.Log().LogStartupEvent("Listener", "API is now listening for requests on port "+strconv.Itoa(port))
	err = api.app.Listen(":"+strconv.Itoa(port), config)
	if err != nil {
		return err
	}

	return nil
}

/*
Stop - Gracefully terminates the API, closes database connections and flushes existing logs to sync
*/
func (api *API) Stop() error {
	api.server.Log().LogShutdownEvent("Listener", "Shutting down API. New requests will not be allowed")

	/*
		First we shut down the API to ensure that any currently processing requests
		finish. Additionally, we don't want new requests coming in as we are shutting
		down the server
	*/
	err := api.app.Shutdown()
	if err != nil {
		return err
	}

	api.server.Log().LogDatabaseEvent("DatabaseDisconnect",
		api.server.Database().Options().Hostname,
		int(api.server.Database().Options().Port),
	)
	/*
		Then we close our connection to the database gracefully.
	*/
	err = api.server.Database().Disconnect()
	if err != nil {
		return err
	}

	api.server.Log().LogShutdownEvent("LogFlush", "Flushing queued logs and closing log file")
	/*
		Then we flush any buffered logs to sync and close the open log file, any errors
		returned from this action will be logged properly
	*/
	err = api.server.Log().CloseLog()
	if err != nil {
		return err
	}

	return nil
}

/*
New - A constructor for the API. This will use the server.Server structure passed in its parameter, and initialize a
new fiber.App for you. Note: The server structure provided here should not have a pre-connected database, this structure
will handle server.Database lifecycle for you
*/
func New(server *server.Server) *API {
	/*
		Realistically, these should probably be exposed to the user for them to modify,
		however they are hardcoded for now to ensure that these will ensure the most performance
	*/
	config := fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "CredStack API",
	}

	return &API{
		server: server,
		app:    fiber.New(config),
	}
}

/*
FromConfig - Initializes the server.Server structure from configuration values provided by viper. Note: The server
structure provided here should not have a pre-connected database, this structure will handle server.Database lifecycle
for you
*/
func FromConfig() *API {
	serv := server.FromConfig()
	return New(serv)
}
