// Package config provides configurations for subcommands.
//
// Default values can be obtained from various sources (constants,
// environment variables, etc.) and then overridden by flags.
//
// As configuration is global you can get it only once for safety:
// you can call only one of Get… functions and call it just once.
package config

import (
	"time"

	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/netx"
	"github.com/powerman/appcfg"
	"github.com/spf13/pflag"
)

// EnvPrefix defines common prefix for environment variables.
const envPrefix = "HLCUP2020_"

// All configurable values of the microservice.
//
// If microservice may runs in different ways (e.g. using CLI subcommands)
// then these subcommands may use subset of these values.
var all = &struct { //nolint:gochecknoglobals // Config is global anyway.
	APIKeyAdmin     appcfg.NotEmptyString `env:"APIKEY_ADMIN"`
	AddrHost        appcfg.NotEmptyString `env:"ADDR_HOST"`
	AddrPort        appcfg.Port           `env:"ADDR_PORT"`
	Duration        appcfg.Duration       `env:"DURATION"`
	MetricsAddrPort appcfg.Port           `env:"METRICS_ADDR_PORT"`
	ResultDir       appcfg.NotEmptyString `env:"RESULT_DIR"`
	WorkDir         appcfg.NotEmptyString `env:"WORK_DIR"`
}{ // Defaults, if any:
	AddrHost:        appcfg.MustNotEmptyString(def.Hostname),
	AddrPort:        appcfg.MustPort("8000"),
	Duration:        appcfg.MustDuration("10m"),
	MetricsAddrPort: appcfg.MustPort("9000"),
	ResultDir:       appcfg.MustNotEmptyString("var/data"),
	WorkDir:         appcfg.MustNotEmptyString("var"),
}

// FlagSets for all CLI subcommands which use flags to set config values.
type FlagSets struct {
	Serve *pflag.FlagSet
}

var fs FlagSets //nolint:gochecknoglobals // Flags are global anyway.

// Init updates config defaults (from env) and setup subcommands flags.
//
// Init must be called once before using this package.
func Init(flagsets FlagSets) error {
	fs = flagsets

	fromEnv := appcfg.NewFromEnv(envPrefix)
	err := appcfg.ProvideStruct(all, fromEnv)
	if err != nil {
		return err
	}

	appcfg.AddPFlag(fs.Serve, &all.AddrHost, "host", "host to serve OpenAPI")
	appcfg.AddPFlag(fs.Serve, &all.AddrPort, "port", "port to serve OpenAPI")
	appcfg.AddPFlag(fs.Serve, &all.Duration, "duration", "overall task duration")
	appcfg.AddPFlag(fs.Serve, &all.MetricsAddrPort, "metrics.port", "port to serve Prometheus metrics")

	return nil
}

// ServeConfig contains configuration for subcommand.
type ServeConfig struct {
	APIKeyAdmin string
	Addr        netx.Addr
	Duration    time.Duration
	MetricsAddr netx.Addr
	ResultDir   string
	WorkDir     string
}

// GetServe validates and returns configuration for subcommand.
func GetServe() (c *ServeConfig, err error) {
	defer cleanup()

	c = &ServeConfig{
		APIKeyAdmin: all.APIKeyAdmin.Value(&err),
		Addr:        netx.NewAddr(all.AddrHost.Value(&err), all.AddrPort.Value(&err)),
		Duration:    all.Duration.Value(&err),
		MetricsAddr: netx.NewAddr(all.AddrHost.Value(&err), all.MetricsAddrPort.Value(&err)),
		ResultDir:   all.ResultDir.Value(&err),
		WorkDir:     all.WorkDir.Value(&err),
	}
	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.Serve, all)
	}
	return c, nil
}

// Cleanup must be called by all Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	all = nil
}
