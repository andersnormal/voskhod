// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"context"
	"os"

	agent "github.com/katallaxie/voskhod/agent/run"
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
