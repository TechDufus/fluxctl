/*
Copyright © 2024 TechDufus

The Config subcommand is used to change or view all aspects of the fluxconfig.
*/
package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configCmdShort = "Display merged fluxconfig settings or a specified fluxconfig file"
	configCmdLong  = `Display merged fluxconfig settings or a specified fluxconfig file.

  This command will display the merged fluxconfig settings or a specified fluxconfig file.`
)

func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: configCmdShort,
		Long:  configCmdLong,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("You must specify a command for config.")
			}
			cmd.Help()
		},
	}
	cmd.AddCommand(NewCmdConfigView())
	cmd.AddCommand(NewCmdConfigSetCluster())
	return cmd
}
