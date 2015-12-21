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

	// SchemaSubmitCmd submits a schema
	SchemaSubmitCmd = &cobra.Command{
		Use:   "submit schema.json",
		Short: "submit a schema",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("expected exactly one JSON file as an argument")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// read JSON
			body, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			var dest interface{}
			err = json.Unmarshal(body, &dest)
			if err != nil {
				return err
			}

			schema, err := app.LoadSchema(dest)
			if err != nil {
				return err
			}

			err = schema.Validate()
			if err != nil {
				return err
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
				return err
			}

			body, err = ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return err
			}

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				return fmt.Errorf("API returned \"%s\":\n\n%s", resp.Status, string(body))
			}

			fmt.Println(string(body))

			return nil
		},
	}

	// SchemaDeleteCmd deletes a schema
	SchemaDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a schema",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("expected exactly one schema name as an argument")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := client.Do(
				"DELETE",
				"/v1/schemas/"+args[0],
				map[string]interface{}{
					"deleteKeys": viper.Get("delete-keys"),
				},
				nil,
			)

			if err != nil {
				return err
			}

			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return err
			}

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				return fmt.Errorf("API returned \"%s\":\n\n%s", resp.Status, string(body))
			}

			if body != nil {
				fmt.Println(string(body))
			}

			return nil
		},
	}
)

// SchemaFlags sets up flags for the client commands
func SchemaFlags() {
	SchemaCmd.PersistentFlags().String("scheme", "http", "set the request scheme")
	SchemaCmd.PersistentFlags().String("host", "localhost:3000", "API hostname")
	SchemaCmd.PersistentFlags().Duration("timeout", 20*time.Second, "set the request timeout")

	SchemaSubmitCmd.Flags().Bool("set-defaults", true, "set defaults when submitting")

	SchemaDeleteCmd.Flags().Bool("delete-keys", false, "also delete configuration keys")

	// set up command hierarchy
	SchemaCmd.AddCommand(SchemaSubmitCmd, SchemaDeleteCmd)
}
