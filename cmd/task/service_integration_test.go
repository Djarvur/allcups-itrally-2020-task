// +build integration

package main

import (
	"context"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client/op"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/def"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/pkg/netx"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/srv/openapi"
	"github.com/powerman/check"
)

func TestSmoke(tt *testing.T) {
	t := check.T(tt)

	s := &service{cfg: cfg}

	ctxStartup, cancel := context.WithTimeout(ctx, 10*def.TestSecond)
	defer cancel()
	ctxShutdown, shutdown := context.WithCancel(ctx)
	errc := make(chan error)
	go func() { errc <- s.runServe(ctxStartup, ctxShutdown, shutdown) }()
	defer func() {
		shutdown()
		if err := <-errc; err != nil {
			t.Log("failed to Serve:", err)
		}
	}()
	t.Must(t.Nil(netx.WaitTCPPort(ctxStartup, cfg.Addr), "connect to service"))

	openapiClient := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Schemes:  []string{"http"},
		Host:     cfg.Addr.String(),
		BasePath: client.DefaultBasePath,
	})

	{
		params := op.NewAddContactParams()
		params.Contact = &model.Contact{Name: apiContact1.Name}
		res, err := openapiClient.Op.AddContact(params, apiKeyUser)
		t.DeepEqual(openapi.ErrPayload(err), apiError403)
		t.Nil(res)
	}
	{
		params := op.NewAddContactParams()
		params.Contact = &model.Contact{Name: apiContact1.Name}
		res, err := openapiClient.Op.AddContact(params, apiKeyAdmin)
		t.Nil(openapi.ErrPayload(err))
		t.DeepEqual(res, &op.AddContactCreated{Payload: apiContact1})
	}
	{
		params := op.NewListContactsParams()
		res, err := openapiClient.Op.ListContacts(params, apiKeyAdmin)
		t.Nil(openapi.ErrPayload(err))
		t.DeepEqual(res, &op.ListContactsOK{Payload: []*model.Contact{apiContact1}})
	}
}
