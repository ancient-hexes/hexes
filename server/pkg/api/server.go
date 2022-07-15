package api

import (
	"net/http"

	eventAPI "hexes/server/pkg/api/event-api"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "hexes/proto/go/hexes/v1"
	environmentAPI "hexes/server/pkg/api/environment-api"
)

func OriginAllowed(_ string) bool {
	return true
}

func WebSocketOriginAllowed(_ *http.Request) bool {
	return true
}

func NewServer() *http.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterEnvironmentAPIServer(grpcServer, environmentAPI.New())
	pb.RegisterEventAPIServer(grpcServer, eventAPI.New())
	reflection.Register(grpcServer)

	httpServer := &http.Server{
		Handler: grpcweb.WrapServer(
			grpcServer,
			grpcweb.WithWebsockets(true),
			grpcweb.WithOriginFunc(OriginAllowed),
			grpcweb.WithWebsocketOriginFunc(WebSocketOriginAllowed),
		),
	}

	return httpServer
}
