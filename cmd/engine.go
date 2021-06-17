package main

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/logs"

	"github.com/gigamono/gigamono-document-engine/pkg/server"
)

func main() {
	// Initialises app.
	app, err := inits.NewApp(inits.BaseEngineMainServer)
	if err != nil {
		logs.FmtPrintln("initialising base engine:", err)
		return
	}

	// Start an engine gRPC server.
	server, err := server.NewBaseEngineServer(app)
	if err != nil {
		logs.FmtPrintln("creating base engine:", err)
	}

	// Listen on port.
	if err := server.Listen(); err != nil {
		logs.FmtPrintln("trying to listen on port specified:", err)
	}
}
