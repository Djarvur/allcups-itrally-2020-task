package openapi_test

import (
	"io"
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/client/op"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
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

func TestListLicenses(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	cleanup, c, _, mockApp := testNewServer(t)
	defer cleanup()
	params := op.NewListLicensesParams()

	mockApp.EXPECT().Licenses(gomock.Any()).Return(nil, io.EOF)
	mockApp.EXPECT().Licenses(gomock.Any()).Return(nil, nil)
	mockApp.EXPECT().Licenses(gomock.Any()).Return([]game.License{
		{ID: 1, DigAllowed: 3, DigUsed: 0},
		{ID: 2, DigAllowed: 5, DigUsed: 1},
	}, nil)

	testCases := []struct {
		want    model.LicenseList
		wantErr *model.Error
	}{
		{model.LicenseList{}, apiError500},
		{model.LicenseList{}, nil},
		{model.LicenseList{
			&model.License{ID: swag.Int64(1), DigAllowed: 3, DigUsed: 0},
			&model.License{ID: swag.Int64(2), DigAllowed: 5, DigUsed: 1},
		}, nil},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			res, err := c.Op.ListLicenses(params)
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

func TestIssueLicense(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	cleanup, c, _, mockApp := testNewServer(t)
	defer cleanup()
	params := op.NewIssueLicenseParams()

	mockApp.EXPECT().IssueLicense(gomock.Any(), []int{}).Return(game.License{}, io.EOF)
	mockApp.EXPECT().IssueLicense(gomock.Any(), []int{0}).Return(game.License{ID: 1, DigAllowed: 3, DigUsed: 2}, nil)
	mockApp.EXPECT().IssueLicense(gomock.Any(), []int{0}).Return(game.License{}, game.ErrBogusCoin)
	mockApp.EXPECT().IssueLicense(gomock.Any(), []int{1, 2}).Return(game.License{}, game.ErrActiveLicenseLimit)

	testCases := []struct {
		wallet  model.Wallet
		want    *model.License
		wantErr *model.Error
	}{
		{model.Wallet{}, nil, apiError500},
		{model.Wallet{0}, &model.License{ID: swag.Int64(1), DigAllowed: 3, DigUsed: 2}, nil},
		{model.Wallet{0}, nil, apiError402},
		{model.Wallet{1, 2}, nil, apiError1002},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			params.Args = tc.wallet
			res, err := c.Op.IssueLicense(params)
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
