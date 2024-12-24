import os
from typing import Literal


SAMPLE = """r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb"""


class Input:
    def __init__(self, filepath: str):
        self.stripes : dict[str, Literal[True]]
        self.stripe_length: set[int]
        self.longest_stripe: int
        self.puzzle: list[str]
        self.__post_init__(filepath)

    def __post_init__(self, filepath: str):
        with open(filepath, "r") as fp:
            content = fp.read()
        # content = SAMPLE
        self.stripes, self.puzzle = content.split("\n\n")
        self.scan_window = max([len(stripe) for stripe in self.stripes.split(", ")])
        self.stripe_length = set(map(len, self.stripes))
        self.stripes = {k: True for k in self.stripes.split(", ")}
        self.puzzle = self.puzzle.split("\n")


MEMO: dict[str, int] = {}


def match_stripe(stripes: dict[str, bool], puzzle_str: str) -> bool:
    if MEMO.get(puzzle_str):
        return MEMO[puzzle_str]
    if not puzzle_str:
        return True
    for word in stripes.keys():
        if puzzle_str.startswith(word):
            matched = match_stripe(stripes, puzzle_str.removeprefix(word))
            if matched:
                MEMO[puzzle_str] = matched
                return True
    MEMO[puzzle_str] = False
    return False


def part01(data: Input):
    num_solvable = 0
    for pzl in data.puzzle:
        print(f"{pzl=}", end=", ", flush=True)
        is_match = match_stripe(data.stripes, pzl)
        print(f"{is_match=}")
        if is_match:
            num_solvable += 1
    print(f"{num_solvable=}")


def part02(data: Input):
    ...


def main():
    print("""
U   U  SSS  EEEE       ggg      ooo
U   U S     E         g        o   o
U   U  SSS  EEE      g   gggg  o   o
u   U     S E         g   g g  o   o
 UUU   SSS  EEEE       ggg  g   ooo

    """)
    os._exit(0)
    data = Input("./day19/input.txt")
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()

