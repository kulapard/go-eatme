package vcs

import "os/exec"

type HgUpdate struct {
	Branch string
}
type HgPull struct {
}
type HgPullUpdate struct {
	Branch string
}
type HgBranch struct {
}

func (cmd HgUpdate) Execute(path string) {
	args := []string{"update", "--repository", path}
	if cmd.Branch != "" {
		args = append(args, "--rev", cmd.Branch)
	}

	systemCmd := exec.Command("hg", args...)
	execCommand(path, systemCmd)
}

func (cmd HgPull) Execute(path string) {
	args := []string{"pull", "--repository", path}
	systemCmd := exec.Command("hg", args...)
	execCommand(path, systemCmd)
}

func (cmd HgPullUpdate) Execute(path string) {
	HgPull{}.Execute(path)
	HgUpdate{Branch: cmd.Branch}.Execute(path)
}

func (cmd HgBranch) Execute(path string) {
	args := []string{"branch", "--repository", path}
	systemCmd := exec.Command("hg", args...)
	execCommand(path, systemCmd)
}
