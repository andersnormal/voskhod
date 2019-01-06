package nats

import (
	"github.com/katallaxie/voskhod/server/config"

	natsd "github.com/nats-io/gnatsd/server"
	stand "github.com/nats-io/nats-streaming-server/server"
	log "github.com/sirupsen/logrus"
)

type Server interface {
	// Start is starting the server
	Start() func() error
	// Stop is stopping the server
	Stop() func() error
}

type server struct {
	cfg *config.Config
	ns  *natsd.Server
	ss  *stand.StanServer

	// logger instance
	logger *log.Entry
}
