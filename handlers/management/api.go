package management

import (
	"github.com/gofiber/fiber/v3"
	server "github.com/stevezaluk/credstack-api/api"
	"github.com/stevezaluk/credstack-api/middleware"
	"github.com/stevezaluk/credstack-lib/api"
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
		return middleware.BindError(c, err)
	}

	return middleware.MarshalProtobuf(c, requestedApi)
}
