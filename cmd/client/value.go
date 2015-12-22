package client

import (
	"bytes"
	"errors"
	"github.com/spf13/cobra"
	"log"
	"path"
	"time"
)

var (
	// ValueCmd is the wrapper command for values
	ValueCmd = &cobra.Command{
		Use:   "value",
		Short: "work with values",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			setupClient()
		},
	}

	// ValueListCmd lists values in a schema
	ValueListCmd = &cobra.Command{
		Use:   "list {schema name}",
		Short: "list the values in a schema",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("expected exactly one name as an argument")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := client.Do(
				"GET",
				path.Join("/v1/schemas", args[0], "values"),
				map[string]interface{}{},
				nil,
			)
			if err != nil {
				log.Fatal(err)
			}

			client.HandleResponse(resp)
		},
	}

	// ValueWriteCmd writes a value
	ValueWriteCmd = &cobra.Command{
		Use:   "write {schema name} {value name} {value}",
		Short: "write a value",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 3 {
				return errors.New("exected exactly two arguments: schema name, value name, and value")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := client.Do(
				"PUT",
				path.Join("/v1/schemas", args[0], "values", args[1]),
				map[string]interface{}{},
				bytes.NewBufferString(args[2]),
			)
			if err != nil {
				log.Fatal(err)
			}

			client.HandleResponse(resp)
		},
	}

	// ValueShowCmd shows a value
	ValueShowCmd = &cobra.Command{
		Use:   "show {schema name} {value name}",
		Short: "show a single value in a schema",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("exected exactly two arguments: schema name and value name")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := client.Do(
				"GET",
				path.Join("/v1/schemas", args[0], "values", args[1]),
				map[string]interface{}{},
				nil,
			)
			if err != nil {
				log.Fatal(err)
			}

			client.HandleResponse(resp)
		},
	}

	// ValueDeleteCmd deletes a value
	ValueDeleteCmd = &cobra.Command{
		Use:   "delete {schema name} {value name}",
		Short: "delete a single value in a schema",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("exected exactly two arguments: schema name and value name")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := client.Do(
				"DELETE",
				path.Join("/v1/schemas", args[0], "values", args[1]),
				map[string]interface{}{},
				nil,
			)
			if err != nil {
				log.Fatal(err)
			}

			client.HandleResponse(resp)
		},
	}
)

// ValueFlags sets up flags for the client commands
func ValueFlags() {
	ValueCmd.PersistentFlags().String("scheme", "http", "set the request scheme")
	ValueCmd.PersistentFlags().String("host", "localhost:3000", "API hostname")
	ValueCmd.PersistentFlags().Duration("timeout", 20*time.Second, "set the request timeout")
	ValueCmd.PersistentFlags().Bool("pretty", true, "pretty-print responses")

	// set up command hierarchy
	ValueCmd.AddCommand(
		ValueListCmd,
		ValueWriteCmd,
		ValueShowCmd,
		ValueDeleteCmd,
	)
}
