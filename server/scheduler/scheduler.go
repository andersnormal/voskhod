package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/andersnormal/voskhod/agent/config"
	"github.com/andersnormal/voskhod/server/nats"

	"github.com/nats-io/stan.go"
)

// New ...
func New(nats nats.Nats, opts ...Opt) Scheduler {
	options := new(Opts)

	s := new(scheduler)
	s.opts = options
	s.nats = nats

	configure(s, opts...)

	return s
}

// WithConfig ...
func WithConfig(c *config.Config) func(o *Opts) {
	return func(o *Opts) {
		o.Config = c
	}
}

// // WithNats ...
// func WithNats(sc stan.Conn) func(o *Opts) {
// 	return func(o *Opts) {
// 		o.Nats = sc
// 	}
// }

// Stop ...
func (s *scheduler) Stop() error {
	return nil
}

// Start ...
func (s *scheduler) Start(ctx context.Context) func() error {
	return func() error {
		// wait for the server to be ready
		time.Sleep(5 * time.Second)

		// connect to cluster
		sc, err := stan.Connect(s.nats.ClusterID(), "server")
		if err != nil {
			return err
		}

		// Simple Synchronous Publisher
		sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

		// Simple Async Subscriber
		sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
			fmt.Printf("Received a message: %s\n", string(m.Data))
		})

		// Unsubscribe
		sub.Unsubscribe()

		// Close connection
		sc.Close()

		return nil
	}
}

// handle ...

func configure(s *scheduler, opts ...Opt) error {
	for _, o := range opts {
		o(s.opts)
	}
	s.cfg = s.opts.Config

	return nil
}
