package config

import (
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

// Config contains a configuration for Voskhod
type Config struct {
	// Verbose toggles the verbosity
	Verbose bool

	// LogLevel is the level with with to log for this config
	LogLevel log.Level

	// ReloadSignal
	ReloadSignal syscall.Signal

	// TermSignal
	TermSignal syscall.Signal

	// KillSignal
	KillSignal syscall.Signal

	// Timeout of the runtime
	Timeout time.Duration

	// DockerReservedPort of engine port
	DockerReservedPort int

	// DockerReservedSSLPort is the default SSL port for the Docker engine
	DockerReservedSSLPort int

	// NatsReadyTimeout is the timeout to wait for NATS to become ready
	NatsReadyTimeout time.Duration

	// NatsFilestoreDir is the directory to persit NATS messages
	NatsFilestoreDir string

	// NatsHTTPPort is the http port that NATS is listening on
	NatsHTTPPort int

	// NatsPort is the port NATS is listing on
	NatsPort int
}
