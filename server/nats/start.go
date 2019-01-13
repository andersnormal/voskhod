package nats

import (
	"time"

	"github.com/katallaxie/voskhod/logger"

	natsd "github.com/nats-io/gnatsd/server"
	stand "github.com/nats-io/nats-streaming-server/server"
	"github.com/nats-io/nats-streaming-server/stores"
)

const (
	defaultStartTimeout     = 2500 * time.Millisecond
	defaultNatsReadyTimeout = 10
	defaultNatsFilestoreDir = "data"
	defaultClusterID        = "voskhod"
)

// Start is starting the queue
func (s *server) Start(ready func()) func() error {
	return func() error {
		var err error

		nopts := &natsd.Options{}
		nopts.HTTPPort = 8223
		nopts.Port = defaultNatsPort
		nopts.NoSigs = true

		s.ns = s.startNatsd(nopts) // wait for the Nats server to come available
		if !s.ns.ReadyForConnections(defaultNatsReadyTimeout * time.Second) {
			return NewError("could not start Nats server in %s seconds", defaultNatsReadyTimeout)
		}

		// verbose
		s.log().Infof("Started NATS server")

		// Get NATS Streaming Server default options
		opts := stand.GetDefaultOptions()
		opts.StoreType = stores.TypeFile
		opts.FilestoreDir = defaultNatsFilestoreDir
		opts.ID = defaultClusterID

		// set custom logger
		logger := logger.New()
		logger.SetLogger(s.log())
		opts.CustomLogger = logger

		// Do not handle signals
		opts.HandleSignals = false
		opts.EnableLogging = true
		opts.Debug = s.cfg.Verbose
		opts.Trace = s.cfg.Tracing

		// Now we want to setup the monitoring port for NATS Streaming.
		// We still need NATS Options to do so, so create NATS Options
		// using the NewNATSOptions() from the streaming server package.
		snopts := stand.NewNATSOptions()
		snopts.HTTPPort = 8222

		// Now run the server with the streaming and streaming/nats options.
		s.ss, err = stand.RunServerWithOpts(opts, snopts)
		if err != nil {
			return err
		}

		// verbose
		s.log().Infof("Started cluster %s", s.ss.ClusterID())

		// wait for the server to be ready
		time.Sleep(defaultNatsReadyTimeout)

		ready()

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
