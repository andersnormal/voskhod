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

	// Tracing toggles the tracing
	Tracing bool

	// LogFormat is the format of the logger to use
	LogFormat string

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

	// Host is the host to listen on
	Host string

	// APIPort is the port for API
	APIPort int

	// DataDir is the directory for etcd
	DataDir string

	// NatsDataDir is the directory for Nats
	NatsDataDir string
}
