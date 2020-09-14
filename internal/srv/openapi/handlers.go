package openapi

import (
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/restapi/op"
	"github.com/go-openapi/swag"
)

func (srv *server) getBalance(params op.GetBalanceParams) op.GetBalanceResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	balance, wallet, err := srv.app.Balance(ctx)
	switch {
	default:
		return errGetBalance(log, err, codeInternal)
	case err == nil:
		return op.NewGetBalanceOK().WithPayload(&model.Balance{
			Balance: swag.Uint32(uint32(balance)),
			Wallet:  apiWallet(wallet),
		})
	}
}
