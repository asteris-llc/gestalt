package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"time"
)

var (
	// SchemaCmd is the wrapper command for schemas
	SchemaCmd = &cobra.Command{
		Use:   "schema",
		Short: "work with schemas",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			setupClient()
		},
	}

	// SchemaListCmd lists schemas
	SchemaListCmd = &cobra.Command{
		Use:   "list",
		Short: "list schemas",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := client.Do(
				"GET",
				"/v1/schemas",
				map[string]interface{}{},
				nil,
			)
			if err != nil {
				log.Fatal(err)
			}

			client.HandleResponse(resp)
		},
	}

	// SchemaSubmitCmd submits a schema
	SchemaSubmitCmd = &cobra.Command{
		Use:   "submit {schema.json}",
		Short: "submit a schema",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("expected exactly one JSON file as an argument")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// read JSON
			body, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}

			var dest interface{}
			err = json.Unmarshal(body, &dest)
			if err != nil {
				log.Fatal(err)
			}

			schema, err := app.LoadSchema(dest)
			if err != nil {
				log.Fatal(err)
			}

			err = schema.Validate()
			if err != nil {
				log.Fatal(err)
			}

			resp, err := client.Do(
				"POST",
				"/v1/schemas",
				map[string]interface{}{
					"setDefaults": viper.Get("set-defaults"),
				},
				bytes.NewBuffer(body),
			)
			if err != nil {
				log.Fatal(err)
			}

			client.HandleResponse(resp)
		},
	}

	// SchemaShowCmd shows a schema
	SchemaShowCmd = &cobra.Command{
		Use:   "show {name}",
		Short: "show a schema",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("expected exactly one name as an argument")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := client.Do(
				"GET",
				"/v1/schemas/"+args[0],
				map[string]interface{}{},
				nil,
			)
			if err != nil {
				log.Fatal(err)
			}

			client.HandleResponse(resp)
		},
	}

	// SchemaDeleteCmd deletes a schema
	SchemaDeleteCmd = &cobra.Command{
		Use:   "delete {name}",
		Short: "delete a schema",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("expected exactly one schema name as an argument")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := client.Do(
				"DELETE",
				"/v1/schemas/"+args[0],
				map[string]interface{}{
					"deleteKeys": viper.Get("delete-keys"),
				},
				nil,
			)
			if err != nil {
				log.Fatal(err)
			}

			client.HandleResponse(resp)
		},
	}

	// SchemaValidateCmd submits a schema
	SchemaValidateCmd = &cobra.Command{
		Use:   "validate {schema.json}",
		Short: "validate a schema",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("expected exactly one JSON file as an argument")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// read JSON
			body, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}

			var dest interface{}
			err = json.Unmarshal(body, &dest)
			if err != nil {
				log.Fatal(err)
			}

			schema, err := app.LoadSchema(dest)
			if err != nil {
				log.Fatal(err)
			}

			err = schema.Validate()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("OK")
		},
	}
)

// SchemaFlags sets up flags for the client commands
func SchemaFlags() {
	SchemaCmd.PersistentFlags().String("scheme", "http", "set the request scheme")
	SchemaCmd.PersistentFlags().String("host", "localhost:3000", "API hostname")
	SchemaCmd.PersistentFlags().Duration("timeout", 20*time.Second, "set the request timeout")
	SchemaCmd.PersistentFlags().Bool("pretty", true, "pretty-print responses")

	SchemaSubmitCmd.Flags().Bool("set-defaults", true, "set defaults when submitting")

	SchemaDeleteCmd.Flags().Bool("delete-keys", false, "also delete configuration keys")

	// set up command hierarchy
	SchemaCmd.AddCommand(
		SchemaListCmd,
		SchemaShowCmd,
		SchemaSubmitCmd,
		SchemaDeleteCmd,
		SchemaValidateCmd,
	)
}
