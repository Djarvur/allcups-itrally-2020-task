//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo,GameFactory

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

	"github.com/powerman/must"
	"github.com/powerman/structlog"

	"github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Errors.
var (
	ErrContactExists    = errors.New("contact already exists")
	errBadPASETOKeySize = errors.New("bad PASETO key size")
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
	// LoadTreasureKey returns treasure key.
	// Errors: none.
	LoadTreasureKey() ([]byte, error)
	// SaveTreasureKey stores treasure key.
	// Errors: none.
	SaveTreasureKey([]byte) error
	// LoadGame returns game state.
	// Errors: none.
	LoadGame() (ReadSeekCloser, error)
	// SaveGame stores game state.
	// Errors: none.
	SaveGame(io.WriterTo) error
	// SaveResult stores final game result.
	// Errors: none.
	SaveResult(int) error
}

type (
	// Contact describes record in address book.
	Contact struct {
		ID   int
		Name string
	}
	// ReadSeekCloser is the interface that groups the basic Read,
	// Seek and Close methods.
	ReadSeekCloser interface {
		io.ReadSeeker
		io.Closer
	}
)

const pasetoKeySize = 32

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
	key       []byte
}

// GameFactory creates and returns new game.
type GameFactory interface {
	New(cfg game.Config) (game.Game, error)
	Continue(r io.ReadSeeker) (game.Game, error)
}

func New(repo Repo, factory GameFactory, cfg Config) (*App, error) {
	a := &App{
		repo:    repo,
		cfg:     cfg,
		started: make(chan time.Time, 1),
		key:     make([]byte, pasetoKeySize),
	}
	t, err := a.repo.LoadStartTime()
	if err != nil {
		return nil, fmt.Errorf("LoadStartTime: %w", err)
	}
	if t.IsZero() {
		err = a.newGame(factory)
	} else {
		err = a.continueGame(factory, *t)
	}
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) newGame(factory GameFactory) (err error) {
	if a.cfg.Game != Difficulty["test"] && a.cfg.Game.Seed == 0 {
		a.cfg.Game.Seed = time.Now().UnixNano()
	}

	_, err = io.ReadFull(rand.Reader, a.key)
	must.NoErr(err)

	a.game, err = factory.New(a.cfg.Game)
	if err != nil {
		return fmt.Errorf("newGame: %w", err)
	}

	err = a.repo.SaveTreasureKey(a.key)
	if err != nil {
		return fmt.Errorf("SaveTreasureKey: %w", err)
	}
	err = a.repo.SaveGame(a.game)
	if err != nil {
		return fmt.Errorf("SaveGame: %w", err)
	}

	structlog.New().Info("new game")
	return nil
}

func (a *App) continueGame(factory GameFactory, t time.Time) (err error) {
	a.key, err = a.repo.LoadTreasureKey()
	if err != nil {
		return fmt.Errorf("LoadTreasureKey: %w", err)
	}
	if len(a.key) != pasetoKeySize {
		return fmt.Errorf("%w: %d", errBadPASETOKeySize, len(a.key))
	}

	f, err := a.repo.LoadGame()
	if err != nil {
		return fmt.Errorf("LoadGame: %w", err)
	}
	a.game, err = factory.Continue(f)
	if err != nil {
		return fmt.Errorf("factory.Continue: %w", err)
	}
	err = f.Close()
	if err != nil {
		return fmt.Errorf("LoadGame.Close: %w", err)
	}

	structlog.New().Info("continue game")
	err = a.Start(t)
	if err != nil {
		return fmt.Errorf("SaveStartTime: %w", err)
	}
	return nil
}

func (a *App) HealthCheck(_ Ctx) (interface{}, error) {
	return "OK", nil
}
