package config

import (
	"fmt"
	"syscall"

	log "github.com/sirupsen/logrus"
)

const (
	// DefaultGrpcAddr is the default address gRPC server
	DefaultGrpcAddr = "localhost:50051"

	// DefaultLogFormat is the default format of the logging.
	// The default is to log to JSON.
	DefaultLogFormat = "json"

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

	// DefaultHost to listen on
	DefaultHost = ""

	// DefaultAPIPort is the default port for API
	DefaultAPIPort = 8888
)

// New returns a new Config
func New() *Config {
	return &Config{
		Verbose:      DefaultVerbose,
		LogLevel:     DefaultLogLevel,
		LogFormat:    DefaultLogFormat,
		ReloadSignal: DefaultReloadSignal,
		TermSignal:   DefaultTermSignal,
		KillSignal:   DefaultKillSignal,
		Timeout:      DefaultTimeout,
		Host:         DefaultHost,
		APIPort:      DefaultAPIPort,
	}
}

// APIListener returns the listener for API
func (c *Config) APIListener() string {
	return fmt.Sprintf("%s:%d", c.Host, c.APIPort)
}
