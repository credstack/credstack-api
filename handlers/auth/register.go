package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/middleware"
	"github.com/stevezaluk/credstack-api/server"
	"github.com/stevezaluk/credstack-lib/options"
	"github.com/stevezaluk/credstack-lib/proto/request"
	"github.com/stevezaluk/credstack-lib/user"
)

/*
RegisterUserHandler - Provides a fiber handler for processing a POST request to /auth/register This should
not be called directly, and should only ever be passed to fiber

TODO: Authentication handler needs to happen here
*/
func RegisterUserHandler(c fiber.Ctx) error {
	var registerRequest request.UserRegisterRequest

	err := c.Bind().JSON(&registerRequest)
	if err != nil {
		wrappedErr := fmt.Errorf("%w (%v)", middleware.ErrFailedToBindResponse, err)
		return middleware.HandleError(c, wrappedErr)
	}

	err = user.RegisterUser(
		server.Server,
		options.Credential().FromConfig(),
		registerRequest.Email,
		registerRequest.Username,
		registerRequest.Password,
	)

	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{"message": "User successfully registered"})
}
