package etcd

import (
	"go.etcd.io/etcd/embed"
	"net/url"
)

// Start is starting the queue
func (s *server) Start() func() error {
	return func() error {
		var err error

		u, err := url.Parse("unix:///var/run/etcd")
		if err != nil {
			fmt.Println("test")
			return err
		}

		cfg := embed.NewConfig()
		cfg.Dir = s.cfg.EtcdFilestoreDir()
		cfg.LCUrls = []url.URL{*u}

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
