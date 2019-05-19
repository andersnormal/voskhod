package scheduler

import (
	"github.com/andersnormal/pkg/server"
	"github.com/andersnormal/voskhod/agent/config"
	"github.com/andersnormal/voskhod/server/nats"
)

type Scheduler interface {
	server.Listener
}

type scheduler struct {
	nats nats.Nats

	cfg  *config.Config
	opts *Opts
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Config *config.Config
}
