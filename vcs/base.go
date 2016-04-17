package vcs

type VcsCommand interface {
	Execute(path string)
}
