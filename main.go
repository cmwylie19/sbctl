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
package main

import (
	"os"

	"github.com/cmwylie19/sbctl/cmd"
	"github.com/cmwylie19/sbctl/pkg/utils"
)

var logger utils.Logger

func main() {
	if err := cmd.GetRootCommand(logger).Execute(); err != nil {
		logger.Infof("Error executing command: %v", err)
		os.Exit(1)
	}
}

func init() {
	logger = utils.Logger{Debug: true}
}
