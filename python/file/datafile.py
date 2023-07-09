value = """#!/bin/sh
branch_name=$VALUE_TO_BE_REPLACED
commit_msg=$(cat $1)
echo "$branch_name: $commit_msg" > $1"""


def replace_value(replacer: str) -> str:
    return value.replace('$VALUE_TO_BE_REPLACED', replacer)
