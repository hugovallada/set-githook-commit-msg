package main

import (
	"fmt"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getBranchName(branch string) string {
	return strings.SplitAfter(branch, "heads/")[1]
}

func writeToFile(path, data string) {
	file, err := os.Create(fmt.Sprintf("%s/.git/hooks/commit-msg", path))
	checkError(err)
	file.WriteString(data)
	os.Chmod(file.Name(), 0775)
}

var value string = `#!/bin/sh
branch_name=$VALUE_TO_BE_REPLACED
commit_msg=$(cat $1)
echo "$branch_name: $commit_msg" > $1`

func main() {
	currentDir, err := os.Getwd()
	checkError(err)
	data, err := os.ReadFile(fmt.Sprintf("%s/.git/HEAD", currentDir))
	checkError(err)
	dataAsString := string(data)
	branchName := getBranchName(dataAsString)
	replacedValue := strings.Replace(value, "$VALUE_TO_BE_REPLACED", branchName, 1)
	writeToFile(currentDir, replacedValue)
}
