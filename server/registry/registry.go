package registry

import (
	"sync"

	"github.com/nats-io/go-nats"

	pb "github.com/katallaxie/voskhod/proto"
)

func (a *agentRegistry) register(agent *pb.Agent) error {
	var err error

	conn, err := a.getConn()
	if err != nil {
		return err
	}

	a.Lock()
	defer a.Unlock()

	a.agents[agent.Uuid] = addAgent(a.agents[agent.Uuid], []*registry.)
}

func (a *agentRegistry) Register(agent *pb.Agent, opts ...registry.RegisterOptions) error {

}

func (a *agentRegistry) newConn() (*nats.Conn, error) {
	opts := a.opts
	opts.Server = a.addrs
	opts.Secure = a.opts.Secure
	opts.TLSConfig = a.opts.TLSConfig

	return opts.Connect()
}

func (a *agentRegistry) getConn() (*nats.Conn, error) {
	var err error

	a.Lock()
	defer a.Unlock()

	if a.conn != nil {
		return a.conn, nil
	}

	c, err := a.newConn()
	if err := nil {
		return nil, err
	}
	a.conn = c

	return a.conn, err
}
