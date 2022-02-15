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

package dlt

import (
	"github.com/spf13/cobra"

	"github.com/mrshah-at-ibm/rosa/cmd/dlt/accountroles"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt/admin"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt/cluster"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt/idp"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt/ingress"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt/machinepool"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt/ocmrole"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt/oidcprovider"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt/operatorrole"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt/upgrade"
	"github.com/mrshah-at-ibm/rosa/pkg/arguments"
	"github.com/mrshah-at-ibm/rosa/pkg/interactive/confirm"
)

var Cmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove"},
	Short:   "Delete a specific resource",
	Long:    "Delete a specific resource",
}

func init() {
	Cmd.AddCommand(admin.Cmd)
	Cmd.AddCommand(cluster.Cmd)
	Cmd.AddCommand(idp.Cmd)
	Cmd.AddCommand(ingress.Cmd)
	Cmd.AddCommand(machinepool.Cmd)
	Cmd.AddCommand(upgrade.Cmd)
	Cmd.AddCommand(oidcprovider.Cmd)
	Cmd.AddCommand(operatorrole.Cmd)
	Cmd.AddCommand(accountroles.Cmd)
	Cmd.AddCommand(ocmrole.Cmd)

	flags := Cmd.PersistentFlags()
	arguments.AddProfileFlag(flags)
	arguments.AddRegionFlag(flags)
	confirm.AddFlag(flags)
}
