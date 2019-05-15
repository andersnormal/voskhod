package stream

import (
	"context"
	"fmt"
	"time"

	"github.com/andersnormal/voskhod/agent/config"

	"github.com/nats-io/stan.go"
)

// New returns a new server
func New(opts ...Opt) Stream {
	options := new(Opts)

	s := new(stream)
	s.opts = options
	s.msg = make(chan *stan.Msg, 64)

	configure(s, opts...)

	return s
}

// WithConfig ...
func WithConfig(c *config.Config) func(o *Opts) {
	return func(o *Opts) {
		o.Config = c
	}
}

// WithNats ...
func WithNats(sc stan.Conn) func(o *Opts) {
	return func(o *Opts) {
		o.Nats = sc
	}
}

// Stop ...
func (s *stream) Stop() error {
	return nil
}

// Start is starting the queue
func (s *stream) Start(ctx context.Context) func() error {
	return func() error {
		// Simple Async Subscriber
		sub, err := s.sc.Subscribe("foo", func(m *stan.Msg) {
			fmt.Printf("Received a message: %s\n", string(m.Data))
		}, stan.DeliverAllAvailable())

		if err != nil {
			return err
		}

		time.Sleep(time.Minute * 5)

		// Unsubscribe
		sub.Unsubscribe() // Close connection
		s.sc.Close()

		return nil
	}
}

func configure(s *stream, opts ...Opt) error {
	for _, o := range opts {
		o(s.opts)
	}

	s.sc = s.opts.Nats
	s.cfg = s.opts.Config

	return nil
}
