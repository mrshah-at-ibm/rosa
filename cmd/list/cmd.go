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

package list

import (
	"github.com/spf13/cobra"

	"github.com/mrshah-at-ibm/rosa/cmd/list/accountroles"
	"github.com/mrshah-at-ibm/rosa/cmd/list/addon"
	"github.com/mrshah-at-ibm/rosa/cmd/list/cluster"
	"github.com/mrshah-at-ibm/rosa/cmd/list/gates"
	"github.com/mrshah-at-ibm/rosa/cmd/list/idp"
	"github.com/mrshah-at-ibm/rosa/cmd/list/ingress"
	"github.com/mrshah-at-ibm/rosa/cmd/list/instancetypes"
	"github.com/mrshah-at-ibm/rosa/cmd/list/machinepool"
	"github.com/mrshah-at-ibm/rosa/cmd/list/ocmroles"
	"github.com/mrshah-at-ibm/rosa/cmd/list/region"
	"github.com/mrshah-at-ibm/rosa/cmd/list/upgrade"
	"github.com/mrshah-at-ibm/rosa/cmd/list/user"
	"github.com/mrshah-at-ibm/rosa/cmd/list/userroles"
	"github.com/mrshah-at-ibm/rosa/cmd/list/version"
	"github.com/mrshah-at-ibm/rosa/pkg/arguments"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List all resources of a specific type",
	Long:  "List all resources of a specific type",
}

func init() {
	Cmd.AddCommand(addon.Cmd)
	Cmd.AddCommand(cluster.Cmd)
	Cmd.AddCommand(gates.Cmd)
	Cmd.AddCommand(idp.Cmd)
	Cmd.AddCommand(ingress.Cmd)
	Cmd.AddCommand(machinepool.Cmd)
	Cmd.AddCommand(region.Cmd)
	Cmd.AddCommand(upgrade.Cmd)
	Cmd.AddCommand(user.Cmd)
	Cmd.AddCommand(version.Cmd)
	Cmd.AddCommand(instancetypes.Cmd)
	Cmd.AddCommand(accountroles.Cmd)
	Cmd.AddCommand(ocmroles.Cmd)
	Cmd.AddCommand(userroles.Cmd)
	flags := Cmd.PersistentFlags()
	arguments.AddProfileFlag(flags)
	arguments.AddRegionFlag(flags)
}
