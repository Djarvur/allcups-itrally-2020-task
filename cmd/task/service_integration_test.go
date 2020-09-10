// +build integration

package main

import (
	"context"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client/op"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/netx"
	"github.com/powerman/check"
)

func TestSmoke(tt *testing.T) {
	t := check.T(tt)

	s := &service{cfg: cfg}

	ctxStartup, cancel := context.WithTimeout(ctx, def.TestTimeout)
	defer cancel()
	ctxShutdown, shutdown := context.WithCancel(ctx)
	errc := make(chan error)
	go func() { errc <- s.runServe(ctxStartup, ctxShutdown, shutdown) }()
	defer func() {
		shutdown()
		t.Nil(<-errc, "RunServe")
	}()
	t.Must(t.Nil(netx.WaitTCPPort(ctxStartup, cfg.Addr), "connect to service"))

	openapiClient := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Schemes:  []string{"http"},
		Host:     cfg.Addr.String(),
		BasePath: client.DefaultBasePath,
	})

	{
		params := op.NewGetBalanceParams()
		res, err := openapiClient.Op.GetBalance(params)
		t.TODO().Nil(err)
		t.TODO().DeepEqual(res, &op.GetBalanceOK{Payload: model.Wallet{}})
	}
}
