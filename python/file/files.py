import os

from dir.directory import get_current_dir


def get_new_file_path() -> str:
    return f'{get_current_dir()}/.git/hooks/commit-msg'


def generate_commit_msg_file(data: str):
    path = get_new_file_path()
    with open(path, 'w') as commit_msg_file:
        commit_msg_file.write(data)
    give_permissions_to_file(path)


def give_permissions_to_file(path):
    os.chmod(path, 0o775)
