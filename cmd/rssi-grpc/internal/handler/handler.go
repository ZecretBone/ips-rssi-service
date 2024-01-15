package handler

import (
	wiregrpc "github.com/RyuChk/wire-provider/grpc"
	v1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
)

type Handlers struct {
	Stat       v1.StatCollectionServiceServer
	Coordinate v1.CoordinateCollectionServiceServer
}

func RegisterService(server wiregrpc.Server, handler *Handlers) {
	v1.RegisterStatCollectionServiceServer(server, handler.Stat)
	v1.RegisterCoordinateCollectionServiceServer(server, handler.Coordinate)
}
