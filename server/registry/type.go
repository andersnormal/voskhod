package registry

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	"github.com/nats-io/go-nats"

	pb "github.com/katallaxie/voskhod/proto"
	"github.com/katallaxie/voskhod/utils"
)

type Registry interface {
	// Options returns the configured options
	Options() *Options
	// Init allows to init with new options
	Init(opts ...Option) error
	// Register allows to register an agent
	Register(a *Agent, opts ...AgentOption) error
	// Deregister allows to deregister an agent
	Deregister(a *Agent) error
	// Watch updates to the registry
	Watch(opts ...utils.WatchOpt) (utils.Watcher, error)
}

type Agent struct {
	Name string
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

	nopts nats.Options
	opts  *Options

	sync.RWMutex
	conn      *nats.Conn
	agents    map[string][]*pb.Agent
	listeners map[string]chan bool
}

type Options struct {
	TTL time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
	// Secure
	Secure bool
	// TLSConfig
	TLSConfig *tls.Config
	// NATS Addresses
	Addrs []string
}
