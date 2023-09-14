package files

import (
	"strings"
)

func GetBranchName(branch string) string {
	return strings.SplitAfter(branch, "heads/")[1]
}
