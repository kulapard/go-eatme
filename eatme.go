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



func findRepo(root string, sign string, c chan string) {
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
			c <- abs_dir
		}
		return nil
	}

	filepath.Walk(root, visit)
	close(c)
}


func runHgCommand(hg_cmd string, path string, args ...string) {
	args = append([]string{hg_cmd, "--repository", path}, args...)
	cmd := exec.Command("hg", args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	color.Green(path)
	color.Yellow("hg %s", strings.Join(args, " "))

	if err != nil {
		color.Red(err.Error())
	}
	fmt.Println(out.String())
}

func hgStatus(path string, wg *sync.WaitGroup, branch string, new_branch bool) {
	runHgCommand("status", path)
	wg.Done()
}

func hgPull(path string, wg *sync.WaitGroup, branch string, new_branch bool) {
	runHgCommand("pull", path)
	wg.Done()
}

func hgPush(path string, wg *sync.WaitGroup, branch string, new_branch bool) {
	if new_branch{
		runHgCommand("push", path, "--new-branch")
	}else {
		runHgCommand("push", path)
	}
	wg.Done()
}

func hgUpdate(path string, wg *sync.WaitGroup, branch string, new_branch bool) {
	if branch != "" {
		runHgCommand("update", path, "--rev", branch)
	} else {
		runHgCommand("update", path)
	}
	wg.Done()
}

func hgPullUpdate(path string, wg *sync.WaitGroup, branch string, new_branch bool) {
	runHgCommand("pull", path)
	if branch != "" {
		runHgCommand("update", path, "--rev", branch)
	} else {
		runHgCommand("update", path)
	}

	wg.Done()
}

func runCommand(cmdFunc func(path string, wg *sync.WaitGroup, branch string, new_branch bool), branch string, new_branch bool){
	t := time.Now()
	wg := new(sync.WaitGroup)
	c := make(chan string)
	count := 0

	go findRepo(".", ".hg", c)

	for path := range c {
		wg.Add(1)
		go cmdFunc(path, wg, branch, new_branch)
		count += 1
	}

	wg.Wait()
	color.Cyan("Done %d repos in %s", count, time.Since(t))
}


func main() {
	var branch string
	var new_branch bool

	var EatMeCmd = &cobra.Command{
		Use: "eatme",
		Short: "pull + update",
		Run: func(cmd *cobra.Command, args []string) {
			runCommand(hgPullUpdate, branch, new_branch)
		},
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "only update",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			runCommand(hgUpdate, branch, new_branch)
		},
	}
	var cmdPull = &cobra.Command{
		Use:   "pull",
		Short: "only pull",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			runCommand(hgPull, branch, new_branch)
		},
	}
	var cmdPush = &cobra.Command{
		Use:   "push",
		Short: "only push",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			runCommand(hgPush, branch, new_branch)
		},
	}

	EatMeCmd.PersistentFlags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")
	cmdPush.Flags().BoolVarP(&new_branch, "new-branch", "n", false, "Create remote new branch")

	EatMeCmd.AddCommand(cmdUpdate)
	EatMeCmd.AddCommand(cmdPull)
	EatMeCmd.AddCommand(cmdPush)
	EatMeCmd.Execute()

}