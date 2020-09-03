package openapi

import (
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client/op"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/go-openapi/swag"
)

// APIError returns model.Error with given code and msg.
func APIError(code int32, msg string) *model.Error {
	return &model.Error{
		Code:    swag.Int32(code),
		Message: swag.String(msg),
	}
}

// ErrPayload returns err.Payload or err for unknown errors.
func ErrPayload(err error) interface{} {
	switch errDefault := err.(type) {
	default:
		return err
	case *op.ListContactsDefault:
		return errDefault.Payload
	case *op.AddContactDefault:
		return errDefault.Payload
	}
}
