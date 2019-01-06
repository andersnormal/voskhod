package cmd

import (
	"fmt"
	"os"

	"github.com/katallaxie/voskhod/server/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg *config.Config
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "voskhod",
	Short: "Distributed scheduler",
	Long:  `Not yet`,
	RunE:  runE,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// init config
	cfg = config.New()

	// initialize cobra
	cobra.OnInitialize(initConfig)

	// adding flags
	addFlags(RootCmd, cfg)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	// setup the logging behavior
	setupLog(cfg)
}
