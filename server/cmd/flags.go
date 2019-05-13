package cmd

import (
	"time"

	"github.com/andersnormal/voskhod/server/config"
	"github.com/spf13/cobra"
)

func addFlags(cmd *cobra.Command, cfg *config.Config) {
	// enable verbose output
	cmd.Flags().BoolVar(&cfg.Verbose, "verbose", config.DefaultVerbose, "enable verbose")

	// enable tracing output
	cmd.Flags().BoolVar(&cfg.Tracing, "tracing", config.DefaultVerbose, "enable tracing")

	// timeout for client operations
	cmd.Flags().DurationVar(&cfg.Timeout, "timeout", config.DefaultTimeout*time.Second, "timeout")

	// Host to listen on
	cmd.Flags().StringVar(&cfg.Host, "host", config.DefaultHost, "host")

	// API Port
	cmd.Flags().IntVar(&cfg.APIPort, "api-port", config.DefaultAPIPort, "api port")
}
