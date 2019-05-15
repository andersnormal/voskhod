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

	// ClusterID
	ClusterID string

	// NatsAddr ...
	NatsAddr string

	// Name
	Name string
}
