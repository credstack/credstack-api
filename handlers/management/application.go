package management

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/api"
	"github.com/stevezaluk/credstack-api/middleware"
	"github.com/stevezaluk/credstack-lib/application"
	applicationModel "github.com/stevezaluk/credstack-lib/proto/application"
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
		return middleware.HandleError(c, err)
	}

	return middleware.MarshalProtobuf(c, app)
}

/*
PostApplicationHandler - Provides a fiber handler for processing a POST request to /management/application This should
not be called directly, and should only ever be passed to fiber

TODO: Authentication handler needs to happen here
TODO: Add client_id in return for new application
*/
func PostApplicationHandler(c fiber.Ctx) error {
	var model applicationModel.Application

	err := c.Bind().JSON(&model)
	if err != nil {
		wrappedErr := fmt.Errorf("%w (%v)", middleware.ErrFailedToBindResponse, err)
		return middleware.HandleError(c, wrappedErr)
	}

	clientId, err := application.NewApplication(api.Server, model.Name, model.RedirectUri, model.GrantType...)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.Status(201).JSON(&fiber.Map{"message": "Created application successfully", "client_id": clientId})
}

/*
PatchApplicationHandler - Provides a fiber handler for processing a PATCH request to /management/application This should
not be called directly, and should only ever be passed to fiber

TODO: Authentication handler needs to happen here
*/
func PatchApplicationHandler(c fiber.Ctx) error {
	clientId := c.Query("client_id")

	var model applicationModel.Application

	err := c.Bind().JSON(&model)
	if err != nil {
		wrappedErr := fmt.Errorf("%w (%v)", middleware.ErrFailedToBindResponse, err)
		return middleware.HandleError(c, wrappedErr)
	}

	err = application.UpdateApplication(api.Server, clientId, &model)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{"message": "Updated application successfully"})
}

/*
DeleteApplicationHandler - Provides a fiber handler for processing a DELETE request to /management/application This should
not be called directly, and should only ever be passed to fiber

TODO: Authentication handler needs to happen here
*/
func DeleteApplicationHandler(c fiber.Ctx) error {
	clientId := c.Query("client_id")

	err := application.DeleteApplication(api.Server, clientId)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{"message": "Deleted application successfully"})
}
