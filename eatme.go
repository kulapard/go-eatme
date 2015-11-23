package main

import (
	"github.com/spf13/cobra"
	"github.com/kulapard/eatme/hg"
)


func main() {
	var branch string
	var new_branch bool
	var clean bool

	var EatMeCmd = &cobra.Command{
		Use: "eatme",
		Short: "pull + update",
		Run: func(cmd *cobra.Command, args []string) {
			pull_cmd := &hg.HGCommand{Cmd: "pull"}
			pull_cmd.RunForAll()

			update_cmd := &hg.HGCommand{Cmd: "update"}
			update_cmd.SetBranch(branch)
			update_cmd.SetClean(clean)
			update_cmd.RunForAll()
		},
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "only update",
		Run: func(cmd *cobra.Command, args []string) {
			update_cmd := &hg.HGCommand{Cmd: "update"}
			update_cmd.SetBranch(branch)
			update_cmd.SetClean(clean)
			update_cmd.RunForAll()
		},
	}
	var cmdPull = &cobra.Command{
		Use:   "pull",
		Short: "only pull",
		Run: func(cmd *cobra.Command, args []string) {
			pull_cmd := &hg.HGCommand{Cmd: "pull"}
			pull_cmd.RunForAll()
		},
	}
	var cmdPush = &cobra.Command{
		Use:   "push",
		Short: "only push",
		Run: func(cmd *cobra.Command, args []string) {
			push_cmd := &hg.HGCommand{Cmd: "push"}
			push_cmd.SetNewBranch(new_branch)
			push_cmd.RunForAll()
		},
	}

	EatMeCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")
	cmdPush.Flags().BoolVarP(&new_branch, "new-branch", "n", false, "Create remote new branch")
	cmdPush.Flags().BoolVarP(&clean, "clean", "C", false, "Clean update")

	EatMeCmd.AddCommand(cmdUpdate)
	EatMeCmd.AddCommand(cmdPull)
	EatMeCmd.AddCommand(cmdPush)
	EatMeCmd.Execute()

}