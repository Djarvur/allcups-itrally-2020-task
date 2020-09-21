package app

import (
	"github.com/o1egl/paseto/v2"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
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

	return paseto.Encrypt(a.key, pos, "")
}

func (a *App) Cash(ctx Ctx, treasure string) (wallet []int, err error) {
	var pos game.Coord
	err = paseto.Decrypt(treasure, a.key, &pos, nil)
	if err != nil {
		return nil, err
	}

	return a.game.Cash(pos)
}
