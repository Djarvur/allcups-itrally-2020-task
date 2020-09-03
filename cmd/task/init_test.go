package main

import (
	"context"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/config"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/def"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/srv/openapi"
	oapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"
	"github.com/powerman/check"
)

func TestMain(m *testing.M) {
	def.Init()
	initMetrics(reg, "test")
	app.InitMetrics(reg)
	openapi.InitMetrics(reg, "test")
	cfg = config.MustGetServeTest()
	check.TestMain(m)
}

// Const shared by tests. Recommended naming scheme: <dataType><Variant>.
var (
	cfg         *config.ServeConfig
	ctx         = context.Background()
	apiError403 = openapi.APIError(403, "only admin can make changes")
	apiKeyAdmin = oapiclient.APIKeyAuth("API-Key", "header", "admin")
	apiKeyUser  = oapiclient.APIKeyAuth("API-Key", "header", "user")
	apiContact1 = &model.Contact{ID: 1, Name: swag.String("A")}
)
