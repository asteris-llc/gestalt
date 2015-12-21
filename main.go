package main

import (
	"github.com/asteris-llc/gestalt/cmd"
	"github.com/asteris-llc/gestalt/cmd/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "gestalt",
}

func init() {
	cmd.ServerFlags()
	client.SchemaFlags()
	client.ValueFlags()

	flagsets := []*pflag.FlagSet{
		cmd.ServerCmd.Flags(),
		client.SchemaCmd.PersistentFlags(),
		client.SchemaSubmitCmd.Flags(),
		client.SchemaDeleteCmd.Flags(),
	}
	for _, flagset := range flagsets {
		if err := viper.BindPFlags(flagset); err != nil {
			log.Fatal(err)
		}
	}

	// set up command hierarchy
	rootCmd.AddCommand(cmd.ServerCmd)
	rootCmd.AddCommand(client.SchemaCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
