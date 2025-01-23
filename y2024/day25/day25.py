from typing import List, Tuple

SAMPLE = """#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####"""
TARGET = 7


class Combination:
    def __init__(self, pattern: str) -> None:
        self.pattern = pattern
        self.combination: List[int]
        self.parse_pattern()

    def parse_pattern(self) -> "Combination":
        lines = self.pattern.splitlines()
        lines = [list(i) for i in lines]
        # print(f"{lines=}")
        num_rows = len(lines)
        num_cols = len(lines[0])
        self.combination = []
        for col in range(num_cols):
            fill_count = "".join([lines[row][col] for row in range(num_rows)])
            self.combination.append(fill_count.count("#"))
        # print(f"{self.pattern}, {self.combination}")
        return self


def process_input() -> Tuple[List, List]:
    with open("./day25/input.txt", "r") as fp:
        content = fp.read()
    # content = SAMPLE
    patterns = content.split("\n\n")
    locks = []
    pins = []
    for p in patterns:
        combination = Combination(pattern=p)
        if p.startswith("#####"):
            locks.append(combination)
        else:
            pins.append(combination)
    return locks, pins


def part01(locks: List[Combination], pins: List[Combination]) -> None:
    num_cols = len(locks[0].combination)
    total = 0
    for loc in locks:
        for pin in pins:
            mix = [loc.combination[i] + pin.combination[i] for i in range(num_cols)]
            if sum([1 for i in mix if i <= TARGET]) == num_cols:
                total += 1

    print(f"num matches= {total}")


def main():
    locks, pins = process_input()
    part01(locks, pins)


if __name__ == "__main__":
    main()
