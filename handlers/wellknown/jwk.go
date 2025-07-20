package wellknown

import (
	"github.com/credstack/credstack-api/middleware"
	"github.com/credstack/credstack-api/server"
	"github.com/credstack/credstack-lib/oauth/jwk"
	"github.com/gofiber/fiber/v3"
)

/*
GetJWKHandler - Provides a Fiber handler for processing a GET request to /.well-known/jwks.json. This should
not be called directly, and should only ever be passed to Fiber
*/
func GetJWKHandler(c fiber.Ctx) error {
	jwks, err := jwk.GetJWKS(server.Server)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return middleware.MarshalProtobuf(c, jwks)
}
