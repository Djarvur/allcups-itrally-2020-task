package dal

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const fnStartTime = "start.time"

func (r *Repo) LoadStartTime() (*time.Time, error) {
	t := new(time.Time)
	path := filepath.Join(r.cfg.WorkDir, fnStartTime)
	buf, err := ioutil.ReadFile(path) //nolint:gosec // False positive.
	switch {
	case errors.Is(err, os.ErrNotExist):
		err = nil
	case err == nil:
		err = t.UnmarshalText(buf)
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *Repo) SaveStartTime(t time.Time) error {
	path := filepath.Join(r.cfg.WorkDir, fnStartTime)
	buf, err := t.MarshalText()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path+".tmp", buf, 0o600)
	if err == nil {
		err = os.Rename(path+".tmp", path)
	}
	return err
}
