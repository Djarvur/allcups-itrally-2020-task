package app_test

import (
	"io"
	"testing"
	"time"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

func TestNew(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockRepo := app.NewMockRepo(ctrl)
	cfg := app.Config{
		Duration: time.Minute,
	}
	start := time.Now()

	mockRepo.EXPECT().LoadStartTime().Return(nil, io.EOF)
	_, err := app.New(mockRepo, cfg)
	t.Err(err, io.EOF)

	mockRepo.EXPECT().LoadStartTime().Return(&time.Time{}, nil)
	_, err = app.New(mockRepo, cfg)
	t.Nil(err)

	mockRepo.EXPECT().LoadStartTime().Return(&start, nil)
	mockRepo.EXPECT().SaveStartTime(start).Return(io.EOF)
	_, err = app.New(mockRepo, cfg)
	t.Err(err, io.EOF)

	mockRepo.EXPECT().LoadStartTime().Return(&start, nil)
	mockRepo.EXPECT().SaveStartTime(start).Return(nil)
	_, err = app.New(mockRepo, cfg)
	t.Nil(err)
}