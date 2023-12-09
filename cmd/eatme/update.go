package main

import (
	"github.com/kulapard/go-eatme/internal/runner"
	"github.com/spf13/cobra"
)

var cmdUpdate = &cobra.Command{
	Use:   "update",
	Short: "run git checkout/hg update",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "update", Branch: branch}
		c.RunRecursively()
	},
}

func init() {
	rootCmd.AddCommand(cmdUpdate)
}
