package vcs

type Command interface {
	Execute(path string)
}
