package management

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/api"
	"github.com/stevezaluk/credstack-api/api/handlers/middleware"
	"github.com/stevezaluk/credstack-lib/application"
)

/*
GetApplicationHandler - Provides a Fiber handler for processing a get request to /management/application. This should
not be called directly, and should only ever be passed to Fiber

TODO: Authentication handler needs to happen here
*/
func GetApplicationHandler(c fiber.Ctx) error {
	clientId := c.Query("client_id")

	app, err := application.GetApplication(api.Server, clientId, true)
	if err != nil {
		return middleware.BindError(c, err)
	}

	return c.Status(200).JSON(&app)
}
