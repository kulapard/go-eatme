package main

import (
	"github.com/kulapard/go-eatme/internal/runner"
	"github.com/spf13/cobra"
)

var cmdPush = &cobra.Command{
	Use:   "push",
	Short: "run git/hg push",
	Long: `Push commits to remote repository.
Supports both git and hg repositories.
Use --all flag to push all branches.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "push", Branch: cfg.branch, All: cfg.all}
		c.RunRecursively()
	},
}

func init() {
	rootCmd.AddCommand(cmdPush)
}
