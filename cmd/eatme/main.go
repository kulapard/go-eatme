package main

import (
	"github.com/kulapard/go-eatme/internal/cli"
)

var (
	version = "unknown-local-build"
)

func main() {
	cli.Execute(version)
}
