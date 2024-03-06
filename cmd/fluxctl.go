/*
Copyright © 2024 TechDufus

The cmd package is the entry point for the fluxctl command-line tool.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/techdufus/fluxctl/cmd/config"
	"github.com/techdufus/fluxctl/cmd/get"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fluxctl",
	Short: "The RunOnFlux command-line tool, fluxctl, allows you to run commands against the flux api.",
	Long: `This tool allows you to control your flux assets and applications.

  This project is still a work in progress, so please report any issues or feature requests to the project's github page at https://github.com/techdufus/fluxctl`,
	Version: "0.0.1",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		config.NewCmdConfig(),
		get.NewCmdGet(),
	)
}
