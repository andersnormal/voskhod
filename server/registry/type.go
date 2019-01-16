package registry

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	"github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"

	pb "github.com/andersnormal/voskhod/proto"
)

type Watcher interface {
	// Next ...
	Next()
}

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
	Watch(cb stan.MsgHandler) (stan.Subscription, error)
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
	clusterID  string
	clientID   string
	queryTopic string // this should reflect back to the cluster
	watchTopic string

	nopts nats.Options
	opts  *Options

	sync.RWMutex
	conn stan.Conn

	agents    map[string][]*pb.Agent
	listeners map[string]chan bool
}

type Options struct {
	// Nats servers
	Addrs []string
	// Timeout for retry
	TTL time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
	// TLS Config
	TLSConfig *tls.Config
	// Secure
	Secure bool
	// ClusterID
	ClusterID string
	// CllientID
	ClientID string
}
