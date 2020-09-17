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
	"github.com/go-openapi/swag"
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

	var (
		area = &model.Area{
			PosX:  swag.Int64(1247),
			PosY:  swag.Int64(1366),
			SizeX: 1,
			SizeY: 1,
		}
		treasure = model.Treasure(`{"X":1247,"Y":1366,"Depth":1}`)
	)

	{
		params := op.NewGetBalanceParams()
		res, err := openapiClient.Op.GetBalance(params)
		t.Nil(err)
		t.DeepEqual(res, &op.GetBalanceOK{Payload: &model.Balance{
			Balance: swag.Uint32(0),
			Wallet:  model.Wallet{},
		}})
	}
	{
		params := op.NewIssueLicenseParams().WithArgs(model.Wallet{})
		res, err := openapiClient.Op.IssueLicense(params)
		t.Nil(err)
		t.DeepEqual(res, &op.IssueLicenseOK{Payload: &model.License{
			ID:         swag.Int64(0),
			DigAllowed: 3,
			DigUsed:    0,
		}})
	}
	{
		params := op.NewExploreAreaParams().WithArgs(area)
		res, err := openapiClient.Op.ExploreArea(params)
		t.Nil(err)
		t.DeepEqual(res, &op.ExploreAreaOK{Payload: &model.Report{
			Area:           area,
			Amount:         1,
			AmountPerDepth: nil,
		}})
	}
	{
		params := op.NewDigParams().WithArgs(&model.Dig{
			LicenseID: swag.Int64(0),
			PosX:      area.PosX,
			PosY:      area.PosY,
			Depth:     swag.Int64(1),
		})
		res, err := openapiClient.Op.Dig(params)
		t.Nil(err)
		t.DeepEqual(res, &op.DigOK{Payload: model.TreasureList{treasure}})
	}
	{
		params := op.NewCashParams().WithArgs(treasure)
		res, err := openapiClient.Op.Cash(params)
		t.Nil(err)
		t.DeepEqual(res, &op.CashOK{Payload: model.Wallet{0}})
	}
	{
		params := op.NewListLicensesParams()
		res, err := openapiClient.Op.ListLicenses(params)
		t.Nil(err)
		t.DeepEqual(res, &op.ListLicensesOK{Payload: model.LicenseList{
			&model.License{
				ID:         swag.Int64(0),
				DigAllowed: 3,
				DigUsed:    1,
			},
		}})
	}
}
