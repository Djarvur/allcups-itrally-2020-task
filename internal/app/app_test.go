package app_test

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/powerman/check"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
)

var (
	start = time.Now()
	dump  = nopCloser{bytes.NewReader([]byte("save"))}
)

func testPrepare(t *check.C) (func(), *app.MockRepo, *game.MockGame, *app.MockGameFactory, app.Config, func(a *app.App, err error)) {
	ctrl := gomock.NewController(t)
	mockRepo := app.NewMockRepo(ctrl)
	mockGame := game.NewMockGame(ctrl)
	mockGameFactory := app.NewMockGameFactory(ctrl)
	cfg := app.Config{
		Duration: time.Minute,
		Game:     app.Difficulty["test"],
	}
	wantErr := func(a *app.App, err error) {
		t.Helper()
		t.Err(err, io.EOF)
		t.Nil(a)
	}
	return ctrl.Finish, mockRepo, mockGame, mockGameFactory, cfg, wantErr
}

func TestNew(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	cleanup, mockRepo, mockGame, mockGameFactory, cfg, wantErr := testPrepare(t)
	defer cleanup()

	mockRepo.EXPECT().LoadStartTime().Return(nil, io.EOF)
	mockRepo.EXPECT().LoadStartTime().Return(&time.Time{}, nil).Times(6)
	wantErr(app.New(mockRepo, mockGameFactory, cfg))

	// Enforce random seed if difficulty is not "test" and seed is 0.
	cfgNormal := app.Difficulty["normal"]
	cfgNormal7 := cfgNormal
	cfgNormal7.Seed = 7
	cfgTest := app.Difficulty["test"]
	mockGameFactory.EXPECT().New(matchRandomSeed(cfgNormal)).Return(nil, io.EOF)
	mockGameFactory.EXPECT().New(cfgNormal7).Return(nil, io.EOF)
	mockGameFactory.EXPECT().New(cfgTest).Return(nil, io.EOF)
	mockGameFactory.EXPECT().New(cfgTest).Return(mockGame, nil).Times(3)
	cfg.Game = cfgNormal
	wantErr(app.New(mockRepo, mockGameFactory, cfg))
	cfg.Game = cfgNormal7
	wantErr(app.New(mockRepo, mockGameFactory, cfg))
	cfg.Game = cfgTest
	wantErr(app.New(mockRepo, mockGameFactory, cfg))

	mockRepo.EXPECT().SaveTreasureKey(gomock.Len(32)).Return(io.EOF)
	mockRepo.EXPECT().SaveTreasureKey(gomock.Len(32)).Return(nil).Times(2)
	wantErr(app.New(mockRepo, mockGameFactory, cfg))

	mockRepo.EXPECT().SaveGame(mockGame).Return(io.EOF)
	mockRepo.EXPECT().SaveGame(mockGame).Return(nil).MinTimes(1)
	wantErr(app.New(mockRepo, mockGameFactory, cfg))

	a, err := app.New(mockRepo, mockGameFactory, cfg)
	t.Nil(err)
	t.NotNil(a)
}

func TestContinue(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	cleanup, mockRepo, mockGame, mockGameFactory, cfg, wantErr := testPrepare(t)
	defer cleanup()

	mockRepo.EXPECT().LoadStartTime().Return(&start, nil).AnyTimes()

	mockRepo.EXPECT().LoadTreasureKey().Return(nil, io.EOF)
	mockRepo.EXPECT().LoadTreasureKey().Return(make([]byte, 33), nil)
	mockRepo.EXPECT().LoadTreasureKey().Return(make([]byte, 32), nil).Times(4)

	wantErr(app.New(mockRepo, mockGameFactory, cfg))

	a, err := app.New(mockRepo, mockGameFactory, cfg)
	t.Match(err, `bad PASETO key size`)
	t.Nil(a)

	mockRepo.EXPECT().LoadGame().Return(nil, io.EOF)
	mockRepo.EXPECT().LoadGame().Return(dump, nil).Times(3)
	wantErr(app.New(mockRepo, mockGameFactory, cfg))

	mockGameFactory.EXPECT().Continue(dump).Return(nil, io.EOF)
	mockGameFactory.EXPECT().Continue(dump).Return(mockGame, nil).Times(2)
	wantErr(app.New(mockRepo, mockGameFactory, cfg))

	mockRepo.EXPECT().SaveStartTime(start).Return(io.EOF)
	wantErr(app.New(mockRepo, mockGameFactory, cfg))

	mockRepo.EXPECT().SaveStartTime(start).Return(nil)
	a, err = app.New(mockRepo, mockGameFactory, cfg)
	t.Nil(err)
	t.NotNil(a)
}

func TestRestoreKey(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	cleanup, mockRepo, mockGame, mockGameFactory, cfg, _ := testPrepare(t)
	defer cleanup()
	var key []byte

	mockRepo.EXPECT().LoadStartTime().Return(&time.Time{}, nil)
	mockGameFactory.EXPECT().New(cfg.Game).Return(mockGame, nil)
	mockRepo.EXPECT().SaveTreasureKey(gomock.Len(32)).DoAndReturn(func(k []byte) error {
		key = k
		return nil
	})
	mockRepo.EXPECT().SaveGame(mockGame).Return(nil).MinTimes(1)
	a, err := app.New(mockRepo, mockGameFactory, cfg)
	t.Nil(err)
	t.NotNil(a)

	mockGame.EXPECT().Dig(1, game.Coord{X: 0, Y: 0, Depth: 1}).Return(true, nil)
	treasure, err := a.Dig(ctx, 1, game.Coord{X: 0, Y: 0, Depth: 1})

	mockRepo.EXPECT().LoadStartTime().Return(&start, nil)
	mockRepo.EXPECT().LoadTreasureKey().Return(key, nil)
	mockRepo.EXPECT().LoadGame().Return(dump, nil)
	mockGameFactory.EXPECT().Continue(dump).Return(mockGame, nil)
	mockRepo.EXPECT().SaveStartTime(start).Return(nil)
	a, err = app.New(mockRepo, mockGameFactory, cfg)
	t.Nil(err)
	t.NotNil(a)

	mockGame.EXPECT().Cash(game.Coord{X: 0, Y: 0, Depth: 1}).Return([]int{42}, nil)
	res, err := a.Cash(ctx, treasure)
	t.Nil(err)
	t.DeepEqual(res, []int{42})
}

type matchRandomSeed game.Config

func (m matchRandomSeed) String() string { return "has random Seed" }
func (m matchRandomSeed) Matches(x interface{}) bool {
	cfg, ok := x.(game.Config)
	if !ok {
		return false
	}
	if m.Seed == 0 && cfg.Seed > 0 {
		cfg.Seed = 0
		return game.Config(m) == cfg
	}
	return false
}

type nopCloser struct{ io.ReadSeeker }

func (nopCloser) Close() error { return nil }
