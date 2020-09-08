//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "ListContacts=AddContact"

package openapi

import (
	"net/http"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/restapi/op"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/go-openapi/swag"
)

func errListContacts(log Log, err error, code errCode) op.ListContactsResponder {
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError { // Do no expose details about internal errors.
		msg = "internal error"
	}

	return op.NewListContactsDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(code.extra),
		Message: swag.String(msg),
	})
}
