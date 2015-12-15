package cmd

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServerCmd starts the server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.WithField("address", viper.GetString("address")).Info("should have listened")
		return errors.New("not implemented")
	},
}

// ServerFlags binds flags to the server
func ServerFlags() {
	ServerCmd.Flags().String("address", ":3000", "address to serve on")
}
