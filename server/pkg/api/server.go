package api

import (
	"net/http"

	pb "hexes/proto/go/hexes/v1"
	environmentAPI "hexes/server/pkg/api/environment-api"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func OriginAllowed(_ string) bool {
	return true
}

func NewServer() *http.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterEnvironmentAPIServer(grpcServer, environmentAPI.New())
	reflection.Register(grpcServer)

	httpServer := &http.Server{
		Handler: grpcweb.WrapServer(
			grpcServer,
			grpcweb.WithOriginFunc(OriginAllowed),
		),
	}

	return httpServer
}
