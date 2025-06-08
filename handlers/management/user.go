package management

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/api"
	"github.com/stevezaluk/credstack-api/middleware"
	"github.com/stevezaluk/credstack-lib/user"
)

/*
GetUserHandler - Provides a Fiber handler for processing a get request to /management/user. This should
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

/*
DeleteUserHandler - Provides a Fiber handler for processing a DELETE request to /management/user. This should
not be called directly, and should only ever be passed to Fiber

TODO: Authentication handler needs to happen here
*/
func DeleteUserHandler(c fiber.Ctx) error {
	email := c.Query("email")

	err := user.DeleteUser(api.Server, email)
	if err != nil {
		return middleware.BindError(c, err)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Successfully deleted user"})
}
