package main

import (
	"context"
	"regexp"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/restapi"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/config"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/dal"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/srv/openapi"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/concurrent"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/serve"
	"github.com/powerman/structlog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

var reg = prometheus.NewPedanticRegistry() //nolint:gochecknoglobals // Metrics are global anyway.

type service struct {
	cfg  *config.ServeConfig
	repo *dal.Repo
	appl *app.App
	srv  *restapi.Server
}

func initService(_, serveCmd *cobra.Command) error {
	namespace := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(def.ProgName, "_")
	initMetrics(reg, namespace)
	app.InitMetrics(reg)
	openapi.InitMetrics(reg, namespace)

	return config.Init(config.FlagSets{
		Serve: serveCmd.Flags(),
	})
}

func (s *service) runServe(ctxStartup, ctxShutdown Ctx, shutdown func()) (err error) {
	log := structlog.FromContext(ctxShutdown, nil)
	if s.cfg == nil {
		s.cfg, err = config.GetServe()
	}
	if err != nil {
		return log.Err("failed to get config", "err", err)
	}

	err = concurrent.Setup(ctxStartup, map[interface{}]concurrent.SetupFunc{
		&s.repo: s.connectRepo,
	})
	if err != nil {
		return log.Err("failed to connect", "err", err)
	}

	if s.appl == nil {
		s.appl, err = app.New(s.repo, game.New, app.Config{
			Duration: s.cfg.Duration,
			Game:     s.cfg.Game,
		})
	}
	if err != nil {
		return log.Err("failed to app.New", "err", err)
	}
	s.srv, err = openapi.NewServer(s.appl, openapi.Config{
		Addr: s.cfg.Addr,
	})
	if err != nil {
		return log.Err("failed to openapi.NewServer", "err", err)
	}

	err = concurrent.Serve(ctxShutdown, shutdown,
		s.serveMetrics,
		s.serveOpenAPI,
		s.appl.Wait,
	)
	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	return nil
}

func (s *service) connectRepo(ctx Ctx) (interface{}, error) {
	return dal.New(ctx, dal.Config{
		ResultDir: s.cfg.ResultDir,
		WorkDir:   s.cfg.WorkDir,
	})
}

func (s *service) serveMetrics(ctx Ctx) error {
	return serve.Metrics(ctx, s.cfg.MetricsAddr, reg)
}

func (s *service) serveOpenAPI(ctx Ctx) error {
	return serve.OpenAPI(ctx, s.srv, "OpenAPI")
}
