package registry

import (
	"context"
	"sync"
	"time"

	"github.com/nats-io/go-nats"

	pb "github.com/katallaxie/voskhod/proto"
)

type Registry interface {
	// Options returns the configured options
	Options() Options
	// Init allows to init with new options
	Init(opts ...Option) error
}

type Agent struct {
}

type AgentOptions struct {
}

type AgentOption func(*AgentOptions)

type Message struct {
	Action string
	Agent  *Agent
}

type Option func(*Options)

type registry struct {
	addrs      []string
	queryTopic string // this should reflect back to the cluster
	watchTopic string

	sync.RWMutex
	conn      *nats.Conn
	agents    map[string][]*pb.Agent
	listeners map[string]chan bool

	opts Options
}

type Options struct {
	TTL time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
