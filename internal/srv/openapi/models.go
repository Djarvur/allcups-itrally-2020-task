package openapi

import (
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
)

func apiWallet(vs []app.Coin) model.Wallet {
	ms := make(model.Wallet, len(vs))
	for i := range vs {
		ms[i] = string(vs[i])
	}
	return ms
}
