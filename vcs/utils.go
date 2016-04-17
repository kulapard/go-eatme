package vcs

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func printCommand(cmd *exec.Cmd) {
	color.Yellow("%s\n", strings.Join(cmd.Args, " "))
	//fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		color.Red(fmt.Sprintf("%s\n", err.Error()))
		//os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("%s\n", string(outs))
		//fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func execCommand(path string, c *exec.Cmd) {
	output, err := c.CombinedOutput()
	color.Green(path)
	printCommand(c)
	printError(err)
	printOutput(output)
}
