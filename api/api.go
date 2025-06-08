package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-lib/server"
	"strconv"
)

// Server - A global variable that all API handlers can use for interacting with server side resources
var Server *server.Server

// App - A global variable that provides interaction with the Fiber Application
var App *fiber.App

/*
New - Constructs a new fiber.App with recommended configurations
*/
func New() *fiber.App {
	/*
		Realistically, these should probably be exposed to the user for them to modify,
		however they are hardcoded for now to ensure that these will ensure the most performance
	*/
	config := fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "CredStack API",
	}

	app := fiber.New(config)

	return app
}

/*
Start - Connects to MongoDB and starts the API
*/
func Start(port int) error {
	Server = server.FromConfig()

	Server.Log().LogDatabaseEvent("DatabaseConnect",
		Server.Database().Options().Hostname,
		int(Server.Database().Options().Port),
	)

	/*
		We still need to connect to our database as the constructors for Server do not
		provide this functionality by default.
	*/
	err := Server.Database().Connect()
	if err != nil {
		Server.Log().LogErrorEvent("Failed to connect to database", err)
		return err
	}

	/*
		Realistically, these should probably be exposed to the user for them to modify,
		however they are hardcoded for now to ensure that these will ensure the most performance
	*/
	listenConfig := fiber.ListenConfig{
		DisableStartupMessage: true,
		EnablePrefork:         false, // this makes log entries duplicate
		ListenerNetwork:       "tcp4",
	}

	/*
		Once our database is connected we can properly start our API
	*/
	Server.Log().LogStartupEvent("API", "API is now listening for requests on port "+strconv.Itoa(port))
	err = App.Listen(":"+strconv.Itoa(port), listenConfig)
	if err != nil {
		return err // log here
	}

	return nil
}

/*
Stop - Gracefully terminates the API, closes database connections and flushes existing logs to sync
*/
func Stop() error {
	Server.Log().LogShutdownEvent("API", "Shutting down API. New requests will not be allowed")

	/*
		First we shut down the API to ensure that any currently processing requests
		finish. Additionally, we don't want new requests coming in as we are shutting
		down the server
	*/
	err := App.Shutdown()
	if err != nil {
		return err // log here
	}

	Server.Log().LogDatabaseEvent("DatabaseDisconnect",
		Server.Database().Options().Hostname,
		int(Server.Database().Options().Port),
	)
	/*
		Then we close our connection to the database gracefully.
	*/
	err = Server.Database().Disconnect()
	if err != nil {
		return err // log here
	}

	Server.Log().LogShutdownEvent("LogFlush", "Flushing queued logs and closing log file")
	/*
		Then we flush any buffered logs to sync and close the open log file, any errors
		returned from this action will be logged properly
	*/
	err = Server.Log().CloseLog()
	if err != nil {
		return err
	}

	return nil
}
