package main

import (
	"context"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/config"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/srv/openapi"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/powerman/check"
)

func TestMain(m *testing.M) {
	def.Init()
	initMetrics(reg, "test")
	app.InitMetrics(reg)
	openapi.InitMetrics(reg, "test")
	cfg = config.MustGetServeTest()
	app.GameLevelNormal.Seed = 3
	check.TestMain(m)
}

// Const shared by tests. Recommended naming scheme: <dataType><Variant>.
var (
	cfg *config.ServeConfig
	ctx = context.Background()
)
