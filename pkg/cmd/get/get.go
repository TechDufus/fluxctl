/*
Copyright © 2024 TechDufus

The get subcommand is used to display one or many resources.
*/
package get

import (
	"github.com/spf13/cobra"
)

var (
	getCmdShort = "Display one or many resources."
	getCmdLong  = `Prints a table of the most important information about the specified resources.
Use "fluxctl explain <resource>" for more information on a specific resource.

See 'fluxctl get -h' for help and examples.`
)

// NewCmdGet returns a new cobra command for `fluxctl get`
// fluxctl get displays one or many resources.
func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: getCmdShort,
		Long:  getCmdLong,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	return cmd
}
