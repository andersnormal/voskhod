package etcd

import (
	"github.com/katallaxie/voskhod/server/config"

	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/embed"
)

type Server interface {
	// Start is starting the server
	Start() func() error
	// Stop is stopping the server
	Stop() func() error
}

type server struct {
	cfg  *config.Config
	etcd *embed.Etcd

	// logger instance
	logger *log.Entry
}
