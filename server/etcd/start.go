package etcd

import (
	"go.etcd.io/etcd/embed"
)

// Start is starting the queue
func (s *server) Start() func() error {
	return func() error {
		var err error

		cfg := embed.NewConfig()
		cfg.Dir = s.cfg.EtcdFilestoreDir()

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
