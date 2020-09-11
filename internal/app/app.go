//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo

// Package app provides business logic.
package app

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Errors.
var (
	ErrContactExists = errors.New("contact already exists")
)

// Appl provides application features (use cases) service.
type Appl interface {
	// Start must be called before any other method to ensure task
	// will be available for cfg.Duration since given time. Second and
	// following calls will have no effect, so it's safe to call Start
	// on every API call.
	// Errors: none.
	Start(time.Time) error
	// Balance returns current balance.
	// Errors: none.
	Balance(Ctx) ([]Coin, error)
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
	// Coin is an unique coin.
	Coin string
	// Contact describes record in address book.
	Contact struct {
		ID   int
		Name string
	}
)

type Config struct {
	Duration time.Duration
}

// App implements interface Appl.
type App struct {
	repo      Repo
	cfg       Config
	started   chan time.Time
	startOnce sync.Once
}

func New(repo Repo, cfg Config) (*App, error) {
	a := &App{
		repo:    repo,
		cfg:     cfg,
		started: make(chan time.Time, 1),
	}

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
