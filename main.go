package main

import (
	"github.com/asteris-llc/gestalt/cmd"
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

	flagsets := []*pflag.FlagSet{
		rootCmd.PersistentFlags(),
		cmd.ServerCmd.Flags(),
	}
	for _, flagset := range flagsets {
		if err := viper.BindPFlags(flagset); err != nil {
			log.Fatal(err)
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
