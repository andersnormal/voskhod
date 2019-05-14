package cmd

import (
	"context"
	"time"

	"github.com/andersnormal/voskhod/server/nats"

	"github.com/andersnormal/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runE(cmd *cobra.Command, args []string) error {
	// create new root command
	root := new(root)

	// setup folders
	// if err = mkdirDataFolder(cfg); err != nil {
	// 	return err
	// }

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
	root.logger.Info("starting server...")

	// // create nats
	nats := nats.New(
		cfg,
		nats.WithID("voskhod"),
		nats.WithTimeout(2500*time.Millisecond),
	)
	s.Listen(nats)

	// // create agent and start
	// server := server.New(root.ctx, cfg)

	// // start the API
	// server.ServeAPI()
	// // start the Nats
	// server.ServeNats(nats)
	// // start etcd
	// server.ServeEtcd(etcd)

	// wait for errors
	if err := s.Wait(); err != nil {
		root.logger.Error(err)
	}

	// noop
	return nil
}
