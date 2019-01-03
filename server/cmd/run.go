package cmd

import (
	"context"
	"os"

	server "github.com/katallaxie/voskhod/server/run"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	log "github.com/sirupsen/logrus"
)

func runE(cmd *cobra.Command, args []string) error {
	var err error
	var root = new(root)

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

	// log
	root.logger.Info("Starting Server ...")

	// create agent and start
	server := server.New(cfg)
	g.Go(server.Start(ctx))

	// wait for errors
	err = g.Wait()

	// again, wait exit
	<-root.exit

	// noop
	return err
}
