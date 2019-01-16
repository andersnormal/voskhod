package cmd

import (
	"time"

	"github.com/andersnormal/voskhod/agent/config"
	c "github.com/andersnormal/voskhod/config"
	"github.com/spf13/cobra"
)

func addFlags(cmd *cobra.Command, cfg *config.Config) {
	// enable verbose output
	cmd.Flags().BoolVar(&cfg.Verbose, "verbose", c.DefaultVerbose, "enable verbose output")

	// timeout for client operations
	cmd.Flags().DurationVar(&cfg.Timeout, "timeout", c.DefaultTimeout*time.Second, "timeout")
}
