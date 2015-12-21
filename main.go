package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/asteris-llc/gestalt/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "gestalt",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		setupLogging()
	},
}

func init() {
	// root and persistent flags
	rootCmd.PersistentFlags().String("log-level", "info", "one of debug, info, warn, error, or fatal")
	rootCmd.PersistentFlags().String("log-format", "text", "output format (text or json)")

	cmd.ServerFlags()

	flagsets := []*pflag.FlagSet{
		rootCmd.PersistentFlags(),
		cmd.ServerCmd.Flags(),
	}
	for _, flagset := range flagsets {
		if err := viper.BindPFlags(flagset); err != nil {
			logrus.WithField("error", err).Fatal("could not bind flags")
		}
	}

	// set up command hierarchy
	rootCmd.AddCommand(cmd.ServerCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
