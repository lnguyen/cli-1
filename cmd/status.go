package cmd

import (
	"fmt"
	"github.com/onepanelio/cli/util"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Check deployment status.",
	Long:    "Check deployment status by checking pods statuses.",
	Example: "status",
	Run: func(cmd *cobra.Command, args []string) {
		ready, err := util.DeploymentStatus()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if ready {
			fmt.Println("Deployment is reporting as ready.")
		} else {
			fmt.Println("Deployment is not ready.")
		}
		fmt.Println("Note: be sure to check that your pods are all running by running this command: kubectl get pods --all-namespaces.")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
