package nats

import (
	"time"

	"github.com/andersnormal/voskhod/server/config"

	s "github.com/andersnormal/pkg/server"
	natsd "github.com/nats-io/gnatsd/server"
	stand "github.com/nats-io/nats-streaming-server/server"
	log "github.com/sirupsen/logrus"
)

type Nats interface {
	s.Listener
}

type nats struct {
	cfg *config.Config
	ns  *natsd.Server
	ss  *stand.StanServer

	opts *Opts

	// logger instance
	logger *log.Entry
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	ID      string
	Timeout time.Duration
}
