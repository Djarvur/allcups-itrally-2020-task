package game_test

import (
	"testing"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
	"github.com/powerman/check"
)

func TestSmoke(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	g, err := game.New(game.Config{
		Seed:              666, // {2 0 1}, {0 0 2}, {0 1 2} and 1 duplicate
		MaxActiveLicenses: 2,
		Density:           3,
		SizeX:             3,
		SizeY:             2,
		Depth:             2,
	})
	t.Nil(err)

	count, err := g.CountTreasures(game.Area{X: 0, Y: 0, SizeX: 3, SizeY: 2}, 1)
	t.Nil(err)
	t.Equal(count, 1)
	count, err = g.CountTreasures(game.Area{X: 0, Y: 0, SizeX: 3, SizeY: 2}, 2)
	t.Nil(err)
	t.Equal(count, 2)

	count, err = g.CountTreasures(game.Area{X: 2, Y: 0, SizeX: 1, SizeY: 1}, 1)
	t.Nil(err)
	t.Equal(count, 1)
	count, err = g.CountTreasures(game.Area{X: 0, Y: 0, SizeX: 1, SizeY: 2}, 2)
	t.Nil(err)
	t.Equal(count, 2)

	count, err = g.CountTreasures(game.Area{X: 0, Y: 0, SizeX: 2, SizeY: 2}, 1)
	t.Nil(err)
	t.Equal(count, 0)

	lic1, err := g.IssueLicense(2)
	t.Nil(err)
	t.DeepEqual(lic1, &game.License{ID: 0, DigAllowed: 2})
	lic2, err := g.IssueLicense(3)
	t.Nil(err)
	t.DeepEqual(lic2, &game.License{ID: 1, DigAllowed: 3})
	lic3, err := g.IssueLicense(1)
	t.Err(err, game.ErrActiveLicenseLimit)
	t.Nil(lic3)

	found, err := g.Dig(lic1.ID, game.Coord{X: 0, Y: 0, Depth: 1})
	t.Nil(err)
	t.False(found)
	found, err = g.Dig(lic1.ID, game.Coord{X: 0, Y: 0, Depth: 2})
	t.Nil(err)
	t.True(found)
	wallet, err := g.Cash(game.Coord{X: 0, Y: 0, Depth: 2})
	t.Nil(err)
	t.Len(wallet, 3)
	t.Len(g.Licenses(), 1)

	found, err = g.Dig(lic2.ID, game.Coord{X: 0, Y: 1, Depth: 2})
	t.Err(err, game.ErrWrongDepth)
	t.False(found)
	found, err = g.Dig(lic2.ID, game.Coord{X: 0, Y: 1, Depth: 1})
	t.Nil(err)
	t.False(found)
	found, err = g.Dig(lic2.ID, game.Coord{X: 0, Y: 1, Depth: 2})
	t.Nil(err)
	t.True(found)
	wallet, err = g.Cash(game.Coord{X: 0, Y: 1, Depth: 2})
	t.Nil(err)
	t.Len(wallet, 4)
	t.Len(g.Licenses(), 0)

	found, err = g.Dig(lic2.ID, game.Coord{X: 2, Y: 0, Depth: 1})
	t.Err(err, game.ErrNoSuchLicense)
	t.False(found)
	wallet, err = g.Cash(game.Coord{X: 2, Y: 0, Depth: 1})
	t.Err(err, game.ErrNotDigged)
	t.Nil(wallet)

	t.Nil(g.Spend([]int{0, 2}))
	balance, wallet := g.Balance()
	t.Equal(balance, 5)
	t.DeepEqual(wallet, []int{6, 5, 4, 3, 1})
}
