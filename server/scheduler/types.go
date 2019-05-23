package scheduler

import (
	"sync"

	"github.com/andersnormal/pkg/server"
	"github.com/andersnormal/voskhod/agent/config"
	pb "github.com/andersnormal/voskhod/proto"
	"github.com/andersnormal/voskhod/server/nats"

	"github.com/nats-io/stan.go"
)

type Scheduler interface {
	server.Listener
}

type scheduler struct {
	nats nats.Nats
	conn stan.Conn

	in  chan pb.Event
	out chan pb.Event

	errOnce sync.Once
	err     error
	wg      sync.WaitGroup

	exit chan struct{}

	cfg  *config.Config
	opts *Opts
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Config *config.Config
}
