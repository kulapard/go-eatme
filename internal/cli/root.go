package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/kulapard/go-eatme/internal/runner"
)

var branch string
var all bool

var rootCmd = &cobra.Command{
	Use:   "eatme",
	Short: "pull + update",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "pull + update", Branch: branch}
		c.RunRecursively()
	},
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func Execute(version string) {
	rootCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")
	cmdPush.PersistentFlags().BoolVarP(&all, "all", "a", false, "Push all branches")
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing CLI '%s'", err)
		os.Exit(1)
	}
}
