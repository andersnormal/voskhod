package server

import (
	"github.com/katallaxie/voskhod/server/nats"
)

func (s *server) ServeNats(n nats.Server) {
	g := s.errG

	s.nats = n

	g.Go(s.serveNats())
}

func (s *server) serveNats() func() error {
	return s.nats.Start()
}

func (s *server) shutdownNats() func() error {
	return s.nats.Stop()
}
