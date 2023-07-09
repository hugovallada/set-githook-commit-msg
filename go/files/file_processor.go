package files

import (
	"fmt"
	"os"
	"strings"

	"github.com/hugovallada/set-branch/branch"
	"github.com/hugovallada/set-branch/checker"
)

type fileProcessor struct {
	branchName  string
	commitTitle string
	fileDir     string
}

func NewFileProcessor(currentDir string, commitTitle ...string) *fileProcessor {
	if len(commitTitle) == 0 {
		commitTitle = append(commitTitle, "")
	}

	return &fileProcessor{
		fileDir:     currentDir,
		commitTitle: commitTitle[0],
	}
}

func (fp *fileProcessor) Execute() {
	fp.getBranchName().
		getCommitTitle().
		writeToFile()
}

func (fp *fileProcessor) getBranchName() *fileProcessor {
	if fp.commitTitle != "" {
		return fp
	}
	data, err := os.ReadFile(fmt.Sprintf("%s/.git/HEAD", fp.fileDir))
	checker.Check(err)
	fp.branchName = strings.SplitAfter(string(data), "heads/")[1]
	return fp
}

func (fp *fileProcessor) getCommitTitle() *fileProcessor {
	if fp.commitTitle != "" {
		return fp
	}
	splitedBranchName := strings.SplitN(fp.branchName, "-", 2)
	name, taskId := splitedBranchName[0], splitedBranchName[1]
	fp.commitTitle = branch.NewBranch(strings.Split(name, "/")[0], taskId).GetCommitTitle()
	return fp
}

func (fp *fileProcessor) writeToFile() {
	file, err := os.Create(fmt.Sprintf("%s/.git/hooks/commit-msg", fp.fileDir))
	checker.Check(err)
	data := ReplaceDefaultValue(fp.commitTitle)
	file.WriteString(data)
	os.Chmod(file.Name(), 0775)
}
