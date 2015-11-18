package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"fmt"
	"time"
	"bytes"
	"sync"
	"strings"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

type HGCommand struct {
	hg_cmd string
	args   []string
}

func findRepo(root string, sign string, path_chan chan string) {
	defer close(path_chan)

	visit := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}

		if info.Name() == sign && info.IsDir() {
			dir, _ := filepath.Split(path)
			abs_dir, err := filepath.Abs(dir)
			if err != nil {
				color.Red(err.Error())
				return nil
			}
			path_chan <- abs_dir
		}
		return nil
	}

	filepath.Walk(root, visit)
}


func (cmd *HGCommand) Run(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	args := append([]string{cmd.hg_cmd, "--repository", path}, cmd.args...)
	system_cmd := exec.Command("hg", args...)

	var out bytes.Buffer
	system_cmd.Stdout = &out

	err := system_cmd.Run()

	color.Green(path)
	color.Yellow("hg %s", strings.Join(args, " "))

	if err != nil {
		color.Red(err.Error())
	}
	fmt.Println(out.String())
}

func (cmd *HGCommand) RunForAll() {
	t := time.Now()
	wg := new(sync.WaitGroup)
	path_chan := make(chan string)
	count := 0

	go findRepo(".", ".hg", path_chan)

	for path := range path_chan {
		wg.Add(1)
		go cmd.Run(path, wg)
		count += 1
	}

	wg.Wait()
	color.Cyan("Done \"hg %s\" for %d repos in %s\n\n", cmd.hg_cmd, count, time.Since(t))
}

func (cmd *HGCommand) SetBranch(branch string) {
	if branch != "" {
		cmd.args = append(cmd.args, "--rev", branch)
	}
}
func (cmd *HGCommand) SetNewBranch(new_branch bool) {
	if new_branch {
		cmd.args = append(cmd.args, "--new-branch")
	}
}

func (cmd *HGCommand) SetClean(clean bool) {
	if clean {
		cmd.args = append(cmd.args, "--clean")
	}
}

func main() {
	var branch string
	var new_branch bool
	var clean bool

	var EatMeCmd = &cobra.Command{
		Use: "eatme",
		Short: "pull + update",
		Run: func(cmd *cobra.Command, args []string) {
			pull_cmd := &HGCommand{hg_cmd: "pull"}
			pull_cmd.RunForAll()

			update_cmd := &HGCommand{hg_cmd: "update"}
			update_cmd.SetClean(clean)
			update_cmd.RunForAll()
		},
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "only update",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			update_cmd := &HGCommand{hg_cmd: "update"}
			update_cmd.SetClean(clean)
			update_cmd.RunForAll()
		},
	}
	var cmdPull = &cobra.Command{
		Use:   "pull",
		Short: "only pull",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			pull_cmd := &HGCommand{hg_cmd: "pull"}
			pull_cmd.RunForAll()
		},
	}
	var cmdPush = &cobra.Command{
		Use:   "push",
		Short: "only push",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			push_cmd := &HGCommand{hg_cmd: "push"}
			push_cmd.SetNewBranch(new_branch)
			push_cmd.RunForAll()
		},
	}

	EatMeCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")
	cmdPush.Flags().BoolVarP(&new_branch, "new-branch", "n", false, "Create remote new branch")
	cmdPush.Flags().BoolVarP(&clean, "clean", "C", false, "Clean update")

	EatMeCmd.AddCommand(cmdUpdate)
	EatMeCmd.AddCommand(cmdPull)
	EatMeCmd.AddCommand(cmdPush)
	EatMeCmd.Execute()

}