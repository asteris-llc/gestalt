package cmd

import (
	"github.com/asteris-llc/gestalt/store"
	"github.com/asteris-llc/gestalt/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
)

// ServerCmd starts the server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// read config file
		configBytes, err := ioutil.ReadFile(viper.GetString("config"))
		if err != nil {
			return err
		}

		config, err := store.FromConfig(configBytes)
		if err != nil {
			return err
		}

		// run server with config file
		web.Run(viper.GetString("address"), config)
		return nil
	},
}

// ServerFlags binds flags to the server
func ServerFlags() {
	ServerCmd.Flags().String("address", ":3000", "address to serve on")
	ServerCmd.Flags().String("config", "gestalt.toml", "config file to read for the server")
}
