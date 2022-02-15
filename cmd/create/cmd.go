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

package create

import (
	"github.com/spf13/cobra"

	"github.com/mrshah-at-ibm/rosa/cmd/create/accountroles"
	"github.com/mrshah-at-ibm/rosa/cmd/create/admin"
	"github.com/mrshah-at-ibm/rosa/cmd/create/cluster"
	"github.com/mrshah-at-ibm/rosa/cmd/create/idp"
	"github.com/mrshah-at-ibm/rosa/cmd/create/ingress"
	"github.com/mrshah-at-ibm/rosa/cmd/create/machinepool"
	"github.com/mrshah-at-ibm/rosa/cmd/create/ocmrole"
	"github.com/mrshah-at-ibm/rosa/cmd/create/oidcprovider"
	"github.com/mrshah-at-ibm/rosa/cmd/create/operatorroles"
	"github.com/mrshah-at-ibm/rosa/cmd/create/userrole"

	"github.com/mrshah-at-ibm/rosa/pkg/arguments"
	"github.com/mrshah-at-ibm/rosa/pkg/interactive/confirm"
)

var Cmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Create a resource from stdin",
	Long:    "Create a resource from stdin",
}

func init() {
	Cmd.AddCommand(accountroles.Cmd)
	Cmd.AddCommand(admin.Cmd)
	Cmd.AddCommand(cluster.Cmd)
	Cmd.AddCommand(idp.Cmd)
	Cmd.AddCommand(ingress.Cmd)
	Cmd.AddCommand(machinepool.Cmd)
	Cmd.AddCommand(oidcprovider.Cmd)
	Cmd.AddCommand(operatorroles.Cmd)
	Cmd.AddCommand(userrole.Cmd)
	Cmd.AddCommand(ocmrole.Cmd)

	flags := Cmd.PersistentFlags()
	arguments.AddProfileFlag(flags)
	arguments.AddRegionFlag(flags)
	confirm.AddFlag(flags)
}
