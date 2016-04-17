package vcs

import "os/exec"

type GitUpdate struct {
	Branch string
}
type GitPull struct {
}
type GitPullUpdate struct {
	Branch string
}
type GitBranch struct {
}
type GitFetch struct {
}
type GitPush struct {
	Branch string
}

func (cmd GitUpdate) Execute(path string) {
	args := []string{"-C", path, "checkout"}

	if cmd.Branch != "" {
		args = append(args, cmd.Branch)
	} else {
		args = append(args, "HEAD")
	}

	systemCmd := exec.Command("git", args...)
	execCommand(path, systemCmd)
}

func (cmd GitPull) Execute(path string) {
	args := []string{"-C", path, "pull"}
	systemCmd := exec.Command("git", args...)
	execCommand(path, systemCmd)
}

func (cmd GitPullUpdate) Execute(path string) {
	GitUpdate{Branch: cmd.Branch}.Execute(path)
	GitPull{}.Execute(path)
}

func (cmd GitBranch) Execute(path string) {
	args := []string{"-C", path, "rev-parse", "--abbrev-ref", "HEAD"}
	systemCmd := exec.Command("git", args...)
	execCommand(path, systemCmd)
}

func (cmd GitFetch) Execute(path string) {
	args := []string{"-C", path, "fetch", "--tags", "--prune"}
	systemCmd := exec.Command("git", args...)
	execCommand(path, systemCmd)
}

func (cmd GitPush) Execute(path string) {
	args := []string{"-C", path, "push", "--set-upstream"}

	if cmd.Branch != "" {
		args = append(args, "origin", cmd.Branch)
	} else {
		args = append(args, "--all")
	}

	systemCmd := exec.Command("git", args...)
	execCommand(path, systemCmd)
}
