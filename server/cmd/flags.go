package cmd

import (
	"time"

	"github.com/katallaxie/voskhod/server/config"
	"github.com/spf13/cobra"
)

func addFlags(cmd *cobra.Command, cfg *config.Config) {
	// enable verbose output
	cmd.Flags().BoolVar(&cfg.Verbose, "verbose", config.DefaultVerbose, "enable verbose output")

	// timeout for client operations
	cmd.Flags().DurationVar(&cfg.Timeout, "timeout", config.DefaultTimeout*time.Second, "timeout")
}
