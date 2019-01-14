package server

import (
	"github.com/katallaxie/voskhod/server/etcd"
)

// ServeEtcd is starting etcd in the errgroup
func (s *server) ServeEtcd(e etcd.Server) {
	g := s.errG

	s.etcd = e

	g.Go(s.serveEtcd())
}

func (s *server) serveEtcd() func() error {
	return s.etcd.Start()
}

func (s *server) shutdownEtcd() func() error {
	return s.etcd.Stop()
}
