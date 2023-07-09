package files

import "strings"

var value string = `#!/bin/sh
branch_name=$VALUE_TO_BE_REPLACED
commit_msg=$(cat $1)
echo "$branch_name: $commit_msg" > $1`

func ReplaceDefaultValue(newValue string) string {
	return strings.Replace(value, "$VALUE_TO_BE_REPLACED", newValue, 1)
}
