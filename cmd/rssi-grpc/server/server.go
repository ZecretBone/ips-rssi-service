package server

import (
	wiregrpc "github.com/RyuChk/wire-provider/grpc"
	"github.com/RyuChk/wire-provider/grpc/provider"
	"github.com/ZecretBone/ips-rssi-service/cmd/rssi-grpc/internal/handler"
)

type grpcServerCustomizer struct {
	handlers *handler.Handlers
}

func (gc *grpcServerCustomizer) Register(server wiregrpc.Server) error {
	handler.RegisterService(server, gc.handlers)
	return nil
}

func (gc *grpcServerCustomizer) Configrue(builder wiregrpc.Builder) error {
	builder.WithServiceName("bff-service").WithListenAddr(":6000")
	return nil
}

func ProvideGRPCServerCustomizer(handlers *handler.Handlers) provider.GRPCServerCustomizer {
	return &grpcServerCustomizer{
		handlers: handlers,
	}
}
