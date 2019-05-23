package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/andersnormal/voskhod/agent/config"
	pb "github.com/andersnormal/voskhod/proto"
	"github.com/andersnormal/voskhod/server/nats"

	"github.com/nats-io/stan.go"
)

// New ...
func New(nats nats.Nats, opts ...Opt) Scheduler {
	options := new(Opts)

	s := new(scheduler)
	s.opts = options
	s.nats = nats

	s.exit = make(chan struct{}, 0)
	s.in = make(chan pb.Event)
	s.out = make(chan pb.Event)

	configure(s, opts...)

	return s
}

// WithConfig ...
func WithConfig(c *config.Config) func(o *Opts) {
	return func(o *Opts) {
		o.Config = c
	}
}

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

		s.conn = sc

		// start subscriptions
		s.run(s.publishEvents)
		s.run(s.subscribeEvents)

		// wait ... :)
		<-s.exit

		// Close connection
		sc.Close()

		return nil
	}
}

func (s *scheduler) subscribeEvents() error {
	// Simple Async Subscriber
	sub, _ := s.conn.Subscribe("events", func(m *stan.Msg) {
		fmt.Printf("Received an event: %s\n", string(m.Data))
	})

	// Unsubscribe
	defer sub.Unsubscribe()

	// wait for the ca
	<-s.exit

	// noop
	return nil
}

func (s *scheduler) publishEvents() error {
	ticker := time.Tick(5 * time.Second)

	for {
		select {
		case <-ticker:
			// Simple Synchronous Publisher
			s.conn.Publish("events", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming
		case <-s.exit:
			return nil
		}
	}
}

func (s *scheduler) run(fn func() error) {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()

		if err := fn(); err != nil {
			s.errOnce.Do(func() {
				s.err = err

				close(s.exit)
			})
		}
	}()
}

func configure(s *scheduler, opts ...Opt) error {
	for _, o := range opts {
		o(s.opts)
	}
	s.cfg = s.opts.Config

	return nil
}
