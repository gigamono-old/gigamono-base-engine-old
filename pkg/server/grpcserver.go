package server

import (
	"context"
	"fmt"
	"net"

	"github.com/gigamono/gigamono/pkg/services/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (server *BaseEngineServer) grpcServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(":", server.Config.Services.BaseEngine.Ports.Private),
	)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	generated.RegisterBaseEngineMainServerServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello says Hello
func (server *BaseEngineServer) SayHello(ctx context.Context, msg *generated.Message) (*generated.Message, error) {
	engineMsg := "Document Engine Main Server replies: " + msg.Content
	fmt.Println(engineMsg)
	response := generated.Message{
		Content: engineMsg,
	}
	return &response, nil
}
