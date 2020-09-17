package openapi_test

import (
	"io"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client/op"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/srv/openapi"
	"github.com/go-openapi/swag"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

func TestGetBalance(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	cleanup, c, _, mockApp := testNewServer(t)
	defer cleanup()
	params := op.NewGetBalanceParams()

	mockApp.EXPECT().Balance(gomock.Any()).Return(0, nil, io.EOF)
	mockApp.EXPECT().Balance(gomock.Any()).Return(0, nil, nil)
	mockApp.EXPECT().Balance(gomock.Any()).Return(42, []int{1, 2}, nil)

	testCases := []struct {
		want    *model.Balance
		wantErr *model.Error
	}{
		{&model.Balance{}, apiError500},
		{&model.Balance{Balance: swag.Uint32(0), Wallet: model.Wallet{}}, nil},
		{&model.Balance{Balance: swag.Uint32(42), Wallet: model.Wallet{1, 2}}, nil},
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
