package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "plan",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		setupLogging()
	},
}

func init() {
	// root and persistent flags
	rootCmd.PersistentFlags().String("log-level", "info", "one of debug, info, warn, error, or fatal")
	rootCmd.PersistentFlags().String("log-format", "text", "output format (text or json)")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		logrus.WithField("error", err).Fatal("could not bind flags")
	}
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		logrus.WithField("error", err).Fatal("exited with error")
	}
}
