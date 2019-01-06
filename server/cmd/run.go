package cmd

import (
	"context"
	"os"

	"github.com/katallaxie/voskhod/server/nats"
	server "github.com/katallaxie/voskhod/server/run"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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

	// log
	root.logger.Info("Starting Server ...")

	// create nats
	nats := nats.New(cfg)

	// create agent and start
	server := server.New(root.ctx, cfg)

	// start the API
	server.ServeAPI()
	// start the Nats
	server.ServeNats(nats)

	// wait for errors
	err = server.Wait()

	// noop
	return err
}
