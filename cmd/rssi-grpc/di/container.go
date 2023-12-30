package di

import wiregrpc "github.com/RyuChk/wire-provider/grpc"

type Container struct {
	server wiregrpc.Server
}
