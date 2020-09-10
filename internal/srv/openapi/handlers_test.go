package openapi_test

import (
	"io"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client/op"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/srv/openapi"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

func TestGetBalance(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	c, _, mockApp := testNewServer(t)
	params := op.NewGetBalanceParams()

	mockApp.EXPECT().Balance(gomock.Any()).Return(nil, io.EOF)
	mockApp.EXPECT().Balance(gomock.Any()).Return(nil, nil)
	mockApp.EXPECT().Balance(gomock.Any()).Return([]app.Coin{"coin1", "coin2"}, nil)

	testCases := []struct {
		want    model.Wallet
		wantErr *model.Error
	}{
		{nil, apiError500},
		{model.Wallet{}, nil},
		{model.Wallet{"coin1", "coin2"}, nil},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			res, err := c.Op.GetBalance(params)
			if tc.wantErr == nil {
				t.Nil(openapi.ErrPayload(err))
				t.DeepEqual(res.Payload, tc.want)
			} else {
				t.DeepEqual(openapi.ErrPayload(err), tc.wantErr)
				t.Nil(res)
			}
		})
	}
}
