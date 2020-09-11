package app

import (
	"time"

	"github.com/powerman/structlog"
)

func (a *App) Wait(ctx Ctx) error {
	log := structlog.FromContext(ctx, nil)
	select {
	case <-ctx.Done():
	case t := <-a.started:
		dur := time.Until(t.Add(a.cfg.Duration))
		log.Info("task started", "dur", dur)
		select {
		case <-ctx.Done():
		case <-time.After(dur):
			log.Info("task finished")
		}
	}
	return nil
}

func (a *App) Start(t time.Time) (err error) {
	a.startOnce.Do(func() {
		a.started <- t
		err = a.repo.SaveStartTime(t)
	})
	return
}
