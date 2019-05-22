package etcd

import (
	"context"
	"net/url"

	s "github.com/andersnormal/pkg/server"

	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/embed"
)

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Dir    string
	LCUrls []url.URL
}

// New returns a new server
func New(opts ...Opt) s.Listener {
	options := new(Opts)

	s := new(server)
	s.cfg = embed.NewConfig()
	s.opts = options

	configure(s, opts...)

	s.logger = log.WithFields(log.Fields{})

	return s
}

// Start is starting the queue
func (s *server) Start(ctx context.Context) func() error {
	return func() error {
		etcd, err := embed.StartEtcd(s.cfg)
		if err != nil {
			return err
		}

		s.etcd = etcd
		err = <-s.etcd.Err()

		// noop
		return err
	}
}

// Stop is stopping the queue
func (s *server) Stop() error {
	s.log().Info("shutting down etcd...")

	if s.etcd != nil {
		s.etcd.Close()
	}

	return nil
}

// WithLCUrls ...
func WithLCUrls(urls []url.URL) func(o *Opts) {
	return func(o *Opts) {
		o.LCUrls = urls
	}
}

// WithDir ...
func WithDir(dir string) func(o *Opts) {
	return func(o *Opts) {
		o.Dir = dir
	}
}

func (s *server) log() *log.Entry {
	return s.logger
}

func configure(s *server, opts ...Opt) error {
	for _, o := range opts {
		o(s.opts)
	}

	s.cfg.Dir = s.opts.Dir
	s.cfg.LCUrls = s.opts.LCUrls

	return nil
}
