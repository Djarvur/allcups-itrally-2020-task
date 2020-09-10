package openapi

import (
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/restapi/op"
)

func (srv *server) getBalance(params op.GetBalanceParams) op.GetBalanceResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	wallet, err := srv.app.Balance(ctx)
	switch {
	default:
		return errGetBalance(log, err, codeInternal)
	case err == nil:
		return op.NewGetBalanceOK().WithPayload(apiWallet(wallet))
	}
}
