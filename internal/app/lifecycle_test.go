package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/powerman/check"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
)

func TestWait(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	cleanup, a, mockRepo, _ := testNew(t)
	defer cleanup()
	mockRepo.EXPECT().SaveResult(0).Return(nil)

	{ // ctxShutdown before a.Start().
		ctx, shutdown := context.WithCancel(ctx)
		errc := make(chan error, 1)
		go func() { errc <- a.Wait(ctx) }()
		go func() { time.Sleep(def.TestSecond / 10); shutdown() }()
		waitErr(t, errc, def.TestSecond/10, nil)
	}
	{ // ctxShutdown after a.Start().
		ctx, shutdown := context.WithCancel(ctx)
		errc := make(chan error, 1)
		go func() { errc <- a.Wait(ctx) }()
		mockRepo.EXPECT().SaveStartTime(gomock.Any()).Return(nil)
		a.Start(time.Now())
		go func() { time.Sleep(def.TestSecond / 10); shutdown() }()
		waitErr(t, errc, def.TestSecond/10, nil)
	}
	{ // No ctxShutdown.
		mockRepo.EXPECT().LoadStartTime().Return(&time.Time{}, nil)
		mockRepo.EXPECT().SaveTreasureKey(gomock.Any()).Return(nil)
		mockRepo.EXPECT().SaveGame(gomock.Any()).Return(nil).AnyTimes()
		a, err := app.New(mockRepo, game.Factory{}, app.Config{
			Duration: def.TestSecond,
			Game:     app.Difficulty["test"],
		})
		t.Must(t.Nil(err))
		ctx, shutdown := context.WithCancel(ctx)
		defer shutdown()
		errc := make(chan error, 1)
		go func() { errc <- a.Wait(ctx) }()
		// Waiting for a.Start().
		select {
		case <-errc:
			t.FailNow()
		case <-time.After(def.TestSecond + def.TestSecond/4):
		}
		// Finish in cfg.Duration after first a.Start().
		// Second Start() will be ignored.
		now := time.Now()
		mockRepo.EXPECT().SaveStartTime(now).Return(nil)
		a.Start(now)
		time.Sleep(def.TestSecond / 2)
		a.Start(now.Add(def.TestSecond / 2))
		waitErr(t, errc, def.TestSecond/2, nil)
	}
}
