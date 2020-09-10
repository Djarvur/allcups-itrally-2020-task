package app_test

import (
	"context"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	def.Init()
	reg := prometheus.NewPedanticRegistry()
	app.InitMetrics(reg)
	check.TestMain(m)
}

type Ctx = context.Context

// Const shared by tests. Recommended naming scheme: <dataType><Variant>.
var (
	ctx = context.Background()
)

func testNew(t *check.C) (*app.App, *app.MockRepo) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockRepo := app.NewMockRepo(ctrl)
	a := app.New(mockRepo)
	return a, mockRepo
}
