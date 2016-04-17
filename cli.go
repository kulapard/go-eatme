package main

import "github.com/spf13/cobra"

func main() {
	var branch string

	var EatMeCmd = &cobra.Command{
		Use:   "eatme",
		Short: "pull + update",
		Run: func(cmd *cobra.Command, args []string) {
			pullUpdateCmd := CliCommand{Name: "pull + update", Branch: branch}
			RunRecursively(pullUpdateCmd)
		},
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "only update",
		Run: func(cmd *cobra.Command, args []string) {
			updateCmd := CliCommand{Name: "update", Branch: branch}
			RunRecursively(updateCmd)
		},
	}
	var cmdPull = &cobra.Command{
		Use:   "pull",
		Short: "only pull",
		Run: func(cmd *cobra.Command, args []string) {
			pullCmd := CliCommand{Name: "pull"}
			RunRecursively(pullCmd)
		},
	}
	var cmdPush = &cobra.Command{
		Use:   "push",
		Short: "only push",
		Run: func(cmd *cobra.Command, args []string) {
			pushCmd := CliCommand{Name: "push"}
			RunRecursively(pushCmd)
		},
	}
	var cmdBranch = &cobra.Command{
		Use:   "branch",
		Short: "show current branch",
		Run: func(cmd *cobra.Command, args []string) {
			pushCmd := CliCommand{Name: "branch"}
			RunRecursively(pushCmd)
		},
	}

	EatMeCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")

	EatMeCmd.AddCommand(cmdUpdate)
	EatMeCmd.AddCommand(cmdPull)
	EatMeCmd.AddCommand(cmdPush)
	EatMeCmd.AddCommand(cmdBranch)
	EatMeCmd.Execute()
}
