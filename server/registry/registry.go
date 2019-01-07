package registry

import (
	"encoding/json"
	"sync"

	"github.com/nats-io/go-nats"
)

const (
	RegisterAction = "register"
)

func New(opts ...Option) Registry {
	var n = new(registry)

	configure(n, opts...)

	return n
}

func (r *registry) Init(opts ...Option) error {
	return configure(r, opts...)
}

func (r *registry) Options() Options {
	return r.opts
}

func (r *registry) Registry(a *Agent, opts ...AgentOption) error {

	b, err := json.Marshal(&Message{Action: RegisterAction, Agent: a})
	if err != nil {
		return err
	}

	return nil
}

func configure(r *registry, opts ...Option) error {
	for _, o := range opts {
		o(&r.opts)
	}

	return nil
}

// func (r *registry) register(agent *pb.Agent) error {
// 	var err error

// 	conn, err := a.getConn()
// 	if err != nil {
// 		return err
// 	}

// 	a.Lock()
// 	defer a.Unlock()

// 	r.agents[agent.Uuid] = addAgent(r.agents[agent.Uuid], []*registry.)
// }

// func (r *registry) Register(agent *pb.Agent, opts ...registry.RegisterOptions) error {

// }

// func (r *registry) newConn() (*nats.Conn, error) {
// 	opts := a.opts
// 	opts.Server = r.addrs
// 	opts.Secure = r.opts.Secure
// 	opts.TLSConfig = r.opts.TLSConfig

// 	return opts.Connect()
// }

// func (r *registry) getConn() (*nats.Conn, error) {
// 	var err error

// 	r.Lock()
// 	defer r.Unlock()

// 	if r.conn != nil {
// 		return r.conn, nil
// 	}

// 	c, err := r.newConn()
// 	if err := nil {
// 		return nil, err
// 	}
// 	r.conn = c

// 	return r.conn, err
// }
