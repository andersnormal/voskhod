package nats

// Stop is stopping the queue
func (s *server) Stop() func() error {
	return func() error {
		s.ss.Shutdown()
		s.ns.Shutdown()

		return nil
	}
}
