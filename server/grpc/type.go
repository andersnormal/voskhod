package grpc

import (
	pb "github.com/katallaxie/voskhod/proto"
)

var _ pb.VoskhodServer = (*grpcServer)(nil)

type grpcServer struct{}
