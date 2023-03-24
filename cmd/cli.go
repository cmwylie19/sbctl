/*
Copyright © 2023 EDA Data Egress edadataegress@redhat.com
*/
package cmd

import (
	"github.com/cmwylie19/sbctl/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	mode            string
	operatingSystem string
)

// cliCmd represents the cli command
func getCliCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "Install the Trino CLI",
		Long: `Install the Trino CLI which provides a terminal-based, 
	interactive shell for running queries.
	
	The CLI is a self-executing JAR file, which means it acts
	like a normal UNIX executable. The CLI uses the Trino client
	REST API over HTTP/HTTPS to communicate with the coordinator
	on the cluster.`,
		Run: func(cmd *cobra.Command, args []string) {
			client := cli.NewCli(operatingSystem)
			if mode == "install" {
				client.Install()
			} else if mode == "uninstall" {
				client.Uninstall()
			} else {
				panic("Invalid mode")
			}
		},
	}
	cmd.PersistentFlags().StringVarP(&mode, "mode", "", "install", "Options: install or uninstall")

	cmd.PersistentFlags().StringVarP(&operatingSystem, "os", "", "mac", "Target Operating System to install the Trino CLI on. Options: mac, linux, or windows.")
	return cmd
}
