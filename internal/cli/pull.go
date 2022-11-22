package cli

import (
	"github.com/kulapard/go-eatme/internal/runner"
	"github.com/spf13/cobra"
)

var cmdPull = &cobra.Command{
	Use:   "pull",
	Short: "run git/hg pull",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "pull"}
		c.RunRecursively()
	},
}

func init() {
	rootCmd.AddCommand(cmdPull)
}
