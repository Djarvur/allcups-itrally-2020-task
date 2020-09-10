//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "GetBalance=Cash,Dig,ExploreArea,IssueLicense,ListLicenses"

package openapi

import (
	"net/http"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/restapi/op"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/go-openapi/swag"
)

func errGetBalance(log Log, err error, code errCode) op.GetBalanceResponder {
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError { // Do no expose details about internal errors.
		msg = "internal error" //nolint:goconst // Duplicated by go:generate.
	}

	return op.NewGetBalanceDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(code.extra),
		Message: swag.String(msg),
	})
}
