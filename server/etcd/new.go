package etcd

import (
	"github.com/andersnormal/voskhod/server/config"

	log "github.com/sirupsen/logrus"
)

const (
	defaultNatsHTTPPort = 8223
	defaultNatsPort     = 4223
)

// New returns a new server
func New(cfg *config.Config) Server {
	var s = new(server)

	s.cfg = cfg
	s.logger = log.WithFields(log.Fields{})

	return s
}

func (s *server) log() *log.Entry {
	return s.logger
}
