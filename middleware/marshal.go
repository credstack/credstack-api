package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

/*
MarshalProtobuf - Marshal's and sends a protobuf while populating any errors that occur during the process. This
function is required as by default protobuf emits empty fields
*/
func MarshalProtobuf(c fiber.Ctx, message proto.Message) error {
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}

	data, err := marshaler.Marshal(message)
	if err != nil {
		wrappedErr := fmt.Errorf("%w (%v)", ErrFailedToBindResponse, err)
		return HandleError(c, wrappedErr)
	}

	c.Set("Content-Type", "application/json; charset=utf-8")
	return c.Status(200).Send(data)
}
