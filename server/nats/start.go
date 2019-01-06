package nats

import (
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
func (s *server) Start() func() error {
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

		s.log().Infof("Started nats server")

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

		// wait for the server to be ready
		time.Sleep(2500 * time.Millisecond)

		// noop
		return nil
	}
}

func (s *server) startNatsd(nopts *natsd.Options) *natsd.Server {

	// Create the NATS Server
	ns := natsd.New(nopts)

	// Start it as a go routine
	go ns.Start()

	return ns
}
