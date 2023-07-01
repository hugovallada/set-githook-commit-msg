import sys
import os

def get_current_dir() -> str:
    return os.getcwd()


def dividir_branch(value: str) -> list[str]:
    return value.split("heads/")

class Branch:
    type: str
    id: str

    def extract_id(self) -> str:
        return self.id.split("-")[-1]
    
    def write_base_message(self) -> str:
        return f'{self.type}_{self.extract_id()}'

value = """#!/bin/sh
branch_name=$VALUE_TO_BE_REPLACED
commit_msg=$(cat $1)
echo "$branch_name: $commit_msg" > $1"""


def set_default_branch():
    branch  = set_branch_class()
    commit_data = replace_value(value, branch)
    write_to_file(commit_data)
    os.chmod(f'{get_current_dir()}/.git/hooks/commit-msg', 0o775)
    

def set_task_id(id: str):
    new_value = value.replace('$VALUE_TO_BE_REPLACED', id)
    write_to_file(new_value)
    os.chmod(f'{get_current_dir()}/.git/hooks/commit-msg', 0o775)

def set_branch_class() -> Branch:
    with open(f'{get_current_dir()}/.git/HEAD') as head:
       line = head.readlines()[0]
       extracted_data = line.split('heads/')[-1]
       branch = Branch()
       branch.type, branch.id = extracted_data.split('/')
       return branch
    
def replace_value(value: str, branch: Branch) -> str:
    return value.replace('$VALUE_TO_BE_REPLACED',branch.write_base_message())

def write_to_file(data: str):
    with open(f'{get_current_dir()}/.git/hooks/commit-msg', 'w') as f:
            f.write(data)


if len(sys.argv) > 1:
    set_task_id(sys.argv[1])
else:
    set_default_branch()
    


    