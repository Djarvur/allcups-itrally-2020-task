package app

import (
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/powerman/must"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
)

const (
	claimPosX     = "pos.X"
	claimPosY     = "pos.Y"
	claimPosDepth = "pos.Depth"
)

func (a *App) Balance(ctx Ctx) (balance int, wallet []int, err error) {
	balance, wallet = a.game.Balance()
	return balance, wallet, nil
}

func (a *App) Licenses(ctx Ctx) ([]game.License, error) {
	return a.game.Licenses(), nil
}

func (a *App) IssueLicense(ctx Ctx, wallet []int) (game.License, error) {
	return a.game.IssueLicense(wallet)
}

func (a *App) ExploreArea(ctx Ctx, area game.Area) (int, error) {
	sum := 0
	for depth := uint8(1); depth <= a.cfg.Game.Depth; depth++ {
		count, err := a.game.CountTreasures(area, depth)
		if err != nil {
			return 0, err
		}
		sum += count
	}
	return sum, nil
}

func (a *App) Dig(ctx Ctx, licenseID int, pos game.Coord) (treasure string, _ error) {
	found, err := a.game.Dig(licenseID, pos)
	if err != nil {
		return "", err
	}
	if !found {
		return "", nil
	}

	t := jwt.New()
	must.NoErr(t.Set(claimPosX, pos.X))
	must.NoErr(t.Set(claimPosY, pos.Y))
	must.NoErr(t.Set(claimPosDepth, pos.Depth))
	token, err := jwt.Sign(t, jwa.HS256, a.key)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (a *App) Cash(ctx Ctx, treasure string) (wallet []int, err error) {
	t, err := jwt.ParseBytes([]byte(treasure), jwt.WithVerify(jwa.HS256, a.key.Octets()))
	if err != nil {
		return nil, err
	}
	claims := t.PrivateClaims()
	pos := game.Coord{
		X:     int(claims[claimPosX].(float64)),
		Y:     int(claims[claimPosY].(float64)),
		Depth: uint8(claims[claimPosDepth].(float64)),
	}
	return a.game.Cash(pos)
}
