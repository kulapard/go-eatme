package cli

import (
	"github.com/kulapard/go-eatme/internal/runner"
	"github.com/spf13/cobra"
)

var cmdBranch = &cobra.Command{
	Use:   "branch",
	Short: "show current branch",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "branch"}
		c.RunRecursively()
	},
}

func init() {
	rootCmd.AddCommand(cmdBranch)
}
