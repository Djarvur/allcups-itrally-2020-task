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

func TestNew(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := app.NewMockRepo(ctrl)
	cfg := app.Config{
		Duration: time.Minute,
		Game:     app.Difficulty["test"],
	}
	start := time.Now()

	mockRepo.EXPECT().LoadStartTime().Return(nil, io.EOF)
	_, err := app.New(mockRepo, game.Factory{}, cfg)
	t.Err(err, io.EOF)

	var dump *bytes.Buffer
	mockRepo.EXPECT().LoadStartTime().Return(&time.Time{}, nil)
	mockRepo.EXPECT().SaveTreasureKey(gomock.Any()).Return(nil)
	mockRepo.EXPECT().SaveGame(gomock.Any()).DoAndReturn(func(g io.WriterTo) error {
		if dump == nil {
			dump = new(bytes.Buffer)
			g.WriteTo(dump)
		}
		return nil
	}).AnyTimes()
	_, err = app.New(mockRepo, game.Factory{}, cfg)
	t.Nil(err)

	mockRepo.EXPECT().LoadStartTime().Return(&start, nil)
	mockRepo.EXPECT().LoadTreasureKey().Return(make([]byte, 32), nil)
	mockRepo.EXPECT().LoadGame().Return(nil, io.EOF)
	_, err = app.New(mockRepo, game.Factory{}, cfg)
	t.Err(err, io.EOF)

	dumpReader := nopCloser{bytes.NewReader(dump.Bytes())}
	mockRepo.EXPECT().LoadStartTime().Return(&start, nil)
	mockRepo.EXPECT().LoadTreasureKey().Return(make([]byte, 32), nil)
	mockRepo.EXPECT().LoadGame().Return(dumpReader, nil)
	mockRepo.EXPECT().SaveStartTime(start).Return(nil)
	_, err = app.New(mockRepo, game.Factory{}, cfg)
	t.Nil(err)
}

type nopCloser struct {
	io.ReadSeeker
}

func (nopCloser) Close() error { return nil }
