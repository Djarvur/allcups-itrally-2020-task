package openapi_test

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/restapi"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/def"
	"github.com/go-openapi/loads"
	"github.com/powerman/check"
)

func TestServeSwagger(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	_, tsURL, _ := testNewServer(t)
	c := &http.Client{}
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	t.Nil(err)
	basePath := swaggerSpec.BasePath()

	testCases := []struct {
		path string
		want int
	}{
		{"/", 404},
		{"/swagger.yml", 404},
		{"/swagger.yaml", 404},
		{"/swagger.json", 200},
		{basePath, 404},
		{path.Join(basePath, "docs"), 200},
		{path.Join(basePath, "swagger.json"), 200},
	}
	for _, tc := range testCases {
		ctx, cancel := context.WithTimeout(context.Background(), 10*def.TestSecond)
		req, err := http.NewRequestWithContext(ctx, "GET", tsURL+tc.path, nil)
		t.Nil(err)
		resp, err := c.Do(req)
		t.Nil(err, tc.path)
		t.Equal(resp.StatusCode, tc.want, tc.path)
		cancel()
	}
}
