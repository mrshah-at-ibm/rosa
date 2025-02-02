/*
Copyright (c) 2021 Red Hat, Inc.

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

package ocm

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mrshah-at-ibm/rosa/pkg/aws"
	"github.com/mrshah-at-ibm/rosa/pkg/logging"
	rprtr "github.com/mrshah-at-ibm/rosa/pkg/reporter"
)

var clusterKey string

func AddClusterFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(
		&clusterKey,
		"cluster",
		"c",
		"",
		"Name or ID of the cluster.",
	)
	cmd.MarkFlagRequired("cluster")
	cmd.RegisterFlagCompletionFunc("cluster", clusterCompletion)
}

func SetClusterKey(key string) {
	clusterKey = key
}

func GetClusterKey() (string, error) {
	// Check that the cluster key (name, identifier or external identifier) given by the user
	// is reasonably safe so that there is no risk of SQL injection:
	if !IsValidClusterKey(clusterKey) {
		return "", fmt.Errorf(
			"Cluster name, identifier or external identifier '%s' isn't valid: it "+
				"must contain only letters, digits, dashes and underscores",
			clusterKey,
		)
	}
	return clusterKey, nil
}

func clusterCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	reporter := rprtr.CreateReporterOrExit()
	logger := logging.CreateLoggerOrExit(reporter)

	ocmClient, err := NewClient().Logger(logger).Build()
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveDefault
	}
	defer ocmClient.Close()

	fmt.Println("Calling aws.NewClient in clusterCompletion function")
	awsClient, err := aws.NewClient().Logger(logger).Build()
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveDefault
	}
	awsCreator, err := awsClient.GetCreator()
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveDefault
	}

	clusters, err := ocmClient.GetClusters(awsCreator, 10)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveDefault
	}
	res := []string{}
	for _, cluster := range clusters {
		res = append(res, cluster.Name())
	}
	return res, cobra.ShellCompDirectiveDefault
}
