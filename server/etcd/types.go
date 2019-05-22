package etcd

import (
	s "github.com/andersnormal/pkg/server"

	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/embed"
)

var _ s.Listener = (*server)(nil)

type server struct {
	etcd *embed.Etcd
	cfg  *embed.Config

	logger *log.Entry

	opts *Opts
}
