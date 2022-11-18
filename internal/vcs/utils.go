package vcs

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func stringCommand(cmd *exec.Cmd) string {
	return color.YellowString("%s\n", strings.Join(cmd.Args, " "))
}

func stringError(err error) string {
	if err != nil {
		return color.RedString(fmt.Sprintf("%s\n", err.Error()))
	}
	return ""
}

func stringOutput(outs []byte) string {
	if len(outs) > 0 {
		return fmt.Sprintf("%s\n", string(outs))
	}
	return ""
}

func execCommand(path string, c *exec.Cmd) {
	output, err := c.CombinedOutput()
	fmt.Printf("%s\n%s%s%s", color.GreenString(path), stringCommand(c), stringError(err), stringOutput(output))
}
