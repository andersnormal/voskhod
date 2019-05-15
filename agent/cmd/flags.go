package cmd

import (
	"time"

	"github.com/andersnormal/voskhod/agent/config"
	"github.com/spf13/cobra"
)

func addFlags(cmd *cobra.Command, cfg *config.Config) {
	// enable verbose output
	cmd.Flags().BoolVar(&cfg.Verbose, "verbose", config.DefaultVerbose, "enable verbose output")

	// timeout for client operations
	cmd.Flags().DurationVar(&cfg.Timeout, "timeout", config.DefaultTimeout*time.Second, "timeout")

	// cluster id
	cmd.Flags().StringVar(&cfg.ClusterID, "cluster", cfg.ClusterID, "cluster")

	// nats addr
	cmd.Flags().StringVar(&cfg.NatsAddr, "addr", cfg.NatsAddr, "address")

	// name
	cmd.Flags().StringVar(&cfg.Name, "name", cfg.Name, "name")
}
