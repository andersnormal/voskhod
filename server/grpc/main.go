package grpc

import (
	"context"

	pb "github.com/andersnormal/voskhod/proto"
)

func New() *grpcServer {
	return &grpcServer{}
}

func (g *grpcServer) RegisterAgent(ctx context.Context, in *pb.RegisterAgentRequest) (*pb.RegisterAgentResponse, error) {
	var err error

	return &pb.RegisterAgentResponse{
		Cluster: &pb.Cluster{},
	}, err
}
