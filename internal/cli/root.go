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
		runner.RunRecursively(c)
	},
}
var cmdUpdate = &cobra.Command{
	Use:   "update",
	Short: "run git checkout/hg update",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "update", Branch: branch}
		runner.RunRecursively(c)
	},
}
var cmdPull = &cobra.Command{
	Use:   "pull",
	Short: "run git/hg pull",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "pull"}
		runner.RunRecursively(c)
	},
}
var cmdPush = &cobra.Command{
	Use:   "push",
	Short: "run git/hg push",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "push", Branch: branch, All: all}
		runner.RunRecursively(c)
	},
}
var cmdBranch = &cobra.Command{
	Use:   "branch",
	Short: "show current branch",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "branch"}
		runner.RunRecursively(c)
	},
}
var cmdFetch = &cobra.Command{
	Use:   "fetch",
	Short: "run git fetch",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "fetch"}
		runner.RunRecursively(c)
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")
	cmdPush.PersistentFlags().BoolVarP(&all, "all", "a", false, "Push all branches")

	rootCmd.AddCommand(cmdUpdate)
	rootCmd.AddCommand(cmdPull)
	rootCmd.AddCommand(cmdPush)
	rootCmd.AddCommand(cmdBranch)
	rootCmd.AddCommand(cmdFetch)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing CLI '%s'", err)
		os.Exit(1)
	}
}
