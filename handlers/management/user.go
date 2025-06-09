package management

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/api"
	"github.com/stevezaluk/credstack-api/middleware"
	userModel "github.com/stevezaluk/credstack-lib/proto/user"
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
PatchUserHandler - Provides a Fiber handler for processing a PATCH request to /management/user. This should
not be called directly, and should only ever be passed to Fiber

TODO: Authentication handler needs to happen here
*/
func PatchUserHandler(c fiber.Ctx) error {
	email := c.Query("email")

	var model userModel.User

	err := c.Bind().JSON(&model)
	if err != nil {
		wrappedErr := fmt.Errorf("%w (%v)", middleware.ErrFailedToBindResponse, err)
		return middleware.BindError(c, wrappedErr)
	}

	err = user.UpdateUser(api.Server, email, &model)
	if err != nil {
		return middleware.BindError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{"message": "Updated user successfully"})
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
