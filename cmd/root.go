/*
Copyright © 2023 EDA Egress Team edadataegress@redhat.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/cmwylie19/sbctl/pkg/utils"
	"github.com/spf13/cobra"
)

func GetRootCommand(logger utils.Logger) *cobra.Command {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(getCliCommand(logger))
	return rootCmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sbctl",
	Short: "Starburst configuration tool",
	Long: `A tool to simplify and automate installations and configurations of
Starburst clients, libraries, and connections for data customers.

The tool is designed to unify the configuration into a single source of truth
in order to simplify the process of configuring, debugging, and generally managing
the Starburst client, libraries, and connections.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
