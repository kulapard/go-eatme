package hg

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type HGCommand struct {
	Cmd  string
	args []string
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
			// ignore hidden directories
			matched, _ := regexp.MatchString("/\\.", abs_dir)
			if matched {
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
	args := append([]string{cmd.Cmd, "--repository", path}, cmd.args...)
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
	color.Cyan("Done \"hg %s\" for %d repos in %s\n\n", cmd.Cmd, count, time.Since(t))
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
