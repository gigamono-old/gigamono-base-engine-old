package server

import (
	"fmt"
	"net"

	"github.com/gigamono/gigamono-document-engine/pkg/engine"
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
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
		logs.FmtPrintln("initialising Document Engine server:", err)
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
	// Listener on TCP port.
	listener, err := net.Listen("tcp", fmt.Sprint(":", server.Config.Services.Types.DocumentEngine.Port))
	if err != nil {
		return err
	}

	// Create multiplexer and delegate content-types.
	multiplexer := cmux.New(listener)
	grpcListener := multiplexer.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := multiplexer.Match(cmux.HTTP1Fast())

	// Run servers concurrently and sync errors.
	grp := new(errgroup.Group)
	grp.Go(func() error { return server.grpcServe(grpcListener) })
	grp.Go(func() error { return server.httpServe(httpListener) })
	grp.Go(func() error { return multiplexer.Serve() })
	return grp.Wait()
}
