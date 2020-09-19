package openapi

import (
	"errors"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/restapi/op"
	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
)

func (srv *server) HealthCheck(params op.HealthCheckParams) op.HealthCheckResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	status, err := srv.app.HealthCheck(ctx)
	switch {
	default:
		return errHealthCheck(log, err, codeInternal)
	case err == nil:
		return op.NewHealthCheckOK().WithPayload(status)
	}
}

func (srv *server) GetBalance(params op.GetBalanceParams) op.GetBalanceResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	balance, wallet, err := srv.app.Balance(ctx)
	switch {
	default:
		return errGetBalance(log, err, codeInternal)
	case err == nil:
		return op.NewGetBalanceOK().WithPayload(apiBalance(balance, wallet))
	}
}

func (srv *server) ListLicenses(params op.ListLicensesParams) op.ListLicensesResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	licenses, err := srv.app.Licenses(ctx)
	switch {
	default:
		return errListLicenses(log, err, codeInternal)
	case err == nil:
		return op.NewListLicensesOK().WithPayload(apiLicenseList(licenses))
	}
}

func (srv *server) IssueLicense(params op.IssueLicenseParams) op.IssueLicenseResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	license, err := srv.app.IssueLicense(ctx, appWallet(params.Args))
	switch {
	default:
		return errIssueLicense(log, err, codeInternal)
	case errors.Is(err, game.ErrActiveLicenseLimit):
		return errIssueLicense(log, err, codeActiveLicenseLimit)
	case errors.Is(err, game.ErrBogusCoin):
		return errIssueLicense(log, err, codePaymentRequired)
	case err == nil:
		return op.NewIssueLicenseOK().WithPayload(apiLicense(license))
	}
}

func (srv *server) ExploreArea(params op.ExploreAreaParams) op.ExploreAreaResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	count, err := srv.app.ExploreArea(ctx, appArea(params.Args))
	switch {
	default:
		return errExploreArea(log, err, codeInternal)
	case errors.Is(err, game.ErrWrongCoord):
		return errExploreArea(log, err, codeWrongCoord)
	case err == nil:
		return op.NewExploreAreaOK().WithPayload(&model.Report{
			Area:   params.Args,
			Amount: model.Amount(count),
		})
	}
}

func (srv *server) Dig(params op.DigParams) op.DigResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	coord := appCoord(params.Args)
	treasure, err := srv.app.Dig(ctx, int(*params.Args.LicenseID), coord)
	switch {
	default:
		return errDig(log, err, codeInternal)
	case errors.Is(err, game.ErrNoSuchLicense):
		return errDig(log, err, codeForbidden)
	case errors.Is(err, game.ErrWrongCoord):
		return errDig(log, err, codeWrongCoord)
	case errors.Is(err, game.ErrWrongDepth):
		return errDig(log, err, codeWrongDepth)
	case err == nil && treasure == "":
		return errDig(log, game.ErrNoThreasure, codeNotFound)
	case err == nil:
		return op.NewDigOK().WithPayload(apiTreasureList(treasure))
	}
}

func (srv *server) Cash(params op.CashParams) op.CashResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	wallet, err := srv.app.Cash(ctx, string(params.Args))
	switch {
	default:
		return errCash(log, err, codeInternal)
	case errors.Is(err, game.ErrWrongCoord):
		return errCash(log, err, codeWrongCoord)
	case errors.Is(err, game.ErrNotDigged):
		return errCash(log, err, codeNotDigged)
	case errors.Is(err, game.ErrNoThreasure):
		return errCash(log, err, codeNotFound)
	case err == nil:
		return op.NewCashOK().WithPayload(apiWallet(wallet))
	}
}
