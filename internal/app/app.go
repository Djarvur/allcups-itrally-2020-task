//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo

// Package app provides business logic.
package app

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/powerman/must"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Errors.
var (
	ErrContactExists = errors.New("contact already exists")
)

// Appl provides application features (use cases) service.
type Appl interface {
	// HealthCheck returns error if service is unhealthy or current
	// status otherwise.
	// Errors: none.
	HealthCheck(Ctx) (interface{}, error)
	// Start must be called before any other method to ensure task
	// will be available for cfg.Duration since given time. Second and
	// following calls will have no effect, so it's safe to call Start
	// on every API call.
	// Errors: none.
	Start(time.Time) error
	// Balance returns current balance and up to 1000 issued coins.
	// Errors: none.
	Balance(Ctx) (balance int, wallet []int, err error)
	// Licenses returns all active licenses.
	// Errors: none.
	Licenses(Ctx) ([]game.License, error)
	// IssueLicense creates and returns a new license with given digAllowed.
	// Errors: game.ErrActiveLicenseLimit, game.ErrBogusCoin.
	IssueLicense(_ Ctx, wallet []int) (game.License, error)
	// ExploreArea returns amount of not-digged-yet treasures in the
	// area at depth.
	// Errors: game.ErrWrongCoord.
	ExploreArea(_ Ctx, area game.Area) (int, error)
	// Dig tries to dig at pos and returns if any treasure was found.
	// The pos depth must be next to current (already digged) one.
	// Also it increment amount of used dig calls in given active license.
	// If amount of used dig calls became equal to amount of allowed dig calls
	// then license will became inactive after the call.
	// Errors: game.ErrNoSuchLicense, game.ErrWrongCoord, game.ErrWrongDepth.
	Dig(_ Ctx, licenseID int, pos game.Coord) (treasure string, _ error)
	// Cash returns coins earned for treasure as given pos.
	// Errors: game.ErrWrongCoord, game.ErrNotDigged, game.ErrAlreadyCached.
	Cash(_ Ctx, treasure string) (wallet []int, err error)
}

// Repo provides data storage.
type Repo interface {
	// LoadStartTime returns start time or zero time if not started.
	// Errors: none.
	LoadStartTime() (*time.Time, error)
	// SaveStartTime stores start time.
	// Errors: none.
	SaveStartTime(t time.Time) error
}

type (
	// Contact describes record in address book.
	Contact struct {
		ID   int
		Name string
	}
)

// Difficulty contains predefined game difficulty levels.
//nolint:gochecknoglobals,gomnd // Const.
var Difficulty = map[string]game.Config{
	"test": {
		MaxActiveLicenses: 3,
		Density:           4,
		SizeX:             5,
		SizeY:             5,
		Depth:             10,
	},
	"normal": {
		MaxActiveLicenses: 3,
		Density:           250,
		SizeX:             3500,
		SizeY:             3500,
		Depth:             10,
	},
}

type Config struct {
	Duration time.Duration
	Game     game.Config
}

// App implements interface Appl.
type App struct {
	repo      Repo
	cfg       Config
	game      game.Game
	started   chan time.Time
	startOnce sync.Once
	key       jwk.SymmetricKey
}

// GameFactory creates and returns new game.
type GameFactory func(game.Config) (game.Game, error)

func New(repo Repo, newGame GameFactory, cfg Config) (*App, error) {
	if cfg.Game != Difficulty["test"] && cfg.Game.Seed == 0 {
		cfg.Game.Seed = time.Now().UnixNano() // TODO Restore after crash?
	}
	g, err := newGame(cfg.Game)
	if err != nil {
		return nil, fmt.Errorf("newGame: %w", err)
	}

	a := &App{
		repo:    repo,
		cfg:     cfg,
		game:    g,
		started: make(chan time.Time, 1),
		key:     jwk.NewSymmetricKey(),
	}

	buf := make([]byte, 16)
	_, err = io.ReadFull(rand.Reader, buf)
	must.NoErr(err)
	must.NoErr(a.key.FromRaw(buf))

	t, err := a.repo.LoadStartTime()
	if err != nil {
		return nil, fmt.Errorf("LoadStartTime: %w", err)
	}
	if !t.IsZero() {
		err = a.Start(*t)
	}
	if err != nil {
		return nil, fmt.Errorf("SaveStartTime: %w", err)
	}
	return a, nil
}

func (a *App) HealthCheck(_ Ctx) (interface{}, error) {
	return "OK", nil
}
