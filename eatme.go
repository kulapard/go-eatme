package main

import (
	"github.com/kulapard/eatme/Godeps/_workspace/src/github.com/spf13/cobra"
	"github.com/kulapard/eatme/hg"
)

func main() {
	var branch string
	var newBranch bool
	var clean bool

	var EatMeCmd = &cobra.Command{
		Use:   "eatme",
		Short: "pull + update",
		Run: func(cmd *cobra.Command, args []string) {
			pullCmd := &hg.HGCommand{Cmd: "pull"}
			pullCmd.RunForAll()

			updateCmd := &hg.HGCommand{Cmd: "update"}
			updateCmd.SetBranch(branch)
			updateCmd.SetClean(clean)
			updateCmd.RunForAll()
		},
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "only update",
		Run: func(cmd *cobra.Command, args []string) {
			updateCmd := &hg.HGCommand{Cmd: "update"}
			updateCmd.SetBranch(branch)
			updateCmd.SetClean(clean)
			updateCmd.RunForAll()
		},
	}
	var cmdPull = &cobra.Command{
		Use:   "pull",
		Short: "only pull",
		Run: func(cmd *cobra.Command, args []string) {
			pullCmd := &hg.HGCommand{Cmd: "pull"}
			pullCmd.RunForAll()
		},
	}
	var cmdPush = &cobra.Command{
		Use:   "push",
		Short: "only push",
		Run: func(cmd *cobra.Command, args []string) {
			pushCmd := &hg.HGCommand{Cmd: "push"}
			pushCmd.SetNewBranch(newBranch)
			pushCmd.RunForAll()
		},
	}

	EatMeCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")
	cmdPush.Flags().BoolVarP(&newBranch, "new-branch", "n", false, "Create remote new branch")
	cmdPush.Flags().BoolVarP(&clean, "clean", "C", false, "Clean update")

	EatMeCmd.AddCommand(cmdUpdate)
	EatMeCmd.AddCommand(cmdPull)
	EatMeCmd.AddCommand(cmdPush)
	EatMeCmd.Execute()

}
