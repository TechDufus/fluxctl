/*
Copyright © 2024 TechDufus

The config view subcommand is used to view the merged fluxconfig settings or a specified fluxconfig file.
*/
package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	viewCmdShort = "Display merged fluxconfig settings or a specified fluxconfig file"
	viewCmdLong  = `Display merged fluxconfig settings or a specified fluxconfig file.

  This command will display the merged fluxconfig settings or a specified fluxconfig file.`
	viewCmdExample = `  # Specify a config file
  fluxctl config view --fluxconfig ~/.flux/config

  # View the default config file
  fluxctl config view`
)

// NewCmdConfigView returns a new cobra command for `fluxctl config view`
// fluxctl config view exports a specified fluxconfig file.
func NewCmdConfigView() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: viewCmdShort,
		Long:  viewCmdLong,
		Run: func(cmd *cobra.Command, args []string) {
			data, err := os.ReadFile(cmd.Flag("fluxconfig").Value.String())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(data))
		},
	}
	cmd.PersistentFlags().String("fluxconfig", "~/.flux/config", "The path to the fluxconfig file. Defaults to '~/.flux/config'.")
	return cmd
}
