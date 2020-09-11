// Package dal implements Data Access Layer using in-memory DB.
package dal

import (
	"context"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

type Config struct {
	WorkDir   string
	ResultDir string
}

// Repo provides access to storage.
type Repo struct {
	cfg Config
}

// New creates and returns new Repo.
func New(_ Ctx, cfg Config) (*Repo, error) {
	return &Repo{cfg: cfg}, nil
}
