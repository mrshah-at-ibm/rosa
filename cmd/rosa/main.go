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

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mrshah-at-ibm/rosa/cmd/completion"
	"github.com/mrshah-at-ibm/rosa/cmd/create"
	"github.com/mrshah-at-ibm/rosa/cmd/describe"
	"github.com/mrshah-at-ibm/rosa/cmd/dlt"
	"github.com/mrshah-at-ibm/rosa/cmd/docs"
	"github.com/mrshah-at-ibm/rosa/cmd/download"
	"github.com/mrshah-at-ibm/rosa/cmd/edit"
	"github.com/mrshah-at-ibm/rosa/cmd/grant"
	"github.com/mrshah-at-ibm/rosa/cmd/hibernate"
	"github.com/mrshah-at-ibm/rosa/cmd/initialize"
	"github.com/mrshah-at-ibm/rosa/cmd/install"
	"github.com/mrshah-at-ibm/rosa/cmd/link"
	"github.com/mrshah-at-ibm/rosa/cmd/list"
	"github.com/mrshah-at-ibm/rosa/cmd/login"
	"github.com/mrshah-at-ibm/rosa/cmd/logout"
	"github.com/mrshah-at-ibm/rosa/cmd/logs"
	"github.com/mrshah-at-ibm/rosa/cmd/resume"
	"github.com/mrshah-at-ibm/rosa/cmd/revoke"
	"github.com/mrshah-at-ibm/rosa/cmd/uninstall"
	"github.com/mrshah-at-ibm/rosa/cmd/unlink"
	"github.com/mrshah-at-ibm/rosa/cmd/upgrade"
	"github.com/mrshah-at-ibm/rosa/cmd/verify"
	"github.com/mrshah-at-ibm/rosa/cmd/version"
	"github.com/mrshah-at-ibm/rosa/cmd/whoami"
	"github.com/mrshah-at-ibm/rosa/pkg/arguments"
)

var root = &cobra.Command{
	Use:   "rosa",
	Short: "Command line tool for ROSA.",
	Long:  "Command line tool for Red Hat OpenShift Service on AWS.",
}

func init() {
	// Add the command line flags:
	fs := root.PersistentFlags()
	arguments.AddDebugFlag(fs)

	// Register the subcommands:
	root.AddCommand(completion.Cmd)
	root.AddCommand(create.Cmd)
	root.AddCommand(describe.Cmd)
	root.AddCommand(dlt.Cmd)
	root.AddCommand(docs.Cmd)
	root.AddCommand(download.Cmd)
	root.AddCommand(edit.Cmd)
	root.AddCommand(grant.Cmd)
	root.AddCommand(list.Cmd)
	root.AddCommand(initialize.Cmd)
	root.AddCommand(install.Cmd)
	root.AddCommand(login.Cmd)
	root.AddCommand(logout.Cmd)
	root.AddCommand(logs.Cmd)
	root.AddCommand(revoke.Cmd)
	root.AddCommand(uninstall.Cmd)
	root.AddCommand(upgrade.Cmd)
	root.AddCommand(verify.Cmd)
	root.AddCommand(version.Cmd)
	root.AddCommand(whoami.Cmd)
	root.AddCommand(hibernate.Cmd)
	root.AddCommand(resume.Cmd)
	root.AddCommand(link.Cmd)
	root.AddCommand(unlink.Cmd)
}

func main() {
	// Execute the root command:
	root.SetArgs(os.Args[1:])
	err := root.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute root command: %s\n", err)
		os.Exit(1)
	}
}
