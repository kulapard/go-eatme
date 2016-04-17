package main

import (
	"os"
	"path/filepath"
	"regexp"

	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/kulapard/go-eatme/vcs"
)

type CliCommand struct {
	Name   string
	Branch string
}

type vcsPath struct {
	Path string
	Sign string
}

func (cmd CliCommand) GetVcsCommand(sign string) vcs.VcsCommand {
	switch {
	case sign == ".hg":
		switch cmd.Name {
		case "pull":
			return vcs.HgPull{}
		case "update":
			return vcs.HgUpdate{Branch: cmd.Branch}
		case "pull + update":
			return vcs.HgPullUpdate{Branch: cmd.Branch}
		case "branch":
			return vcs.HgBranch{}

		}
	case sign == ".git":
		switch cmd.Name {
		case "pull":
			return vcs.GitPull{}
		case "update":
			return vcs.GitUpdate{Branch: cmd.Branch}
		case "pull + update":
			return vcs.GitPullUpdate{Branch: cmd.Branch}
		case "branch":
			return vcs.GitBranch{}
		}
	}
	return nil
}

func RunRecursively(cmd CliCommand) {
	t := time.Now()
	wg := new(sync.WaitGroup)
	pathChan := make(chan vcsPath)
	count := 0

	go findRepositories(".", pathChan)

	for p := range pathChan {
		vcsCmd := cmd.GetVcsCommand(p.Sign)
		if vcsCmd != nil {
			wg.Add(1)
			go func(path string, wg *sync.WaitGroup) {
				vcsCmd.Execute(p.Path)
				wg.Done()
			}(p.Path, wg)
			count++
		}
	}

	wg.Wait()
	color.Cyan("Done \"%s\" command for %d repos in %s\n\n", cmd.Name, count, time.Since(t))
}

func findRepositories(root string, pathChan chan vcsPath) {
	defer close(pathChan)
	visit := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}

		if info.IsDir() {
			switch info.Name() {
			case ".hg", ".git":
				sign := info.Name()
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

				pathChan <- vcsPath{Path: absDir, Sign: sign}

			}
		}

		return nil
	}

	filepath.Walk(root, visit)
}
