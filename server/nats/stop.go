package nats

import ()

// Stop is stopping the queue
func (s *server) Stop() func() error {
	return func() error {
		s.log().Info("shutting down nats...")

		if s.ss != nil {
			s.ss.Shutdown()
		}

		if s.ns != nil {
			s.ns.Shutdown()
		}

		return nil
	}
}
