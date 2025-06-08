package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	credstackErrors "github.com/stevezaluk/credstack-lib/errors" // this needs to be fixed
)

/*
BindError - Takes a CredStack error and marshal's it into a JSON response
*/
func BindError(c fiber.Ctx, err error) error {
	var casted credstackErrors.CredstackError

	if !errors.As(err, &casted) {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(casted.HTTPStatusCode).JSON(fiber.Map{"errorCode": casted.Short(), "error": casted.Error()})
}
