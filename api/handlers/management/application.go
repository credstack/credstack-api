package management

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/api"
	"github.com/stevezaluk/credstack-api/api/handlers/middleware"
	"github.com/stevezaluk/credstack-lib/application"
	applicationModel "github.com/stevezaluk/credstack-lib/proto/application"
)

/*
GetApplicationHandler - Provides a Fiber handler for processing a get request to /management/application. This should
not be called directly, and should only ever be passed to Fiber

TODO: Authentication handler needs to happen here
TODO: Update handler to process withCredentials parameter
*/
func GetApplicationHandler(c fiber.Ctx) error {
	clientId := c.Query("client_id")

	app, err := application.GetApplication(api.Server, clientId, true)
	if err != nil {
		return middleware.BindError(c, err)
	}

	return c.Status(200).JSON(&app)
}

/*
PostApplicationHandler - Provides a fiber handler for processing a POST request to /mangement/application This should
not be called directly, and should only ever be passed to fiber

TODO: Authentication handler needs to happen here
TODO: Update NewApplication to accept an optional list of request URI's
*/
func PostApplicationHandler(c fiber.Ctx) error {
	var model applicationModel.Application

	err := c.Bind().JSON(&model)
	if err != nil {
		wrappedErr := fmt.Errorf("%w (%v)", middleware.ErrFailedToBindResponse, err)
		return middleware.BindError(c, wrappedErr)
	}

	err = application.NewApplication(api.Server, model.Name, model.GrantType)
	if err != nil {
		return middleware.BindError(c, err)
	}

	return c.Status(201).JSON(&fiber.Map{"message": "Created application successfully"})
}
