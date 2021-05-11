package server

import (
	"github.com/gigamono/gigamono-document-engine/pkg/engine"
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/gigamono/gigamono/pkg/inits"
)

// DocumentEngineServer is a grpc server with an engine.
type DocumentEngineServer struct {
	inits.App
	GinEngine      *gin.Engine
	DocumentEngine engine.DocumentEngine
}

// NewDocumentEngineServer creates a new server instance.
func NewDocumentEngineServer(app inits.App) (DocumentEngineServer, error) {
	engine, err := engine.NewDocumentEngine(&app)
	if err != nil {
		logs.FmtPrintln("initialising document engine server:", err)
		return DocumentEngineServer{}, err
	}
	return DocumentEngineServer{
		App:            app,
		GinEngine:      gin.Default(),
		DocumentEngine: engine,
	}, nil
}

// Listen makes the server listen on specified port.
func (server *DocumentEngineServer) Listen() error {
	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return server.grpcServe() })
	grp.Go(func() error { return server.httpServe() })
	return grp.Wait()
}
