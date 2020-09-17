package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
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

func testNew(t *check.C) (func(), *app.App, *app.MockRepo, *game.MockGame) {
	t.Helper()
	ctrl := gomock.NewController(t)

	mockRepo := app.NewMockRepo(ctrl)
	mockRepo.EXPECT().LoadStartTime().Return(&time.Time{}, nil)
	mockGame := game.NewMockGame(ctrl)
	newGame := func(_ game.Config) (game.Game, error) { return mockGame, nil }

	a, err := app.New(mockRepo, newGame, app.Config{
		Duration: 60 * def.TestSecond,
		Game:     app.Difficulty["test"],
	})
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
