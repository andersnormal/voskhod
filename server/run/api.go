package server

import (
	"context"
	"net"

	pb "github.com/andersnormal/voskhod/proto"

	"google.golang.org/grpc"
)

func (s *server) ServeAPI() {
	g := s.errG

	s.api = grpc.NewServer()
	pb.RegisterVoskhodServer(s.api, &API{s.config(), s.log()})

	g.Go(s.serveAPI())
}

func (s *server) serveAPI() func() error {
	return func() error {
		var err error

		lis, err := net.Listen("tcp", s.cfg.APIListener())
		if err != nil {
			s.log().Error(err)
			return err
		}

		s.log().Infof("Listening on %s", lis.Addr())

		if err = s.api.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) shutdownAPI() func() error {
	return func() error {
		s.api.GracefulStop()

		return nil
	}
}

// API

func (a *API) RegisterAgent(ctx context.Context, req *pb.RegisterAgentRequest) (*pb.RegisterAgentResponse, error) {
	return &pb.RegisterAgentResponse{}, nil
}
