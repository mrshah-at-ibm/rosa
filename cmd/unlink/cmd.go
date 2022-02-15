/*
Copyright (c) 2022 Red Hat, Inc.
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

package unlink

import (
	"github.com/spf13/cobra"

	"github.com/mrshah-at-ibm/rosa/cmd/unlink/ocmrole"
	"github.com/mrshah-at-ibm/rosa/cmd/unlink/userrole"
	"github.com/mrshah-at-ibm/rosa/pkg/arguments"
	"github.com/mrshah-at-ibm/rosa/pkg/interactive/confirm"
)

var Cmd = &cobra.Command{
	Use:     "unlink",
	Aliases: []string{"unlink"},
	Short:   "Unlink a resource",
	Long:    "Unlink a resource",
	Hidden:  true,
}

func init() {
	Cmd.AddCommand(ocmrole.Cmd)
	Cmd.AddCommand(userrole.Cmd)

	flags := Cmd.PersistentFlags()
	arguments.AddProfileFlag(flags)
	confirm.AddFlag(flags)
}
