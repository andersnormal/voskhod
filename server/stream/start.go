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

package stream

import (
	"log"
	"time"

	natsd "github.com/nats-io/gnatsd/server"
	stand "github.com/nats-io/nats-streaming-server/server"
	"github.com/nats-io/nats-streaming-server/stores"
)

const (
	defaultNatsReadyTimeout = 10
	defaultNatsFilestoreDir = "data"
)

// Start is starting the queue
func (s *Stream) Start() func() error {
	return func() error {
		var err error

		nopts := &natsd.Options{}
		nopts.HTTPPort = defaultNatsHTTPPort
		nopts.Port = defaultNatsPort
		nopts.NoSigs = true

		s.ns = s.startNatsd(nopts) // wait for the Nats server to come available
		if !s.ns.ReadyForConnections(defaultNatsReadyTimeout * time.Second) {
			return NewError("could not start Nats server in %s seconds", defaultNatsReadyTimeout)
		}

		log.Print("Started nats server")

		// Get NATS Streaming Server default options
		opts := stand.GetDefaultOptions()
		opts.StoreType = stores.TypeFile
		opts.FilestoreDir = defaultNatsFilestoreDir

		// Point to the NATS Server with host/port used above
		opts.NATSServerURL = "nats://localhost:4223"

		// Do not handle signals
		opts.HandleSignals = false

		// Now we want to setup the monitoring port for NATS Streaming.
		// We still need NATS Options to do so, so create NATS Options
		// using the NewNATSOptions() from the streaming server package.
		snopts := stand.NewNATSOptions()
		snopts.HTTPPort = 8222
		snopts.NoSigs = true

		// Now run the server with the streaming and streaming/nats options.
		s.ss, err = stand.RunServerWithOpts(opts, snopts)
		if err != nil {
			return err
		}

		// wait for the server
		time.Sleep(2500 * time.Millisecond)

		for {
			select {
			case <-s.ctx.Done():
				s.Stop()
				return err
			default:
			}
		}
	}
}

func (s *Stream) startNatsd(nopts *natsd.Options) *natsd.Server {

	// Create the NATS Server
	ns := natsd.New(nopts)

	// Start it as a go routine
	go ns.Start()

	return ns
}
