package stream

import (
	"github.com/andersnormal/pkg/server"
	"github.com/andersnormal/voskhod/agent/config"

	"github.com/nats-io/stan.go"
)

type Stream interface {
	server.Listener
}

type stream struct {
	done chan error

	sc stan.Conn

	cfg  *config.Config
	opts *Opts
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Config *config.Config
	Nats   stan.Conn
}
