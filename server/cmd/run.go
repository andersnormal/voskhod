package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/katallaxie/voskhod/server/nats"
	"github.com/katallaxie/voskhod/server/registry"
	server "github.com/katallaxie/voskhod/server/run"

	stan "github.com/nats-io/go-nats-streaming"
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

	opts := []registry.Option{}
	agents := registry.New(opts...)

	go func() error {
		time.Sleep(2500 * time.Millisecond)

		sub, err := agents.Watch(func(msg *stan.Msg) {
			fmt.Println("ttest")
			fmt.Println(msg)
		})

		if err != nil {
			fmt.Println("here")
			fmt.Println(err)

			return err
		}

		fmt.Println("here2")

		// wait for context
		time.Sleep(time.Second * 3600)

		sub.Close()

		return err
	}()

	// wait for the server to be ready
	time.Sleep(3500 * time.Millisecond)

	agents.Register(&registry.Agent{Name: "test"})

	// wait for errors
	err = server.Wait()

	// noop
	return err
}
