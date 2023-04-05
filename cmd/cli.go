/*
Copyright © 2023 EDA Data Egress edadataegress@redhat.com
*/
package cmd

import (
	"github.com/cmwylie19/sbctl/pkg/cli"
	"github.com/cmwylie19/sbctl/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	mode            string
	operatingSystem string
)

// cliCmd represents the cli command
func getCliCommand(logger utils.Log) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "Install/Uninstall the Trino CLI",
		Long: `Install/Uninstall the Trino CLI which provides a terminal-based, 
interactive shell for running queries.

The CLI is a self-executing JAR file, which means it acts
like a normal UNIX executable. The CLI uses the Trino client
REST API over HTTP/HTTPS to communicate with the coordinator
on the cluster. 

More info can be found here: https://starburstdata.github.io/latest/client/cli.html

Examples:
  sbctl cli --os=linux --mode=install

  sbctl cli --os=mac --mode=uninstall

  sbctl cli --os=windows --mode=install
`,
		Run: func(cmd *cobra.Command, args []string) {
			client := cli.NewCli(operatingSystem, logger)
			if mode == "install" {
				client.Install()
			} else if mode == "uninstall" {
				client.Uninstall()
			} else {

				client.Logger.Printf("%s Valid options are install, or uninstall.\n\nEXAMPLE:\nsbctl cli --os=linux --mode=install\n\nsbctl: try 'sbctl cli -h' for more information.\n\n", "invalid mode!")
			}
		},
	}
	cmd.PersistentFlags().StringVarP(&mode, "mode", "", "", "Options: install or uninstall")

	cmd.PersistentFlags().StringVarP(&operatingSystem, "os", "", "", "Target Operating System to install the Trino CLI on. Options: mac, linux, or windows.")
	return cmd
}
