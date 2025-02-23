package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/kulapard/go-eatme/internal/runner"
)

type config struct {
	branch string
	all    bool
}

var cfg = &config{}

var rootCmd = &cobra.Command{
	Use:   "eatme",
	Short: "pull + update",
	Run: func(cmd *cobra.Command, args []string) {
		c := runner.CliCommand{Name: "pull + update", Branch: cfg.branch}
		c.RunRecursively()
	},
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func Execute(version string) {
	rootCmd.PersistentFlags().StringVarP(&cfg.branch, "branch", "b", "", "Branch or Tag name")
	cmdPush.PersistentFlags().BoolVarP(&cfg.all, "all", "a", false, "Push all branches")
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}
