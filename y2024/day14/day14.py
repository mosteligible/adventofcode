import time
from copy import deepcopy
from dataclasses import dataclass
from pprint import pprint

SAMPLE = """p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
"""
# SAMPLE = "p=2,4 v=2,-3"


@dataclass
class Coordinate:
    row: int
    col: int


@dataclass
class Robot:
    position: Coordinate
    velocity: Coordinate


def process_input() -> list[Robot]:
    with open("./day14/input.txt", "r") as fp:
        content = fp.readlines()
    # content = SAMPLE.splitlines()

    robots = []
    for line in content:
        line = line.strip()
        pos, vel = line.split()
        pos = pos[2:]
        vel = vel[2:]
        pos = pos.split(",")
        vel = vel.split(",")
        robot = Robot(
            position=Coordinate(row=int(pos[1]), col=int(pos[0])),
            velocity=Coordinate(row=int(vel[1]), col=int(vel[0])),
        )
        robots.append(robot)
    # print(f"{robots=}")
    return robots


def get_grid(
    robots: list[Robot], cols: int = 101, rows: int = 103
) -> dict[tuple[int, int], list[Robot]]:
    grid: dict[tuple[int.int], list[Robot]] = {}
    for row_num in range(rows):
        for col_num in range(cols):
            grid[(row_num, col_num)] = []

    for r in robots:
        grid[(r.position.row, r.position.col)].append(r)

    return grid


def show_grid(
    grid: dict[tuple[int, int], list[Robot]], rows: int, cols: int, final: bool = False
) -> None:
    output: list[list[int]] = [[0] * cols for _ in range(rows)]
    for k, v in grid.items():
        output[k[0]][k[1]] += len(v)
    for row_num, row in enumerate(output):
        if row_num == rows // 2 and final:
            print()
            continue
        for col_num, col in enumerate(row):
            if col_num == cols // 2 and final:
                print(" ", end="")
                continue
            if col == 0:
                print(".", end="")
            else:
                print(col, end="")
        print()

    # pprint(output)


def robot_move(
    grid: dict[tuple[int, int], list[Robot]], coordinate: tuple[int, int]
) -> None:
    ...


def move_all_robots(
    grid: dict[tuple[int, int], list[Robot]], row_lim: int, col_lim: int
) -> None:
    new_dict: dict[tuple[int, int], list[Robot]] = {}
    for coordinate, robots in grid.items():
        num_robots = len(robots)
        while num_robots > 0:
            r = robots.pop(0)
            new_row = coordinate[0] + r.velocity.row
            new_col = coordinate[1] + r.velocity.col
            if new_row >= row_lim:
                new_row -= row_lim
            elif new_row < 0:
                new_row += row_lim
            if new_col >= col_lim:
                new_col -= col_lim
            elif new_col < 0:
                new_col += col_lim
            r.position.row, r.position.col = new_row, new_col
            # print(f"row={r}")
            # causes in place replacement
            if new_dict.get((new_row, new_col), None) is not None:
                new_dict[(new_row, new_col)].append(r)
            else:
                new_dict[(new_row, new_col)] = [r]
            num_robots -= 1
    grid.update(new_dict)


def get_safety_factor(
    grid: dict[tuple[int, int], list[Robot]], size: tuple[int, int]
) -> int:
    row_to_avoid = size[0] // 2
    col_to_aviod = size[1] // 2
    safety_factor = 1
    q1, q2, q3, q4 = 0, 0, 0, 0
    for k, v in grid.items():
        if k[0] < row_to_avoid and k[1] < col_to_aviod:
            q1 += len(v)
        elif k[0] < row_to_avoid and k[1] > col_to_aviod:
            q2 += len(v)
        elif k[0] > row_to_avoid and k[1] < col_to_aviod:
            q3 += len(v)
        elif k[0] > row_to_avoid and k[1] > col_to_aviod:
            q4 += len(v)
    print(f"{q1=}, {q2=}, {q3=}, {q4=}")
    return q1 * q2 * q3 * q4


def part01(grid: dict[tuple[int, int], list[Robot]], size: tuple[int, int]):
    num_seconds = 100
    for _ in range(num_seconds):
        move_all_robots(grid=grid, row_lim=size[0], col_lim=size[1])

    print("*" * 42)
    # show_grid(grid=grid, rows=size[0], cols=size[1], final=True)

    print(f" [o] {get_safety_factor(grid, size)=}")


def part02_eyeball_it(grid: dict[tuple[int, int], list[Robot]], size: tuple[int, int]):
    first = True
    for num_seconds in range(10000):
        move_all_robots(
            grid=grid,
            row_lim=size[0],
            col_lim=size[1],
        )
        print(f"{'*'*25} seconds: {num_seconds} {'*'*25}")
        if num_seconds > 7100:
            show_grid(grid=grid, rows=size[0], cols=size[1])
            if first:
                time.sleep(1)
                first = False

            time.sleep(0.25)


def check_linearity(
    grid: dict[tuple[int, int], list[Robot]], rows: int, cols: int
) -> bool:
    for row_num in range(rows):
        row = "".join([str(len(grid[(row_num, col_num)])) for col_num in range(cols)])
        if "1" * 16 in row:
            return True

    return False


def part02(grid: dict[tuple[int, int], list[Robot]], size: tuple[int, int]):
    for num_seconds in range(10000):
        move_all_robots(
            grid=grid,
            row_lim=size[0],
            col_lim=size[1],
        )
        if check_linearity(grid, size[0], size[1]):
            # show_grid(grid=grid, rows=size[0], cols=size[1], final=True)
            print(f" [o] {num_seconds=}")
            break


def main():
    robots = process_input()
    size = (103, 101)
    # size = (7, 11)
    grid = get_grid(robots=robots, rows=size[0], cols=size[1])
    print("             PART 01")
    start = time.time()
    part01(grid=deepcopy(grid), size=size)
    print(f"time taken: {time.time()-start}")
    start = time.time()
    print("PART 02")
    part02(grid=grid, size=size)
    print(f"time taken: {time.time() - start}")


if __name__ == "__main__":
    main()
