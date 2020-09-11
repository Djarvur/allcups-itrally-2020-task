// +build integration

package main

import (
	"context"
	"testing"
	"time"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client/op"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/netx"
	"github.com/powerman/check"
)

func TestTaskDuration(tt *testing.T) {
	t := check.T(tt)

	s := &service{cfg: cfg}
	s.cfg.Duration = def.TestSecond
	s.cfg.WorkDir = t.TempDir()
	s.cfg.ResultDir = t.TempDir()

	ctxStartup, cancel := context.WithTimeout(ctx, def.TestTimeout)
	defer cancel()
	ctxShutdown, shutdown := context.WithCancel(ctx)
	errc := make(chan error, 1)
	go func() { errc <- s.runServe(ctxStartup, ctxShutdown, shutdown) }()
	t.Must(t.Nil(netx.WaitTCPPort(ctxStartup, cfg.Addr), "connect to service"))

	openapiClient := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Schemes:  []string{"http"},
		Host:     cfg.Addr.String(),
		BasePath: client.DefaultBasePath,
	})
	openapiClient.Op.GetBalance(op.NewGetBalanceParams())

	start := time.Now()
	select {
	case err := <-errc:
		t.Nil(err)
		t.Between(time.Since(start), def.TestSecond/2, def.TestSecond*2)
	case <-time.After(def.TestTimeout):
		t.Fail()
	}
	shutdown()
}

func TestSmoke(tt *testing.T) {
	t := check.T(tt)

	s := &service{cfg: cfg}
	s.cfg.WorkDir = t.TempDir()
	s.cfg.ResultDir = t.TempDir()

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
