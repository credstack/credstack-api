package auth

import (
	"github.com/credstack/credstack-api/middleware"
	"github.com/credstack/credstack-api/server"
	"github.com/credstack/credstack-lib/options"
	"github.com/credstack/credstack-lib/user"
	"github.com/credstack/credstack-models/proto/request"
	"github.com/gofiber/fiber/v3"
)

/*
RegisterUserHandler - Provides a fiber handler for processing a POST request to /auth/register This should
not be called directly, and should only ever be passed to fiber

TODO: Authentication handler needs to happen here
*/
func RegisterUserHandler(c fiber.Ctx) error {
	var registerRequest request.UserRegisterRequest

	err := middleware.BindJSON(c, &registerRequest)
	if err != nil {
		return err
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
