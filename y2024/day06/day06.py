import os
import time
from copy import deepcopy
from enum import Enum, auto
from pprint import pprint
from typing import Literal


class Directions(Enum):
    east = auto()
    west = auto()
    north = auto()
    south = auto()


POSITIONS = """....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."""


def process_input():
    with open("./day06/input.txt", "r") as fp:
        content = fp.readlines()
    content = POSITIONS.splitlines()
    content = [list(line) for line in content]
    # pprint(content)
    return content


def get_guard_position_orientation(
    positions: list[list[str]],
) -> tuple[tuple[int, int], Directions]:
    for index, row in enumerate(positions):
        try:
            col = row.index("^")
            return (index, col), Directions.north
        except:
            continue
    raise ValueError("Guard not found!")


def calculate_next_pos(
    curr_pos: tuple[int, int], direction: Directions
) -> tuple[int, int]:
    if direction == Directions.east:
        next_pos = curr_pos[0], curr_pos[1] + 1
    elif direction == Directions.west:
        next_pos = curr_pos[0], curr_pos[1] - 1
    elif direction == Directions.north:
        next_pos = curr_pos[0] - 1, curr_pos[1]
    else:
        next_pos = curr_pos[0] + 1, curr_pos[1]
    return next_pos


def reorient(direction: Directions) -> tuple[tuple[int, int], Directions]:
    if direction == Directions.north:
        direction = Directions.east
    elif direction == Directions.east:
        direction = Directions.south
    elif direction == Directions.south:
        direction = Directions.west
    else:
        direction = Directions.north
    return direction


def get_tick(direction: Directions) -> str:
    if direction == Directions.north:
        return "^"
    elif direction == Directions.east:
        return ">"
    elif direction == Directions.south:
        return "v"
    else:
        return "<"


def print_pos(positions: list[list[str]]) -> None:
    pos = ["".join(i) for i in positions]
    pos = "\n".join(pos)
    print(pos)


def move(
    positions: list[list[str]],
    direction: Directions,
    curr_pos: tuple[int, int],
) -> dict[tuple[int, int], Literal[True]]:
    moved_positions = {}
    rows, columns = len(positions), len(positions[0])
    next_pos = curr_pos
    while True:
        # tick = get_tick(direction)
        # positions[curr_pos[0]][curr_pos[1]] = tick
        # os.system("clear")
        # print_pos(positions)

        moved_positions[curr_pos] = True
        next_pos = calculate_next_pos(curr_pos, direction)
        if (
            next_pos[0] >= rows
            or next_pos[0] < 0
            or next_pos[1] >= columns
            or next_pos[1] < 0
        ):
            break
        if positions[next_pos[0]][next_pos[1]] == "#":
            direction = reorient(direction)
            continue
        curr_pos = next_pos
        # time.sleep(0.25)
    return moved_positions


def move_loop(
    positions: list[list[str]], direction: Directions, curr_pos: tuple[int, int]
) -> True:
    moved_positions: dict[tuple[int, int, int], Literal[True]] = {}
    rows, columns = len(positions), len(positions[0])
    next_pos = curr_pos
    while True:
        # tick = get_tick(direction)
        # positions[curr_pos[0]][curr_pos[1]] = tick
        # os.system("clear")
        # print_pos(positions)
        mp_key = (*curr_pos, direction.name)
        if moved_positions.get(mp_key, None):
            moved_positions[mp_key] += 1
            return True
        else:
            moved_positions[mp_key] = 1
        next_pos = calculate_next_pos(curr_pos, direction)
        if (
            next_pos[0] >= rows
            or next_pos[0] < 0
            or next_pos[1] >= columns
            or next_pos[1] < 0
        ):
            return False
        if positions[next_pos[0]][next_pos[1]] == "#":
            direction = reorient(direction)
            continue
        curr_pos = next_pos
        check = [v for _, v in moved_positions.items() if v > 3]
        if sum(check) > len(moved_positions) * 5:
            print(f"loop identified, all position traversed at least once!")
            return True
    return moved_positions


def part01(positions) -> dict[tuple[int, int], Literal[True]]:
    pos, direction = get_guard_position_orientation(positions)
    moved_positions = move(positions, direction, pos)
    print("Part 01:", len(moved_positions))
    # not at start
    return moved_positions


def part02(positions: list[list[str]]):
    moved_positions = part01(deepcopy(positions))
    pos, direction = get_guard_position_orientation(positions)
    loop_count = 0
    pkeys = list(moved_positions.keys())
    perv_pos_index = 0
    for _, k in enumerate(pkeys[1:]):
        positions[k[0]][k[1]] = "#"
        is_loop = move_loop(positions, direction, pos)
        positions[k[0]][k[1]] = "."
        perv_pos_index += 1
        if is_loop:
            loop_count += 1


def main():
    data = process_input()
    part02(data)


if __name__ == "__main__":
    main()
