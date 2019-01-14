package etcd

// Stop is stopping the queue
func (s *server) Stop() func() error {
	return func() error {
		s.etcd.Close()

		return nil
	}
}
