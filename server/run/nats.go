package server

import (
	"github.com/katallaxie/voskhod/server/nats"
)

// ServeNats is starting NATs in the errgroup
func (s *server) ServeNats(n nats.Server) {
	g := s.errG

	s.nats = n

	g.Go(s.serveNats())
}

func (s *server) serveNats() func() error {
	s.ready.Register(&NatsReady{})
	return s.nats.Start()
}

func (s *server) shutdownNats() func() error {
	return s.nats.Stop()
}
