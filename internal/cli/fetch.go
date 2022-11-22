package cli

import (
	"github.com/kulapard/go-eatme/internal/runner"
	"github.com/spf13/cobra"
)

var cmdFetch = &cobra.Command{
	Use:   "fetch",
	Short: "run git fetch",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "fetch"}
		c.RunRecursively()
	},
}

func init() {
	rootCmd.AddCommand(cmdFetch)
}
