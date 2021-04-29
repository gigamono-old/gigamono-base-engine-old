package engine

import (
	"github.com/gigamono/gigamono/pkg/inits"
)

// DocumentEngine represents an engine instance.
type DocumentEngine struct {
	*inits.App
}

// NewDocumentEngine creates a new document engine.
func NewDocumentEngine(app *inits.App) (DocumentEngine, error) {
	return DocumentEngine{
		App: app,
	}, nil
}
