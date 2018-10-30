// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package server

import (
	"context"
	"net"

	"github.com/katallaxie/voskhod/server/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/katallaxie/voskhod/proto"
	log "github.com/sirupsen/logrus"

	grpcServer "github.com/katallaxie/voskhod/server/grpc"
)

var _ Server = (*server)(nil)

// New is returning a new agent
func New(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
	}
}

// Start is starting the agent
func (s *server) Start(ctx context.Context) func() error {
	// set custom logger for the agent itself
	s.logger = log.WithFields(log.Fields{})

	return func() error {
		var err error

		// creates a new gGRPC server
		s.grpc = grpc.NewServer()
		pb.RegisterVoskhodServer(s.grpc, grpcServer.New())

		// creates a new listener
		lis, err := net.Listen("tcp", s.cfg.GrpcAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		// start couratin
		go s.Stop(ctx)()

		// Register reflection service on gRPC server.
		reflection.Register(s.grpc)
		if err := s.grpc.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		// noop
		return err
	}
}

// Stop is actually stopping the server and tearing down everything.
// Cleaning up the mess.
func (s *server) Stop(ctx context.Context) func() {
	return func() {
		defer s.grpc.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}
}

// func (s *server) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
// 	return &pb.CreateTaskResponse{Task: &pb.Task{Uuid: "test"}}, nil
// }
