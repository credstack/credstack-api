package oauth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stevezaluk/credstack-api/middleware"
	"github.com/stevezaluk/credstack-api/server"
	"github.com/stevezaluk/credstack-lib/oauth/token"
	"github.com/stevezaluk/credstack-lib/proto/request"
)

/*
GetTokenHandler - Provides a fiber handler for processing a GET request to /oauth2/token This should
not be called directly, and should only ever be passed to fiber
*/
func GetTokenHandler(c fiber.Ctx) error {
	req := new(request.TokenRequest)

	if err := c.Bind().Query(req); err != nil {
		return middleware.HandleError(c, err)
	}

	resp, err := token.IssueToken(server.Server, req, "") // not good, need a way of allowing the user to set the issuer
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return middleware.MarshalProtobuf(c, resp)
}
