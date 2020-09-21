package app_test

import (
	"io"
	"testing"

	"github.com/powerman/check"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
)

func TestBalance(tt *testing.T) {
	t := check.T(tt)
	cleanup, a, _, mockGame := testNew(t)
	defer cleanup()

	mockGame.EXPECT().Balance().Return(0, nil)
	balance, wallet, err := a.Balance(ctx)
	t.Nil(err)
	t.Equal(balance, 0)
	t.Len(wallet, 0)
}

func TestExploreArea(tt *testing.T) {
	t := check.T(tt)
	cleanup, a, _, mockGame := testNew(t)
	defer cleanup()

	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(1)).Return(5, nil)
	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(2)).Return(4, nil)
	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(3)).Return(3, nil)
	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(4)).Return(2, nil)
	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(5)).Return(1, nil)
	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(6)).Return(0, nil)
	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(7)).Return(0, nil)
	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(8)).Return(1, nil)
	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(9)).Return(2, nil)
	mockGame.EXPECT().CountTreasures(game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5}, uint8(10)).Return(3, nil)
	count, err := a.ExploreArea(ctx, game.Area{X: 0, Y: 0, SizeX: 5, SizeY: 5})
	t.Nil(err)
	t.Equal(count, 21)

	mockGame.EXPECT().CountTreasures(game.Area{X: 5, Y: 0, SizeX: 1, SizeY: 1}, uint8(1)).Return(0, io.EOF)
	count, err = a.ExploreArea(ctx, game.Area{X: 5, Y: 0, SizeX: 1, SizeY: 1})
	t.Err(err, io.EOF)
	t.Equal(count, 0)
}
