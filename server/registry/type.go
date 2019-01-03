package registry

import (
	"sync"
	"time"

	"github.com/nats-io/go-nats"

	pb "github.com/katallaxie/voskhod/proto"
)

type agentRegistry struct {
	addrs      []string
	queryTopic string // this should reflect back to the cluster
	watchTopic string

	sync.RWMutex
	conn      *nats.Conn
	agents    map[string][]*pb.Agent
	listeners map[string]chan bool
}

type RegisterOptions struct {
	TTL time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
