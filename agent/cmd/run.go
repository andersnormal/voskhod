package cmd

import (
	"context"

	"github.com/andersnormal/pkg/server"
	"github.com/andersnormal/voskhod/agent/stream"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runE(cmd *cobra.Command, args []string) error {
	// create new root command
	root := new(root)

	// init logger
	root.logger = log.WithFields(log.Fields{
		"verbose": cfg.Verbose,
	})

	// create new root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create server
	s := server.NewServer(ctx)

	// log
	root.logger.Info("starting Agent ...")

	nc, err := nats.Connect(cfg.NatsAddr)
	if err != nil {
		return err
	}

	// Connect to a server
	sc, err := stan.Connect("voskhod", "test", stan.NatsConn(nc))
	if err != nil {
		return err
	}

	// Simple Publisher
	if err := sc.Publish("foo", []byte("Hello World")); err != nil {
		return err
	}

	// create stream
	ss := stream.New(
		stream.WithConfig(cfg),
		stream.WithNats(sc),
	)
	s.Listen(ss)

	// wait for errors
	if err := s.Wait(); err != nil {
		root.logger.Error(err)
	}

	// noop
	return nil
}
