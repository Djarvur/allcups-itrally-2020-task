package config

import (
	"os"
	"testing"
	"time"

	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/netx"
	"github.com/powerman/check"
)

func Test(t *testing.T) {
	want := &ServeConfig{
		APIKeyAdmin: "admin",
		Addr:        netx.NewAddr(def.Hostname, 8000),
		Duration:    10 * time.Minute,
		MetricsAddr: netx.NewAddr(def.Hostname, 9000),
		ResultDir:   "/data",
		WorkDir:     "/tmp",
	}

	t.Run("required", func(tt *testing.T) {
		t := check.T(tt)
		require(t, "APIKeyAdmin")
		os.Setenv("HLCUP2020_APIKEY_ADMIN", "admin")
	})
	t.Run("default", func(tt *testing.T) {
		t := check.T(tt)
		c, err := testGetServe()
		t.Nil(err)
		t.DeepEqual(c, want)
	})
	t.Run("constraint", func(tt *testing.T) {
		t := check.T(tt)
		constraint(t, "HLCUP2020_ADDR_PORT", "x", `^AddrPort .* invalid syntax`)
		constraint(t, "HLCUP2020_DURATION", "x", `^Duration .* invalid duration`)
		constraint(t, "HLCUP2020_METRICS_ADDR_PORT", "x", `^MetricsAddrPort .* invalid syntax`)
		constraint(t, "HLCUP2020_RESULT_DIR", "", `^ResultDir .* empty`)
		constraint(t, "HLCUP2020_WORK_DIR", "", `^WorkDir .* empty`)
	})
	t.Run("env", func(tt *testing.T) {
		t := check.T(tt)
		os.Setenv("HLCUP2020_APIKEY_ADMIN", "admin3")
		os.Setenv("HLCUP2020_ADDR_HOST", "localhost3")
		os.Setenv("HLCUP2020_ADDR_PORT", "8003")
		os.Setenv("HLCUP2020_DURATION", "3s")
		os.Setenv("HLCUP2020_METRICS_ADDR_PORT", "9003")
		os.Setenv("HLCUP2020_RESULT_DIR", "/data/3")
		os.Setenv("HLCUP2020_WORK_DIR", "/work/3")
		c, err := testGetServe()
		t.Nil(err)
		want.APIKeyAdmin = "admin3"
		want.Addr = netx.NewAddr("localhost3", 8003)
		want.Duration = 3 * time.Second
		want.MetricsAddr = netx.NewAddr("localhost3", 9003)
		want.ResultDir = "/data/3"
		want.WorkDir = "/work/3"
		t.DeepEqual(c, want)
	})
	t.Run("flag", func(tt *testing.T) {
		t := check.T(tt)
		c, err := testGetServe(
			"--host=localhost4",
			"--port=8004",
			"--duration=4ms",
			"--metrics.port=9004",
		)
		t.Nil(err)
		want.Addr = netx.NewAddr("localhost4", 8004)
		want.Duration = 4 * time.Millisecond
		want.MetricsAddr = netx.NewAddr("localhost4", 9004)
		t.DeepEqual(c, want)
	})
	t.Run("cleanup", func(tt *testing.T) {
		t := check.T(tt)
		t.Panic(func() { GetServe() })
	})
}
