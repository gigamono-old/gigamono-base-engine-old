package server

import (
	"github.com/gigamono/gigamono-document-engine/pkg/engine"
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/gigamono/gigamono/pkg/inits"
)

// BaseEngineServer is a grpc server with an engine.
type BaseEngineServer struct {
	inits.App
	GinEngine  *gin.Engine
	BaseEngine engine.BaseEngine
}

// NewBaseEngineServer creates a new server instance.
func NewBaseEngineServer(app inits.App) (BaseEngineServer, error) {
	engine, err := engine.NewBaseEngine(&app)
	if err != nil {
		logs.FmtPrintln("initialising base engine server:", err)
		return BaseEngineServer{}, err
	}

	return BaseEngineServer{
		App:        app,
		GinEngine:  gin.Default(),
		BaseEngine: engine,
	}, nil
}

// Listen makes the server listen on specified port.
func (server *BaseEngineServer) Listen() error {
	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return server.grpcServe() })
	grp.Go(func() error { return server.httpServe() })
	return grp.Wait()
}
