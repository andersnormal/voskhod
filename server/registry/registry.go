package registry

import (
	"encoding/json"
	"time"

	"github.com/katallaxie/voskhod/utils"

	"github.com/nats-io/go-nats"
)

const (
	RegisterAction   = "register"
	DeregisterAction = "deregister"
)

const (
	defaultWatchTopic = "voskhod.registry.watch"
)

func New(opts ...Option) Registry {
	options := &Options{
		TTL: time.Millisecond * 100,
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

func (r *registry) Watch(opts ...utils.WatchOpt) (utils.Watcher, error) {
	conn, err := r.getConn()
	if err != nil {
		return nil, err
	}

	sub, err := conn.SubscribeSync(r.watchTopic)
	if err != nil {
		return nil, err
	}

	wo := &utils.WatchOpts{}
	for _, o := range opts {
		o(wo)
	}

	w := utils.NewWatcher(wo, sub)

	return w, nil
}

func (r *registry) newConn() (*nats.Conn, error) {
	opts := r.nopts
	opts.Servers = r.addrs
	opts.Secure = r.opts.Secure
	opts.TLSConfig = r.opts.TLSConfig

	if opts.TLSConfig != nil {
		opts.Secure = true
	}

	return opts.Connect()
}

func (r *registry) getConn() (*nats.Conn, error) {
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

	nopts := nats.GetDefaultOptions()

	watchTopic := defaultWatchTopic

	if len(r.opts.Addrs) == 0 {
		r.opts.Addrs = nopts.Servers
	}

	if r.opts.TLSConfig == nil {
		r.opts.TLSConfig = nopts.TLSConfig
	}

	r.addrs = r.opts.Addrs
	r.nopts = nopts
	r.watchTopic = watchTopic

	return nil
}
