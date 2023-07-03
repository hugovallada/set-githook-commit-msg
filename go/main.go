package main

import (
	"os"

	"github.com/hugovallada/set-branch/checker"
	"github.com/hugovallada/set-branch/files"
)

func setCommitTitleByBranch(currentDir string) {
	commitTitle := files.GetCommitTitle(currentDir)
	generateFile(currentDir, commitTitle)
}

func setCommitTitleByMessage(currentDir, message string) {
	generateFile(currentDir, message)
}

func generateFile(dir, message string) {
	commitHook := files.ReplaceDefaultValue(message)
	files.WriteToFile(dir, commitHook)
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
