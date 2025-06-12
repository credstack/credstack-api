package management

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/middleware"
	"github.com/stevezaluk/credstack-api/server"
	"github.com/stevezaluk/credstack-lib/api"
	apiModel "github.com/stevezaluk/credstack-lib/proto/api"
)

/*
GetAPIHandler - Provides a Fiber handler for processing a GET request to /management/api. This should
not be called directly, and should only ever be passed to Fiber

TODO: Authentication handler needs to happen here
*/
func GetAPIHandler(c fiber.Ctx) error {
	audience := c.Query("audience")

	requestedApi, err := api.GetAPI(server.Server, audience)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return middleware.MarshalProtobuf(c, requestedApi)
}

/*
PostAPIHandler - Provides a Fiber handler for processing a POST request to /management/api. This should
not be called directly, and should only ever be passed to Fiber

TODO: Authentication handler needs to happen here
TODO: Underlying functions need domain validation in place
TODO: Underlying functions need to be updated here so that we can assign applications at birth
*/
func PostAPIHandler(c fiber.Ctx) error {
	var model apiModel.API

	err := middleware.BindJSON(c, &model)
	if err != nil {
		return err
	}

	err = api.NewAPI(server.Server, model.Name, model.Domain, model.TokenType)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.Status(201).JSON(&fiber.Map{"message": "Created API successfully"})
}

/*
PatchAPIHandler - Provides a Fiber handler for processing a PATCH request to /management/api. This should
not be called directly, and should only ever be passed to Fiber

TODO: Authentication handler needs to happen here
*/
func PatchAPIHandler(c fiber.Ctx) error {
	domain := c.Query("audience")

	var model apiModel.API

	err := middleware.BindJSON(c, &model)
	if err != nil {
		return err
	}

	err = api.UpdateAPI(server.Server, domain, &model)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.Status(201).JSON(&fiber.Map{"message": "Updated API successfully"})
}

/*
DeleteAPIHandler - Provides a Fiber handler for processing a DELETE request to /management/api. This should
not be called directly, and should only ever be passed to Fiber

TODO: Authentication handler needs to happen here
*/
func DeleteAPIHandler(c fiber.Ctx) error {
	domain := c.Query("audience")

	err := api.DeleteAPI(server.Server, domain)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.Status(201).JSON(&fiber.Map{"message": "Deleted API successfully"})
}
