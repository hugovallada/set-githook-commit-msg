from __future__ import annotations

from dataclasses import dataclass

from dir.directory import get_current_dir


@dataclass
class Branch:
    type: str
    id: str

    def extract_id(self) -> str:
        return self.id.split('-')[-1]

    def get_commit_title(self):
        return f'{self.type}_{self.extract_id()}'

    @staticmethod
    def generate_branch() -> Branch:
        with open(f'{get_current_dir()}/.git/HEAD') as head_file:
            line = head_file.readlines()[0]
            new_type, new_id = line.split('heads/')[-1].split('/')
            return Branch(new_type, new_id)
