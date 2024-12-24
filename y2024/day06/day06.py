import os
import time
from copy import deepcopy
from enum import auto, Enum
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
    # content = POSITIONS.splitlines()
    content = [list(line) for line in content]
    # pprint(content)
    return content


def get_guard_position_orientation(positions: list[list[str]]) -> tuple[tuple[int, int], Directions]:
    for index, row in enumerate(positions):
        try:
            col = row.index("^")
            return (index, col), Directions.north
        except:
            continue
    raise ValueError("Guard not found!")


def calculate_next_pos(curr_pos: tuple[int, int], direction: Directions) -> tuple[int, int]:
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
    positions: list[list[str]], direction: Directions, curr_pos: tuple[int, int],
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
            next_pos[0] >= rows or next_pos[0] < 0
            or next_pos[1] >= columns or next_pos[1] < 0
        ):
            # print(f"terminating, last_position: {curr_pos}")
            break
        if positions[next_pos[0]][next_pos[1]] == "#":
            # print(f"reorienting: {curr_pos=} | {direction=} -> {next_pos=}", end = " => ")
            direction = reorient(direction)
            # print(f"{curr_pos=} | {direction=} -> {next_pos=}")
            continue
        curr_pos = next_pos
        # print(curr_pos)
        # time.sleep(0.07)
    return moved_positions


def move_loop(
    positions: list[list[str]], direction: Directions, curr_pos: tuple[int, int]
) -> True:
    moved_positions = {}
    rows, columns = len(positions), len(positions[0])
    next_pos = curr_pos
    while True:
        # tick = get_tick(direction)
        # positions[curr_pos[0]][curr_pos[1]] = tick
        # os.system("clear")
        # print_pos(positions)
        if moved_positions.get(curr_pos, None):
            moved_positions[curr_pos] += 1
        else:
            moved_positions[curr_pos] = 1
        next_pos = calculate_next_pos(curr_pos, direction)
        if (
            next_pos[0] >= rows or next_pos[0] < 0
            or next_pos[1] >= columns or next_pos[1] < 0
        ):
            # print(f"terminating, last_position: {curr_pos}")
            # print(f"-- {moved_positions=}")
            # print(f"-- {check=}")
            # time.sleep(1.5)
            return False
        # print(f"before # check")
        if positions[next_pos[0]][next_pos[1]] == "#":
            # print(f"reorienting: {curr_pos=} | {direction=} -> {next_pos=}", end = " => ")
            direction = reorient(direction)
            # print(f"{curr_pos=} | {direction=} -> {next_pos=}")
            continue
        # print(f"after # check")
        curr_pos = next_pos
        check = [v for _, v in moved_positions.items() if v > 3]
        # print(f"-- {check=}")
        # print(f"-- {moved_positions=}")
        if sum(check) > len(moved_positions) * 2:
            print(f"loop identified, all position traversed at least once!")
            # time.sleep(0.5)
            return True
        # time.sleep(0.1)
    return moved_positions


def part01(positions) -> dict[tuple[int, int], Literal[True]]:
    pos, direction = get_guard_position_orientation(positions)
    moved_positions = move(positions, direction, pos)
    # not at start
    return moved_positions

def part02(positions: list[list[str]]):
    pos, direction = get_guard_position_orientation(positions)
    moved_positions = part01(deepcopy(positions))
    time.sleep(1)
    moved_positions.pop(pos)
    loop_count = 0
    counts = 0
    locations = list(moved_positions.keys())
    for k, _ in moved_positions.items():
        positions_copy = deepcopy(positions)
        positions_copy[k[0]][k[1]] = "#"
        is_loop = move_loop(positions_copy, direction, pos)
        # print("move loop?", is_loop)
        counts += 1
        if is_loop:
            loop_count += 1
            print(f"num loops: {loop_count} - traversed: {counts}")


def main():
    data = process_input()
    part02(data)


if __name__ == "__main__":
    main()
