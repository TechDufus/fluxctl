/*
Copyright © 2024 TechDufus

The get subcommand is used to display one or many resources.
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
	apis "github.com/techdufus/fluxctl/pkg/apis/public"
)

var (
	appsCmdShort = "Display one or more apps."
	// getCmdLong  = ``
)

// NewCmdGet returns a new cobra command for `fluxctl get`
// fluxctl get displays one or many resources.
func NewCmdGetApps() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apps",
		Short: appsCmdShort,
		// Long:  getCmdLong,
		Run: func(cmd *cobra.Command, args []string) {
			data, err := apis.GetListAllApps()
			if err != nil {
				fmt.Errorf("Error getting apps: %v", err)
			}
			apis.FormatListAllApps(data.Data, cmd.Flag("output").Value.String())
		},
	}
	cmd.Flags().StringP("output", "o", "default", "Output format. One of: default, wide")
	return cmd
}
