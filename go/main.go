package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hugovallada/set-branch/checker"
	"github.com/hugovallada/set-branch/files"
)

var title string
var start bool
var destroy bool
var branchName string
var deleteBranch bool

func init() {
	flag.StringVar(&title, "t", "", "Commit Title")
	flag.BoolVar(&destroy, "d", false, "Delete commit message hook")
	flag.BoolVar(&start, "i", false, "Start new branch")
	flag.StringVar(&branchName, "b", "", "New branch")
	flag.BoolVar(&deleteBranch, "db", false, "Delete current branch")
	flag.Parse()
}

func main() {
	currentDir, err := os.Getwd()
	checker.Check(err)
	if destroy {
		os.Remove(fmt.Sprintf("%s/.git/hooks/commit-msg", currentDir))
		if deleteBranch {
			if branchName == "" {
				branchName = "main"
			}
			data, err := os.ReadFile(fmt.Sprintf("%s/.git/HEAD", currentDir))
			checker.Check(err)
			cmd := exec.Command("git", "switch", branchName)
			err = cmd.Run()
			checker.Check(err)
			cmd = exec.Command("git", "branch", "-D", strings.TrimSpace(files.GetBranchName(string(data))))
			err = cmd.Run()
			checker.Check(err)
		}
		return
	}
	if start {
		if branchName != "" {
			cmd := exec.Command("git", "switch", "-c", branchName)
			err := cmd.Run()
			checker.Check(err)
		} else {
			fmt.Println("You need to specify a branch name with the flag -b")
			return
		}
	}
	files.NewFileProcessor(currentDir, title).Execute()
}
