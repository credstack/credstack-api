package management

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/api"
	"github.com/stevezaluk/credstack-api/middleware"
	"github.com/stevezaluk/credstack-lib/user"
)

/*
GetUserHandler - Provides a Fiber handler for processing a get request to /management/application. This should
not be called directly, and should only ever be passed to Fiber

TODO: Authentication handler needs to happen here
*/
func GetUserHandler(c fiber.Ctx) error {
	email := c.Query("email")

	requestedUser, err := user.GetUser(api.Server, email, false)
	if err != nil {
		return middleware.BindError(c, err)
	}

	return middleware.MarshalProtobuf(c, requestedUser)
}
