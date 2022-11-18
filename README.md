# GoEatMe

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/kulapard/go-eatme/blob/master/LICENSE)
[![Build Status](https://github.com/kulapard/go-eatme/actions/workflows/release.yml/badge.svg)](https://github.com/kulapard/go-eatme/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kulapard/go-eatme)](https://goreportcard.com/report/github.com/kulapard/go-eatme)

Simple tool to manage multiple git/hg repositories at once.

## Install ##
Using [Homebrew](http://brew.sh/) (OS X / Linux)
```shell
brew install kulapard/tap/eatme
```

## Update ##
Using [Homebrew](http://brew.sh/) (OS X / Linux)
```shell
brew upgrade kulapard/tap/eatme
```

## Usage ##
```shell
eatme [options] [command]
```
By default it runs `fetch` + `pull` + `update` commands at once.
To specify branch use `-b`/`--branch` option:
```shell
eatme -b foo/bar
```

| Command  |           Action           |
|----------|----------------------------|
| `branch` | Show current branch        |
| `fetch`  | Run git fetch              |
| `pull`   | Run git/hg pull            |
| `push`   | Run git/hg push            |
| `update` | Run git checkout/hg update |
| `help`   | Help about any command     |
