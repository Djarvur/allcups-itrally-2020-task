package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/smartystreets/goconvey/convey"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
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
	cfg = app.Config{
		Duration:       60 * def.TestSecond,
		Game:           app.Difficulty["test"],
		AutosavePeriod: def.TestSecond,
	}
)

func testNew(t *check.C) (func(), *app.App, *app.MockRepo, *game.MockGame) {
	t.Helper()
	ctrl := gomock.NewController(t)

	mockRepo := app.NewMockRepo(ctrl)
	mockRepo.EXPECT().LoadStartTime().Return(&time.Time{}, nil)
	mockRepo.EXPECT().SaveTreasureKey(gomock.Any()).Return(nil)
	mockRepo.EXPECT().SaveGame(gomock.Any()).Return(nil).MinTimes(1)
	mockGame := game.NewMockGame(ctrl)
	mockGameFactory := app.NewMockGameFactory(ctrl)
	mockGameFactory.EXPECT().New(app.Difficulty["test"]).Return(mockGame, nil)

	a, err := app.New(mockRepo, mockGameFactory, cfg)
	t.Must(t.Nil(err))
	return ctrl.Finish, a, mockRepo, mockGame
}

func waitErr(t *check.C, errc <-chan error, wait time.Duration, wantErr error) {
	t.Helper()
	now := time.Now()
	select {
	case err := <-errc:
		t.Between(time.Since(now), wait-wait/4, wait+wait/4)
		t.Err(err, wantErr)
	case <-time.After(def.TestTimeout):
		t.FailNow()
	}
}
