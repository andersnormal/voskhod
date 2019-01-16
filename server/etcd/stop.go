package etcd

// Stop is stopping the queue
func (s *server) Stop() func() error {
	return func() error {
		s.log().Info("shutting down etcd...")

		if s.etcd != nil {
			s.etcd.Close()
		}

		return nil
	}
}
