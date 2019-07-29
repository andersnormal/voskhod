package stream

import (
	"context"
	"fmt"

	"github.com/andersnormal/voskhod/agent/config"
	pb "github.com/andersnormal/voskhod/proto"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/stan.go"
)

// New returns a new server
func New(opts ...Opt) Stream {
	options := new(Opts)

	s := new(stream)
	s.opts = options
	s.done = make(chan error)

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
func (s *stream) Start(ctx context.Context, ready func()) func() error {
	return func() error {
		sub, err := s.sc.Subscribe("foo", s.handleMessage, stan.DurableName(s.cfg.Name), stan.MaxInflight(1))
		if err != nil {
			return err
		}

		// Unsubscribe
		defer sub.Unsubscribe() // Close connection
		defer s.sc.Close()

		// wait for something to end this
		err = <-s.done

		return err
	}
}

// handleMessages ...
func (s *stream) handleMessage(m *stan.Msg) {
	event := new(pb.Event)
	if err := proto.Unmarshal(m.Data, event); err != nil {
		// todo: reply with error
	}

	// if err := m.Ack(); err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Printf("Received a message: %s\n", event.GetEvent())
}

func configure(s *stream, opts ...Opt) error {
	for _, o := range opts {
		o(s.opts)
	}

	s.sc = s.opts.Nats
	s.cfg = s.opts.Config

	return nil
}
