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
)

var _ Server = (*server)(nil)

// New is returning a new agent
func New(ctx context.Context, cfg *config.Config) Server {
	return &server{
		cfg: cfg,
		ctx: ctx,
	}
}

// Start is starting the agent
func (s *server) Start() func() error {
	// set custom logger for the agent itself
	s.logger = log.WithFields(log.Fields{})

	return func() error {
		var err error

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterVoskhodServer(s, &server{})

		// Register reflection service on gRPC server.
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		// a.dc = dc // assign the client to the agent
		// a.events = e.New(dc.ContainerEvents(a.ctx))

		// // generating a event listener channel
		// l := make(chan events.Message)
		// a.events.AddEventListener(l)

		// // go to func to golem
		// go func() {
		// 	for {
		// 		msg := <-l
		// 		a.handleMessage(msg) // handle the event message, and translate to our events
		// 	}
		// }()

		// init hearbeat for docker client,
		// because we do not know if the client still live

		// a.logger.Infof(fmt.Sprintf("Agent succesfully started ..."))

		// // just have one channel to end it,\
		// // so not using something differne
		// <-a.ctx.Done()

		// cleanup
		// err = a.Stop()

		// noop
		return err
	}
}

// Stop is actually stopping the agent and tearing down everything.
// Cleaning up the mess.
func (s *server) Stop() error {
	var err error
	return err
}

func (s *server) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	return &pb.CreateTaskResponse{Task: &pb.Task{Uuid: "test"}}, nil
}
