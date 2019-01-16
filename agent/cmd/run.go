package cmd

import (
	"context"
	"os"

	agent "github.com/andersnormal/voskhod/agent/run"
	"github.com/andersnormal/voskhod/agent/stream"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	log "github.com/sirupsen/logrus"
)

func runE(cmd *cobra.Command, args []string) error {
	var err error
	var root = new(root)

	// ready
	ready := make(chan bool)
	defer close(ready)

	// init logger
	root.logger = log.WithFields(log.Fields{})

	// create sys channel
	root.sys = make(chan os.Signal, 1)
	root.exit = make(chan int, 1)

	// create root context
	root.ctx, root.cancel = context.WithCancel(context.Background())

	// watch syscalls and cancel upon need
	go root.watchSignals(cfg)

	// create errgroup
	g, ctx := errgroup.WithContext(root.ctx)

	// create agent and start
	stream := stream.New(ctx, cfg)
	g.Go(stream.Start(ready))

	// wait for the next
	<-ready

	// log
	root.logger.Info("Starting Agent ...")

	// create agent and start
	agent := agent.New(ctx, cfg)
	g.Go(agent.Start())

	// wait for errors
	err = g.Wait()

	// again, wait exit
	<-root.exit

	// noop
	return err
}
