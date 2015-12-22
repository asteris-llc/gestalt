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

// Name indicates the tool name
const Name = "gestalt"

// Version indicates the tool version
const Version = "1.0.0"

var (
	rootCmd = &cobra.Command{
		Use:   Name,
		Short: "server and client root",
		Long:  "Gestalt is a wrapper around K/V stores. It provides type checks so that invalid values are not set, and to enable the user to get an overview of the keys in their store.",
	}

	markdownCmd = &cobra.Command{
		Use:    "__markdown",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			cobra.GenMarkdownTree(rootCmd, "docs/cli/")
		},
	}
)

func init() {
	cmd.ServerFlags()
	client.SchemaFlags()
	client.ValueFlags()

	flagsets := []*pflag.FlagSet{
		cmd.ServerCmd.Flags(),

		client.SchemaCmd.PersistentFlags(),
		client.SchemaSubmitCmd.Flags(),
		client.SchemaDeleteCmd.Flags(),

		client.ValueCmd.PersistentFlags(),
	}
	for _, flagset := range flagsets {
		if err := viper.BindPFlags(flagset); err != nil {
			log.Fatal(err)
		}
	}

	// set up command hierarchy
	rootCmd.AddCommand(
		cmd.ServerCmd,
		client.SchemaCmd,
		client.ValueCmd,
		markdownCmd,
	)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
