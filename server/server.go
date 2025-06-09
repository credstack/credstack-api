package server

import "github.com/stevezaluk/credstack-lib/server"

var Server *server.Server

/*
InitServer - Initializes server side resources
*/
func InitServer() error {
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

	return nil
}

/*
CloseServer - Gracefully terminates server side resources
*/
func CloseServer() error {
	Server.Log().LogDatabaseEvent("DatabaseDisconnect",
		Server.Database().Options().Hostname,
		int(Server.Database().Options().Port),
	)
	/*
		Then we close our connection to the database gracefully.
	*/
	err := Server.Database().Disconnect()
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
		if err.Error() == "sync /dev/stdout: invalid argument" { // we explicitly ignore this error as it /dev/stdout is not a real file that supports the Sync system call
			return nil
		}
		return err
	}

	return nil
}
