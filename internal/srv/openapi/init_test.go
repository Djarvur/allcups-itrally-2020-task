package openapi_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/srv/openapi"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/Djarvur/allcups-itrally-2020-task/pkg/netx"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	def.Init()
	reg := prometheus.NewPedanticRegistry()
	app.InitMetrics(reg)
	openapi.InitMetrics(reg, "test")
	check.TestMain(m)
}

// Const shared by tests. Recommended naming scheme: <dataType><Variant>.
var (
	apiError402  = openapi.APIError(402, "bogus coin")
	apiError403  = openapi.APIError(403, "no such license")
	apiError404  = openapi.APIError(404, "no treasure")
	apiError500  = openapi.APIError(500, "internal error")
	apiError1000 = openapi.APIError(1000, "wrong coordinates")
	apiError1001 = openapi.APIError(1001, "wrong depth")
	apiError1002 = openapi.APIError(1002, "no more active licenses allowed")
	apiError1003 = openapi.APIError(1003, "treasure is not digged")
)

func testNewServer(t *check.C) (cleanup func(), c *client.HighLoadCup2020, url string, mockAppl *app.MockAppl) {
	t.Helper()
	ctrl := gomock.NewController(t)

	mockAppl = app.NewMockAppl(ctrl)
	mockAppl.EXPECT().Start(gomock.Any()).Return(nil).AnyTimes()

	server, err := openapi.NewServer(mockAppl, openapi.Config{
		Addr: netx.NewAddr("localhost", 0),
	})
	t.Must(t.Nil(err, "NewServer"))
	t.Must(t.Nil(server.Listen(), "server.Listen"))
	errc := make(chan error, 1)
	go func() { errc <- server.Serve() }()

	cleanup = func() {
		t.Helper()
		t.Nil(server.Shutdown(), "server.Shutdown")
		t.Nil(<-errc, "server.Serve")
		ctrl.Finish()
	}

	ln, err := server.HTTPListener()
	t.Must(t.Nil(err, "server.HTTPListener"))
	c = client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Schemes:  []string{"http"},
		Host:     ln.Addr().String(),
		BasePath: client.DefaultBasePath,
	})
	url = fmt.Sprintf("http://%s", ln.Addr().String())

	// Avoid race between server.Serve() and server.Shutdown().
	ctx, cancel := context.WithTimeout(context.Background(), def.TestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	t.Must(t.Nil(err))
	_, err = (&http.Client{}).Do(req)
	t.Must(t.Nil(err, "connect to service"))

	return cleanup, c, url, mockAppl
}
