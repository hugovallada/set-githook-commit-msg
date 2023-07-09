import sys

from entity.branch import Branch
from file.datafile import replace_value
from file.files import generate_commit_msg_file


def set_default_branch():
    branch = Branch.generate_branch()
    replaced_value = replace_value(branch.get_commit_title())
    generate_commit_msg_file(replaced_value)


def set_task_id(commit_title: str):
    replaced_value = replace_value(commit_title)
    generate_commit_msg_file(replaced_value)


if len(sys.argv) > 1:
    set_task_id(sys.argv[1])
else:
    set_default_branch()
