package config

import (
	"syscall"

	log "github.com/sirupsen/logrus"
)

const (
	// DefaultDockerReservedPort is the default port for the Docker engine
	// http://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml?search=docker
	DefaultDockerReservedPort = 2375

	// DefaultDockerReservedSSLPort is the default SSL port for the Docker engine
	DefaultDockerReservedSSLPort = 2376

	// DefaultLogLevel is the default logging level.
	DefaultLogLevel = log.WarnLevel

	// DefaultTermSignal is the signal to term the agent.
	DefaultTermSignal = syscall.SIGTERM

	// DefaultReloadSignal is the default signal for reload.
	DefaultReloadSignal = syscall.SIGHUP

	// DefaultKillSignal is the default signal for termination.
	DefaultKillSignal = syscall.SIGINT

	// DefaultVerbose is the default verbosity.
	DefaultVerbose = false

	// DefaultTimeout is the default time to configure the runtime
	DefaultTimeout = 60

	// DefaultNatsReadyTimeout is the default timeout to wait for NATS to become ready
	DefaultNatsReadyTimeout = 10

	// DefaultsNatsFilestoreDir is the default directory to persit NATS messages
	DefaultNatsFilestoreDir = "data"

	// DefaultNatsHTTPPort is the default http port NATS listen on
	DefaultNatsHTTPPort = 8223

	// DefaultNatsPort is the default port is listening on
	DefaultNatsPort = 4223
)

// New returns a new Config
func New() *Config {
	return &Config{
		Verbose:               DefaultVerbose,
		LogLevel:              DefaultLogLevel,
		ReloadSignal:          DefaultReloadSignal,
		TermSignal:            DefaultTermSignal,
		KillSignal:            DefaultKillSignal,
		Timeout:               DefaultTimeout,
		DockerReservedPort:    DefaultDockerReservedPort,
		DockerReservedSSLPort: DefaultDockerReservedSSLPort,
		NatsFilestoreDir:      DefaultNatsFilestoreDir,
		NatsReadyTimeout:      DefaultNatsReadyTimeout,
		NatsPort:              DefaultNatsPort,
		NatsHTTPPort:          DefaultNatsHTTPPort,
	}
}
