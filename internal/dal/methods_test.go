package dal_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/dal"
	"github.com/powerman/check"
)

func TestStartTime(tt *testing.T) {
	t := check.T(tt)
	cfg := dal.Config{
		ResultDir: t.TempDir(),
		WorkDir:   t.TempDir(),
	}
	r, err := dal.New(ctx, cfg)
	t.Nil(err)

	start, err := r.LoadStartTime()
	t.Nil(err)
	t.DeepEqual(start, new(time.Time))

	prev := time.Date(2020, 1, 2, 3, 4, 5, 6, time.UTC)
	t.Nil(r.SaveStartTime(prev))

	start, err = r.LoadStartTime()
	t.Nil(err)
	t.DeepEqual(start, &prev)

	t.Nil(os.Chmod(filepath.Join(cfg.WorkDir, "start.time"), 0o000))
	start, err = r.LoadStartTime()
	t.Match(err, "permission denied")
	t.Nil(start)

	t.Nil(os.Chmod(cfg.WorkDir, 0o500))
	defer os.Chmod(cfg.WorkDir, 0o700)
	err = r.SaveStartTime(prev)
	t.Match(err, "permission denied")
}
