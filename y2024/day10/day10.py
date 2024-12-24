import sys
from pprint import pprint

sys.setrecursionlimit(1000000)

"""
8 9 0 1 0 1 2 3
7 8 1 2 1 8 7 4
8 7 4 3 0 9 6 5
9 6 5 4 9 8 7 4
4 5 6 7 8 9 0 3
3 2 0 1 9 0 1 2
0 1 3 2 9 8 0 1
1 0 4 5 6 7 3 2
"""

SAMPLE = """89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732"""
# SAMPLE = """9990999
# 9991999
# 9992999
# 6543456
# 7111117
# 8111118
# 9111119"""
# SAMPLE="""0123
# 1234
# 8765
# 9876"""
# SAMPLE="""7790339
# 8881798
# 8882227
# 6543456
# 7653987
# 8761111
# 9871111
# """


class Node:
    def __init__(self, curr: tuple[str, str]):
        self.current = curr
        self.left: "Node" = None
        self.right: "Node" = None
        self.up: "Node" = None
        self.down: "Node" = None

    def add_next(self, next: "Node") -> None:
        if self.next is None:
            self.next = set()
        self.next.add(next)

    def __eq__(self, value: "Node") -> bool:
        return self.current == value.current

    def __hash__(self):
        return hash(self.current)

    def __str__(self) -> str:
        return f"""
            {self.up}
              ^
              |
{self.left} <- {self.current} -> {self.right}
              |
              v
            {self.down}
"""


ORIG_DEST_TRACKER: dict[tuple[int, int], list[tuple[int, int]]] = {}


def process_input():
    with open("./day10/input.txt", "r") as fp:
        content = fp.readlines()
    # content = SAMPLE.splitlines()
    content = [list(i.strip()) for i in content]
    # pprint(content)
    return content


def trailheads(
    data: list[list[str]], row: int, col: int, prev_no: str, size: tuple[int, int],
    trail_tracker: list, origin: tuple[int, int], start: bool,
) -> int:
    curr_no = data[row][col]
    curr_node = Node(curr=(row, col))
    num_right, num_left, num_down, num_up = 0, 0, 0, 0

    if not start:
        if ord(curr_no) - ord(prev_no) != 1:
            return 0
        # if prev_no == "0" and ord(curr_no) - ord(prev_no) != 1
        if data[row][col] == "9" and prev_no != "0":
            if (row, col) in ORIG_DEST_TRACKER[origin]:
                return 0
            # ORIG_DEST_TRACKER[origin].append((row, col))
            trail_tracker.append((row, col))
            print(f"{curr_no=}, {row=}, {col=}, num_solve={1}")
            return 1
    else:
        start = False
    # if (row, col) in trail_tracker:
    #     return 0
    # check up
    if row - 1 >= 0:
        num_up = trailheads(
            data, row=row-1, col=col, prev_no=curr_no,
            size=size, trail_tracker=trail_tracker,
            origin=origin, start=start
        )
        # print(f"going up, {row=}, {col=}")
        if num_up > 0:
            # print(f"{curr_no=}, {row=}, {col=}, {num_up=}")
            trail_tracker.append((row, col))
        # num_trailheads += num_up
    # check down
    if row + 1 < size[0]:
        # print(f"going down, {row=}, {col=}")
        num_down = trailheads(
            data, row=row+1, col=col, prev_no=curr_no,
            size=size, trail_tracker=trail_tracker,
            origin=origin, start=start
        )
        if num_down > 0:
            # print(f"{curr_no=}, {row=}, {col=}, {num_down=}")
            trail_tracker.append((row, col))
        # num_trailheads += num_down
    # check left
    if col - 1 >= 0:
        num_left = trailheads(
            data, row=row, col=col-1, prev_no=curr_no,
            size=size, trail_tracker=trail_tracker,
            origin=origin, start=start
        )
        if num_left > 0:
            # print(f"{curr_no=}, {row=}, {col=}, {num_left=}")
            trail_tracker.append((row, col))
        # num_trailheads += num_left
    # check right
    if col + 1 < size[1]:
        # print(f"going right, {row=}, {col=}")
        num_right = trailheads(
            data, row=row, col=col+1, prev_no=curr_no,
            size=size, trail_tracker=trail_tracker,
            origin=origin, start=start
        )
        if num_right > 0:
            trail_tracker.append((row, col))
            # print(f"{curr_no=}, {row=}, {col=}, {num_right=}")
        # num_trailheads += num_right
    # print(f"NODE: {curr_no=}, {prev_no=}, {row=}, {col=}, total={num_up + num_down + num_left + num_right}")
    return num_up + num_down + num_left + num_right


def bfs_trailheads(data: list[list[str]], size: tuple[int, int], curr_coord: tuple[int, int], origin: tuple[int, int]) -> int:
    o1, o2 = origin
    c1, c2 = curr_coord
    if data[o1][o2] != "0" and data[c1][c2] == "9":
        if curr_coord in ORIG_DEST_TRACKER[origin]:
            return 0
        ORIG_DEST_TRACKER[origin].append(curr_coord)
        return 1
    # can move up?
    if c1 - 1 >= 0:
        ...


def part01(data):
    trail_coordinate_tracker = []
    destination_tracker = set()
    size = (len(data), len(data[0]))
    print(f"{size=}")
    num_trails = 0
    # ORIG_DEST_TRACKER[(6,6)] = []
    # start = True
    # num_trails = trailheads(
    #             data, 6, 6, "0", size=size,
    #             trail_tracker=trail_coordinate_tracker,
    #             origin=(6,6), start=start
    #         )
    for row_num, row in enumerate(data):
        for col_num, col in enumerate(row):
            if col != "0":
                continue
            origin = (row_num, col_num)
            ORIG_DEST_TRACKER[origin] = []
            trail_coordinate_tracker.append((row_num, col_num))
            n_trail = trailheads(
                data, row_num, col_num, "0", size=size,
                trail_tracker=trail_coordinate_tracker,
                origin=origin, start=True
            )
            print(f"0 at ({row_num}, {col_num}) has {n_trail} trails")
            num_trails += n_trail
    print(f"num_trailheads: {num_trails}")
    print()
    # print(trail_coordinate_tracker)

    # for row_num, row in enumerate(data):
    #     for col_num, col in enumerate(data):
    #         if (row_num, col_num) not in trail_coordinate_tracker:
    #             data[row_num][col_num] = "."
    # pprint(data)


def part02(data):
    ...


def main():
    data = process_input()
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()

