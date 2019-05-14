package cmd

import (
	"context"

	"github.com/andersnormal/pkg/server"
	// agent "github.com/andersnormal/voskhod/agent/run"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
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

	// // create agent and start
	// agent := agent.New(ctx, cfg)
	// g.Go(agent.Start())

	// wait for errors
	if err := s.Wait(); err != nil {
		root.logger.Error(err)
	}

	// noop
	return nil
}
