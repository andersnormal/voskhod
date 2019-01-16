package server

import (
	"context"
	// "fmt"
	"time"

	"github.com/katallaxie/voskhod/server/config"

	"github.com/andersnormal/pkg"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var _ Server = (*server)(nil)

// New is returning a new agent
func New(ctx context.Context, cfg *config.Config) Server {
	g, gtx := errgroup.WithContext(ctx)

	s := &server{
		cfg:    cfg,
		errCtx: gtx,
		errG:   g,
		logger: log.WithFields(log.Fields{}),
		ready:  pkg.NewReadyEvents(cfg.ReadyTimeout),
	}

	return s
}

// Ready is returning the wait signal for the server to be ready
func (s *server) Ready() error {
	return s.ready.Wait()
}

// Wait is returning the wait signal of the underlying errgroup
func (s *server) Wait() error {
	for {
		select {
		case <-s.errCtx.Done():
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
			defer cancel()

			g, _ := errgroup.WithContext(ctx)

			if s.api != nil {
				g.Go(s.shutdownAPI())
			}

			if s.nats != nil {
				g.Go(s.shutdownNats())
			}

			if s.etcd != nil {
				g.Go(s.shutdownEtcd())
			}

			if err := g.Wait(); err != nil {
				return err
			}

			return nil
		}
	}
}

func (s *server) config() *config.Config {
	return s.cfg
}

func (s *server) log() *log.Entry {
	return s.logger
}
