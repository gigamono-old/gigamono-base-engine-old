package engine

import (
	"github.com/gigamono/gigamono/pkg/inits"
)

// BaseEngine represents an engine instance.
type BaseEngine struct {
	*inits.App
}

// NewBaseEngine creates a new base engine.
func NewBaseEngine(app *inits.App) (BaseEngine, error) {
	return BaseEngine{
		App: app,
	}, nil
}
