package etcd

import (
	"time"

	"go.etcd.io/etcd/embed"
)

const (
	defaultStartTimeout     = 2500 * time.Millisecond
	defaultNatsReadyTimeout = 10
	defaultNatsFilestoreDir = "data"
	defaultClusterID        = "voskhod"
)

// Start is starting the queue
func (s *server) Start() func() error {
	return func() error {
		var err error

		cfg := embed.NewConfig()

		etcd, err := embed.StartEtcd(cfg)
		if err != nil {
			return err
		}

		s.etcd = etcd
		err = <-s.etcd.Err()

		// noop
		return err
	}
}
