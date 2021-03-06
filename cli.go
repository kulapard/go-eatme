package main

import "github.com/spf13/cobra"

func main() {
	var branch string
	var all bool

	var EatMeCmd = &cobra.Command{
		Use:   "eatme",
		Short: "pull + update",
		Run: func(cmd *cobra.Command, args []string) {
			c := CliCommand{Name: "pull + update", Branch: branch}
			RunRecursively(c)
		},
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "run git checkout/hg update",
		Run: func(cmd *cobra.Command, args []string) {
			c := CliCommand{Name: "update", Branch: branch}
			RunRecursively(c)
		},
	}
	var cmdPull = &cobra.Command{
		Use:   "pull",
		Short: "run git/hg pull",
		Run: func(cmd *cobra.Command, args []string) {
			c := CliCommand{Name: "pull"}
			RunRecursively(c)
		},
	}
	var cmdPush = &cobra.Command{
		Use:   "push",
		Short: "run git/hg push",
		Run: func(cmd *cobra.Command, args []string) {
			c := CliCommand{Name: "push", Branch: branch, All: all}
			RunRecursively(c)
		},
	}
	var cmdBranch = &cobra.Command{
		Use:   "branch",
		Short: "show current branch",
		Run: func(cmd *cobra.Command, args []string) {
			c := CliCommand{Name: "branch"}
			RunRecursively(c)
		},
	}
	var cmdFetch = &cobra.Command{
		Use:   "fetch",
		Short: "run git fetch",
		Run: func(cmd *cobra.Command, args []string) {
			c := CliCommand{Name: "fetch"}
			RunRecursively(c)
		},
	}

	EatMeCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")
	cmdPush.PersistentFlags().BoolVarP(&all, "all", "a", false, "Push all branches")

	EatMeCmd.AddCommand(cmdUpdate)
	EatMeCmd.AddCommand(cmdPull)
	EatMeCmd.AddCommand(cmdPush)
	EatMeCmd.AddCommand(cmdBranch)
	EatMeCmd.AddCommand(cmdFetch)
	EatMeCmd.Execute()
}
