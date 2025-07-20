package oauth

import (
	"github.com/credstack/credstack-api/middleware"
	"github.com/credstack/credstack-api/server"
	"github.com/credstack/credstack-lib/oauth/flow"
	"github.com/credstack/credstack-models/proto/request"
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
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

	resp, err := flow.IssueTokenForFlow(server.Server, req, viper.GetString("issuer"))
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return middleware.MarshalProtobuf(c, resp)
}
