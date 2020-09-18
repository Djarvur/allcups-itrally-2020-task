// Package openapi implements OpenAPI server.
package openapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/restapi"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/restapi/op"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/netx"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/powerman/structlog"
	"github.com/sebest/xff"
)

type (
	// Ctx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
	// Config contains configuration for OpenAPI server.
	Config struct {
		Addr     netx.Addr
		BasePath string
	}
	server struct {
		app app.Appl
		cfg Config
	}
)

// NewServer returns OpenAPI server configured to listen on the TCP network
// address cfg.Host:cfg.Port and handle requests on incoming connections.
func NewServer(appl app.Appl, cfg Config) (*restapi.Server, error) {
	srv := &server{
		app: appl,
		cfg: cfg,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("load embedded swagger spec: %w", err)
	}
	if cfg.BasePath == "" {
		cfg.BasePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = cfg.BasePath

	api := op.NewHighLoadCup2020API(swaggerSpec)
	api.Logger = structlog.New(structlog.KeyUnit, "swagger").Printf

	api.GetBalanceHandler = op.GetBalanceHandlerFunc(srv.GetBalance)
	api.ListLicensesHandler = op.ListLicensesHandlerFunc(srv.ListLicenses)
	api.IssueLicenseHandler = op.IssueLicenseHandlerFunc(srv.IssueLicense)
	api.ExploreAreaHandler = op.ExploreAreaHandlerFunc(srv.ExploreArea)
	api.DigHandler = op.DigHandlerFunc(srv.Dig)
	api.CashHandler = op.CashHandlerFunc(srv.Cash)

	server := restapi.NewServer(api)
	server.Host = cfg.Addr.Host()
	server.Port = cfg.Addr.Port()

	// The middleware executes before anything.
	api.UseSwaggerUI()
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := makeLogger(cfg.BasePath)
		accesslog := makeAccessLog(cfg.BasePath)
		return noCache(xffmw.Handler(logger(recovery(accesslog(
			middleware.Spec(cfg.BasePath, restapi.FlatSwaggerJSON,
				cors(handler)))))))
	}
	// The middleware executes after serving /swagger.json and routing,
	// but before authentication, binding and validation.
	middlewares := func(handler http.Handler) http.Handler {
		appStart := makeAppStart(srv.app)
		return appStart(handler)
	}
	server.SetHandler(globalMiddlewares(api.Serve(middlewares)))

	log := structlog.New()
	log.Info("OpenAPI protocol", "version", swaggerSpec.Spec().Info.Version)
	return server, nil
}

func fromRequest(r *http.Request) (Ctx, Log) {
	ctx := r.Context()
	log := structlog.FromContext(ctx, nil)
	return ctx, log
}
