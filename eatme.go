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

func hgStatus(path string, wg *sync.WaitGroup) {
	runHgCommand("status", path)
	wg.Done()
}

func hgPull(path string, wg *sync.WaitGroup) {
	runHgCommand("pull", path)
	wg.Done()
}

func hgUpdate(path string, wg *sync.WaitGroup) {
	runHgCommand("update", path)
	wg.Done()
}

func hgPullUpdate(path string, wg *sync.WaitGroup, branch string) {
	runHgCommand("pull", path)
	if branch != "" {
		runHgCommand("update", path, "--rev", branch)
	} else {
		runHgCommand("update", path)
	}

	wg.Done()
}

func runCommand(){
	t := time.Now()
	wg := new(sync.WaitGroup)
	c := make(chan string)
	count := 0

	go findRepo(".", ".hg", c)

	for path := range c {
		wg.Add(1)
		go hgPullUpdate(path, wg, "")
		count += 1
	}

	wg.Wait()
	color.Cyan("Done %d repos in %s", count, time.Since(t))
}


func main() {
	var branch string

	var EatMeCmd = &cobra.Command{
		Use: "go-eatme",
		Short: "pull + update",
		Run: func(cmd *cobra.Command, args []string) {
			runCommand()
		},
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "only update",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			runCommand()
		},
	}

	EatMeCmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")
	cmdUpdate.Flags().StringVarP(&branch, "branch", "b", "", "Branch or Tag name")

	EatMeCmd.AddCommand(cmdUpdate)
	EatMeCmd.Execute()

}