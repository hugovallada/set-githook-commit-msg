package branch

import "fmt"

type branch struct {
	branchType string
	branchID   string
}

func NewBranch(branchType string, branchID string) branch {
	return branch{
		branchType: branchType,
		branchID:   branchID,
	}
}

func (b branch) GetCommitTitle() string {
	return fmt.Sprintf("%s_%s", b.branchType, b.branchID)
}
