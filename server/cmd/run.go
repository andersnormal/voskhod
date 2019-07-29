package cmd

import (
	"context"
	"time"

	"github.com/andersnormal/voskhod/pkg/nats"

	"github.com/andersnormal/pkg/server"
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
	root.logger.Info("starting server...")

	// create nats ... as singular instance right now
	n := nats.New(
		nats.WithDebug(),
		nats.WithVerbose(),
		nats.WithDataDir(cfg.NatsFilestoreDir()),
		nats.WithID("voskhod"),
		nats.WithTimeout(1*time.Millisecond),
	)
	s.Listen(n, true)

	// u, err := url.Parse("http://localhost:2379")
	// if err != nil {
	// 	return err
	// }

	// create embed etcd
	// e := etcd.New(
	// 	etcd.WithDir(cfg.DataDir),
	// 	etcd.WithLCUrls([]url.URL{*u}),
	// )
	// s.Listen(e, false)

	// // wait for etcd to become available
	// time.Sleep(5 * time.Second)

	// // create agent and start
	// sched := scheduler.New(nats)
	// s.Listen(sched, false)

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
