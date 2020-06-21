# git-fixup

```
$ go get -u akhil.cc/git-fixup

usage:
	git-fixup [--root] [HEAD~N]
	git-fixup edit filename
```
Squash the last N commits of the current branch into a single commit. This
tool automatically performs an interactive rebase for you with either
`HEAD~N`, to specify the last N commits, or `--root` to go to the initial
commit. For instance, if `git fixup HEAD~3` is specified, it automatically
applies the following:
```
reword AAAAAAA Message
fixup  BBBBBBB Message
fixup  CCCCCCC Message
```
This gives you a chance to reword the oldest of the N commits.

`git-fixup edit` is the entrypoint for the `GIT_SEQUENCE_EDITOR` that applies
edits to the interactive rebase on your behalf.