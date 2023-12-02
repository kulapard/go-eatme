package runner

import (
	"github.com/kulapard/go-eatme/internal/vcs"
	"reflect"
	"testing"
)

func TestCliCommand_GetVcsCommand_Git(t *testing.T) {
	cmdMap := map[string]interface{}{
		"pull":          &vcs.GitPull{},
		"update":        &vcs.GitUpdate{},
		"push":          &vcs.GitPush{},
		"pull + update": &vcs.GitPullUpdate{},
		"branch":        &vcs.GitBranch{},
		"fetch":         &vcs.GitFetch{},
	}

	for cmdStr, vcsCmdExp := range cmdMap {
		cli := CliCommand{Name: cmdStr}
		vcsCmd := cli.GetVcsCommand(".git")
		typeOfCmdExp := reflect.TypeOf(vcsCmdExp)
		typeOfCmd := reflect.TypeOf(vcsCmd)

		if typeOfCmd != typeOfCmdExp {
			t.Errorf("Git `%s` should be %s, got %s", cmdStr, typeOfCmdExp, typeOfCmd)
		}
	}
}

func TestCliCommand_GetVcsCommand_Hg(t *testing.T) {
	cmdMap := map[string]interface{}{
		"pull":          &vcs.HgPull{},
		"update":        &vcs.HgUpdate{},
		"push":          &vcs.HgPush{},
		"pull + update": &vcs.HgPullUpdate{},
		"branch":        &vcs.HgBranch{},
	}

	for cmdStr, vcsCmdExp := range cmdMap {
		cli := CliCommand{Name: cmdStr}
		vcsCmd := cli.GetVcsCommand(".hg")
		typeOfCmdExp := reflect.TypeOf(vcsCmdExp)
		typeOfCmd := reflect.TypeOf(vcsCmd)

		if typeOfCmd != typeOfCmdExp {
			t.Errorf("Mercurial `%s` should be %s, got %s", cmdStr, typeOfCmdExp, typeOfCmd)
		}
	}
}
