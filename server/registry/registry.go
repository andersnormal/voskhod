package registry

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/satori/go.uuid"
)

const (
	RegisterAction   = "register"
	DeregisterAction = "deregister"
)

const (
	defaultClusterID  = "voskhod"
	defaultNATSURL    = "nats://localhost:4222"
	defaultWatchTopic = "voskhod.registry.watch"
)

func New(opts ...Option) Registry {
	options := &Options{
		TTL: time.Second * 25,
	}

	var n = new(registry)
	n.opts = options

	configure(n, opts...)

	return n
}

func (r *registry) Init(opts ...Option) error {
	return configure(r, opts...)
}

func (r *registry) Options() *Options {
	return r.opts
}

func (r *registry) Register(a *Agent, opts ...AgentOption) error {
	conn, err := r.getConn()
	if err != nil {
		return err
	}

	msg, err := json.Marshal(&Message{Action: RegisterAction, Agent: a})
	if err != nil {
		return err
	}

	return conn.Publish(r.watchTopic, msg)
}

func (r *registry) Deregister(a *Agent) error {
	conn, err := r.getConn()
	if err != nil {
		return err
	}

	msg, err := json.Marshal(&Message{Action: DeregisterAction, Agent: a})
	if err != nil {
		return err
	}

	return conn.Publish(r.watchTopic, msg)
}

func (r *registry) Watch(cb stan.MsgHandler) (stan.Subscription, error) {
	conn, err := r.getConn()
	if err != nil {
		return nil, err
	}

	return conn.Subscribe(r.watchTopic, cb, stan.DeliverAllAvailable())
}

func (r *registry) newConn() (stan.Conn, error) {
	opts := r.nopts
	opts.Secure = r.opts.Secure
	opts.TLSConfig = r.opts.TLSConfig

	clusterID := r.clusterID
	clientID := r.clientID
	if r.clientID == "" {
		clientID = uuid.NewV4().String()
	}

	// security
	if opts.TLSConfig != nil {
		opts.Secure = true
	}

	nc, err := opts.Connect()
	if err != nil {
		return nil, err
	}

	// more options
	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		fmt.Println(err)
	}

	return sc, err
}

func (r *registry) getConn() (stan.Conn, error) {
	r.Lock()
	defer r.Unlock()

	if r.conn != nil {
		return r.conn, nil
	}

	c, err := r.newConn()
	if err != nil {
		return nil, err
	}

	r.conn = c

	return r.conn, nil
}

func configure(r *registry, opts ...Option) error {
	for _, o := range opts {
		o(r.opts)
	}

	// mixin default opts
	nopts := nats.DefaultOptions

	watchTopic := defaultWatchTopic

	if r.opts.ClusterID == "" {
		r.opts.ClusterID = defaultClusterID
	}

	if len(r.opts.Addrs) == 0 {
		r.opts.Addrs = nopts.Servers
	}

	if !r.opts.Secure {
		r.opts.Secure = nopts.Secure
	}

	if r.opts.TLSConfig == nil {
		r.opts.TLSConfig = nopts.TLSConfig
	}

	r.clientID = r.opts.ClientID
	r.clusterID = r.opts.ClusterID

	r.nopts = nopts
	r.watchTopic = watchTopic

	return nil
}
