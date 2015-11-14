package main

import (
	"path/filepath"
	"os"
	"flag"
	"fmt"
	"os/exec"
	"time"
	"github.com/fatih/color"
	"bytes"
	"sync"
	"strings"
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

func hgPullUpdate(path string, branch string, wg *sync.WaitGroup) {
	runHgCommand("pull", path)
	if branch != "" {
		runHgCommand("update", path, "--rev", branch)
	} else {
		runHgCommand("update", path)
	}

	wg.Done()
}


func main() {
	t := time.Now()
	wg := new(sync.WaitGroup)

	flag.Parse()
	branch := flag.Arg(0)
	c := make(chan string)

	go findRepo(".", ".hg", c)

	count := 0
	for path := range c {
		wg.Add(1)
		go hgPullUpdate(path, branch, wg)
		count += 1
	}
	wg.Wait()
	color.Cyan("Done %d repos in %s", count, time.Since(t))
}