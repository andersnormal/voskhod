package nats

import (
	"context"
	"time"

	"github.com/andersnormal/voskhod/logger"
	"github.com/andersnormal/voskhod/server/config"

	natsd "github.com/nats-io/gnatsd/server"
	stand "github.com/nats-io/nats-streaming-server/server"
	"github.com/nats-io/nats-streaming-server/stores"

	log "github.com/sirupsen/logrus"
)

const (
	defaultNatsHTTPPort = 8223
	defaultNatsPort     = 4223
)

// New returns a new server
func New(cfg *config.Config, opts ...Opt) Nats {
	options := new(Opts)

	n := new(nats)
	n.opts = options
	n.cfg = cfg

	n.logger = log.WithFields(log.Fields{})

	configure(n, opts...)

	return n
}

// Stop is stopping the queue
func (n *nats) Stop() error {
	n.log().Info("shutting down nats...")

	if n.ss != nil {
		n.ss.Shutdown()
	}

	if n.ns != nil {
		n.ns.Shutdown()
	}

	return nil
}

const (
	defaultStartTimeout     = 2500 * time.Millisecond
	defaultNatsReadyTimeout = 10
	defaultClusterID        = "voskhod"
)

// Start is starting the queue
func (n *nats) Start(ctx context.Context) func() error {
	return func() error {
		var err error

		nopts := &natsd.Options{}
		nopts.HTTPPort = 8223
		nopts.Port = defaultNatsPort
		nopts.NoSigs = true

		n.ns = n.startNatsd(nopts) // wait for the Nats server to come available
		if !n.ns.ReadyForConnections(defaultNatsReadyTimeout * time.Second) {
			return NewError("could not start Nats server in %s seconds", defaultNatsReadyTimeout)
		}

		// verbose
		n.log().Infof("Started NATS server")

		// Get NATS Streaming Server default options
		opts := stand.GetDefaultOptions()
		opts.StoreType = stores.TypeFile
		opts.FilestoreDir = n.cfg.NatsFilestoreDir()
		opts.ID = defaultClusterID

		// set custom logger
		logger := logger.New()
		logger.SetLogger(n.log())
		opts.CustomLogger = logger

		// Do not handle signals
		opts.HandleSignals = false
		opts.EnableLogging = true
		opts.Debug = n.cfg.Verbose
		opts.Trace = n.cfg.Tracing

		// Now we want to setup the monitoring port for NATS Streaming.
		// We still need NATS Options to do so, so create NATS Options
		// using the NewNATSOptions() from the streaming server package.
		snopts := stand.NewNATSOptions()
		snopts.HTTPPort = 8222
		snopts.NoSigs = true

		// Now run the server with the streaming and streaming/nats options.
		n.ss, err = stand.RunServerWithOpts(opts, snopts)
		if err != nil {
			return err
		}

		// verbose
		n.log().Infof("Started cluster %s", n.ss.ClusterID())

		// wait for the server to be ready
		time.Sleep(defaultNatsReadyTimeout)

		// noop
		return nil
	}
}

func (n *nats) startNatsd(nopts *natsd.Options) *natsd.Server {
	// Create the NATS Server
	ns := natsd.New(nopts)

	// Start it as a go routine
	go ns.Start()

	return ns
}

func (n *nats) log() *log.Entry {
	return n.logger
}

func configure(n *nats, opts ...Opt) error {
	for _, o := range opts {
		o(n.opts)
	}

	return nil
}
