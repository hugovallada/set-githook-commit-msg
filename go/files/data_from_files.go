package files

import (
	"fmt"
	"os"
	"strings"

	"github.com/hugovallada/set-branch/branch"
	"github.com/hugovallada/set-branch/checker"
)

func GetCommitTitle(currentDir string) string {
	data, err := os.ReadFile(fmt.Sprintf("%s/.git/HEAD", currentDir))
	checker.Check(err)
	branchName := getBranchName(string(data))
	branchType, branchID := getTypeAndId(branchName)
	currentBranch := branch.NewBranch(branchType, branchID)
	return currentBranch.GetCommitTitle()
}

func getBranchName(branch string) string {
	return strings.SplitAfter(branch, "heads/")[1]
}

func getTypeAndId(branchName string) (string, string) {
	splitedBranchName := strings.Split(branchName, "-")
	name, taskId := splitedBranchName[0], splitedBranchName[1]
	return strings.Split(name, "/")[0], taskId
}

func WriteToFile(path string, data string) {
	file, err := os.Create(fmt.Sprintf("%s/.git/hooks/commit-msg", path))
	checker.Check(err)
	file.WriteString(data)
	os.Chmod(file.Name(), 0775)
}
