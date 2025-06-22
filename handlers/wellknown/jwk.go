package wellknown

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/middleware"
	"github.com/stevezaluk/credstack-api/server"
	"github.com/stevezaluk/credstack-lib/key"
)

/*
GetJWKHandler - Provides a Fiber handler for processing a GET request to /.well-known/jwks.json. This should
not be called directly, and should only ever be passed to Fiber
*/
func GetJWKHandler(c fiber.Ctx) error {
	jwks, err := key.GetJWKS(server.Server)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return middleware.MarshalProtobuf(c, jwks)
}
