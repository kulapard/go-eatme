package cli

import (
	"github.com/kulapard/go-eatme/internal/runner"
	"github.com/spf13/cobra"
)

var cmdPush = &cobra.Command{
	Use:   "push",
	Short: "run git/hg push",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "push", Branch: branch, All: all}
		c.RunRecursively()
	},
}

func init() {
	rootCmd.AddCommand(cmdPush)
}
