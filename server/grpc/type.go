package grpc

import (
	pb "github.com/andersnormal/voskhod/proto"
)

var _ pb.VoskhodServer = (*grpcServer)(nil)

type grpcServer struct{}
