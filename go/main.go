package main

import (
	"os"

	"github.com/hugovallada/set-branch/checker"
	"github.com/hugovallada/set-branch/files"
)

func setCommitTitleByBranch(currentDir string) {
	commitTitle := files.GetCommitTitle(currentDir)
	commitHook := files.ReplaceDefaultValue(commitTitle)
	files.WriteToFile(currentDir, commitHook)
}

func setCommitTitleByMessage(currentDir, message string) {
	commitHook := files.ReplaceDefaultValue(message)
	files.WriteToFile(currentDir, commitHook)
}

func main() {
	currentDir, err := os.Getwd()
	checker.Check(err)
	if len(os.Args) > 1 {
		setCommitTitleByMessage(currentDir, os.Args[1])
	} else {
		setCommitTitleByBranch(currentDir)
	}
}
