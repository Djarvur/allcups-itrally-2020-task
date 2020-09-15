// Package game implements treasure hunting game.
package game

import (
	"errors"
	"fmt"
	prng "math/rand"
	"sync"

	"github.com/powerman/structlog"
)

const (
	maxSizeX, maxSizeY, maxDepth = 6000, 6000, 10 // About 1GB RAM without bit-optimization.
)

// Errors.
var (
	ErrOutOfBounds        = errors.New("out of bounds")
	ErrActiveLicenseLimit = errors.New("no more active licenses allowed")
	ErrNoSuchLicense      = errors.New("no such license")
	ErrNotEnoughMoney     = errors.New("not enough money")
	ErrCoinNotIssued      = errors.New("coin not issued")
	ErrCoinNotExists      = errors.New("coin does not exist")
	ErrWrongAmount        = errors.New("wrong amount of coins")
	ErrAreaCoord          = errors.New("wrong area coordinates")
	ErrWrongCoord         = errors.New("wrong coordinates")
	ErrWrongDepth         = errors.New("wrong depth")
	ErrAlreadyCached      = errors.New("already cashed")
	ErrNotDigged          = errors.New("treasure is not digged")
)

// Game implements treasure hunting game.
type Game struct {
	cfg      Config
	log      *structlog.Logger
	licenses *licenses
	bank     *bank
	field    *field
	muPRNG   sync.Mutex
	prng     *prng.Rand
}

// New creates and returns new game.
func New(cfg Config) (*Game, error) {
	switch {
	case cfg.Density <= 0, cfg.Density > cfg.volume(): // Min 1 treasure.
		return nil, fmt.Errorf("%w: Density", ErrOutOfBounds)
	case cfg.SizeX <= 0, cfg.SizeX > maxSizeX:
		return nil, fmt.Errorf("%w: SizeX", ErrOutOfBounds)
	case cfg.SizeY <= 0, cfg.SizeY > maxSizeY:
		return nil, fmt.Errorf("%w: SizeY", ErrOutOfBounds)
	case cfg.Depth <= 0, cfg.Depth > maxDepth:
		return nil, fmt.Errorf("%w: Depth", ErrOutOfBounds)
	}

	g := &Game{
		cfg:      cfg,
		log:      structlog.New(),
		licenses: newLicenses(cfg.MaxActiveLicenses),
		bank:     newBank(cfg.totalCash()),
		field:    newField(cfg),
		prng:     prng.New(prng.NewSource(cfg.Seed)), //nolint:gosec // We need repeatable game results.
	}

	for i := 0; i < cfg.treasures(); i++ {
		pos := Coord{
			X:     g.prng.Intn(cfg.SizeX),
			Y:     g.prng.Intn(cfg.SizeY),
			Depth: uint8(g.prng.Intn(int(cfg.Depth)) + 1),
		}
		if !g.field.addTreasure(pos) {
			g.log.Warn("skip adding duplicate treasure")
		} else {
			g.log.Info("added treasure", "pos", pos)
		}
	}
	return g, nil
}

// Balance returns current balance and up to 1000 issued coins.
func (g *Game) Balance() (balance int, wallet []int) {
	return g.bank.getBalance()
}

// Spend mark given coins as not issued (returns them into the bank).
func (g *Game) Spend(wallet []int) error {
	return g.bank.spend(wallet)
}

// Licenses returns all active licenses.
func (g *Game) Licenses() []License {
	return g.licenses.active()
}

// IssueLicense creates and returns a new license with given digAllowed.
func (g *Game) IssueLicense(digAllowed int) (*License, error) {
	return g.licenses.issue(digAllowed)
}

// CountTreasures returns amount of not-digged-yet treasures in the area
// at depth.
func (g *Game) CountTreasures(area Area, depth uint8) (int, error) {
	return g.field.countTreasures(area, depth)
}

// Dig tries to dig at pos and returns if any treasure was found.
// The pos depth must be next to current (already digged) one.
// Also it increment amount of used dig calls in given active license.
// If amount of used dig calls became equal to amount of allowed dig calls
// then license will became inactive after the call.
func (g *Game) Dig(licenseID int, pos Coord) (found bool, _ error) {
	err := g.licenses.use(licenseID)
	if err != nil {
		return false, err
	}
	return g.field.dig(pos)
}

// Cash returns coins earned for treasure as given pos.
func (g *Game) Cash(pos Coord) (wallet []int, err error) {
	err = g.field.cash(pos)
	if err != nil {
		return nil, err
	}

	g.muPRNG.Lock()
	defer g.muPRNG.Unlock()
	min, max := treasureCostAt(pos.Depth)
	amount := min + g.prng.Intn(max-min+1)

	return g.bank.earn(amount)
}

func treasureCostAt(depth uint8) (min, max int) {
	return int(depth), int(depth * 2) //nolint:gomnd // TODO Balance?
}
