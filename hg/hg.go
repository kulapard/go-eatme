package hg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type HGCommand struct {
	Cmd  string
	args []string
}

func findRepo(root string, sign string, pathChan chan string) {
	defer close(pathChan)

	visit := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}

		if info.Name() == sign && info.IsDir() {
			dir, _ := filepath.Split(path)
			absDir, err := filepath.Abs(dir)
			if err != nil {
				color.Red(err.Error())
				return nil
			}
			// ignore hidden directories
			matched, _ := regexp.MatchString("/\\.", absDir)
			if matched {
				return nil
			}

			pathChan <- absDir
		}
		return nil
	}

	filepath.Walk(root, visit)
}

func (cmd *HGCommand) Run(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	args := append([]string{cmd.Cmd, "--repository", path}, cmd.args...)
	systemCmd := exec.Command("hg", args...)

	var out bytes.Buffer
	systemCmd.Stdout = &out

	err := systemCmd.Run()

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
	pathChan := make(chan string)
	count := 0

	go findRepo(".", ".hg", pathChan)

	for path := range pathChan {
		wg.Add(1)
		go cmd.Run(path, wg)
		count += 1
	}

	wg.Wait()
	color.Cyan("Done \"hg %s\" for %d repos in %s\n\n", cmd.Cmd, count, time.Since(t))
}

func (cmd *HGCommand) SetBranch(branch string) {
	if branch != "" {
		cmd.args = append(cmd.args, "--rev", branch)
	}
}
func (cmd *HGCommand) SetNewBranch(newBranch bool) {
	if newBranch {
		cmd.args = append(cmd.args, "--new-branch")
	}
}

func (cmd *HGCommand) SetClean(clean bool) {
	if clean {
		cmd.args = append(cmd.args, "--clean")
	}
}
