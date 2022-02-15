/*
Copyright (c) 2020 Red Hat, Inc.

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

package edit

import (
	"github.com/spf13/cobra"

	"github.com/mrshah-at-ibm/rosa/cmd/edit/addon"
	"github.com/mrshah-at-ibm/rosa/cmd/edit/cluster"
	"github.com/mrshah-at-ibm/rosa/cmd/edit/ingress"
	"github.com/mrshah-at-ibm/rosa/cmd/edit/machinepool"
	"github.com/mrshah-at-ibm/rosa/pkg/arguments"
	"github.com/mrshah-at-ibm/rosa/pkg/interactive"
)

var Cmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"update"},
	Short:   "Edit a specific resource",
	Long:    "Edit a specific resource",
}

func init() {
	Cmd.AddCommand(addon.Cmd)
	Cmd.AddCommand(cluster.Cmd)
	Cmd.AddCommand(ingress.Cmd)
	Cmd.AddCommand(machinepool.Cmd)

	flags := Cmd.PersistentFlags()
	arguments.AddProfileFlag(flags)
	arguments.AddRegionFlag(flags)
	interactive.AddFlag(flags)
}
