/*
Copyright © 2024 TechDufus

The config setCluster subcommand is used to setCluster the merged fluxconfig settings or a specified fluxconfig file.
*/
package config

import (
	"fmt"

	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
)

type setClusterOptions struct {
	// configPath            config.ConfigAccess
	name                  string
	server                cliflag.StringFlag
	tlsServerName         cliflag.StringFlag
	insecureSkipTLSVerify cliflag.Tristate
	certificateAuthority  cliflag.StringFlag
	embedCAData           cliflag.Tristate
	proxyURL              cliflag.StringFlag
}

var (
	setClusterCmdShort = "Set a cluster in the fluxconfig file."
	setClusterCmdLong  = `Specifying a name that already exists will merge new fields on top of existing
values for those fields.`
	setClusterCmdExample = `  # Specify a config file
  fluxctl config set-cluster --fluxconfig ~/.flux/config

  # setCluster the default config file
  fluxctl config set-cluster`
)

// NewCmdConfigsetCluster returns a new cobra command for `fluxctl config setCluster`
// fluxctl config setCluster exports a specified fluxconfig file.
func NewCmdConfigSetCluster() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-cluster",
		Short: setClusterCmdShort,
		Long:  setClusterCmdLong,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("set-cluster called")
		},
	}
	cmd.PersistentFlags().String("fluxconfig", "~/.flux/config", "The path to the fluxconfig file. Defaults to '~/.flux/config'.")
	return cmd
}
