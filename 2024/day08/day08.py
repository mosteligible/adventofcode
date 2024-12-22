from copy import deepcopy
from pprint import pprint

"""
c0 = (1,1)
c1 = (5,2)
delta = c0 - c1 = (-4, -1)
c0 + delta = (-3, 0)
c0 - 2 x delta = (5, 2)
"""

EPSILON = 0.009
POS = []
ANTINODE_POS = set()
ANTENNA_POS = []
SAMPLE = """............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............"""


def process_input():
    with open("./day08/input.txt", "r") as fp:
        content = fp.readlines()
    # content = SAMPLE.splitlines()
    content = [list(i.strip()) for i in content]
    return content


def get_antenna_coordinates(positions: list[list[str]]) -> dict[str, tuple[int, int]]:
    coordinates: dict[str, list] = {}
    antenna_count = 0
    for row_num, row in enumerate(positions):
        for col_num, col in enumerate(row):
            if col != ".":
                antenna_count += 1
                antenna = coordinates.get(col, "")
                if antenna:
                    coordinates[col].append((row_num, col_num))
                else:
                    coordinates[col] = [(row_num, col_num),]
    return coordinates


def coordinates_difference(c1: tuple[int, int], c2: tuple[int, int]) -> tuple[int, int]:
    delta_x = c1[0] - c2[0]
    delta_y = c1[1] - c2[1]
    return delta_x, delta_y


def coordinates_add(c1: tuple[int, int], c2: tuple[int, int]) -> tuple[int, int]:
    add_x = c1[0] + c2[0]
    add_y = c1[1] + c2[1]
    return add_x, add_y


def distance(c1: tuple[int, int], c2: tuple[int, int]) -> float:
    dist = ((c1[0]-c2[0])**2 + (c1[1]-c2[1])**2)**0.5
    return round(dist, 2)


def calculate_antinode_pos(coordinates: list[tuple[int, int]], num_rows: int, num_cols: int, antenna_pos: list) -> int:
    global POS
    for index, c0 in enumerate(coordinates):
        for c1 in coordinates[index + 1:]:
            delta = coordinates_difference(c0, c1)
            potential_antinode_0 = coordinates_add(c0, delta)
            potential_antinode_1 = coordinates_difference(c0, (2*delta[0], 2*delta[1]))
            if (potential_antinode_1[0] >= 0 and potential_antinode_1[0] < num_rows
                and potential_antinode_1[1] >= 0 and potential_antinode_1[1] < num_cols
                and potential_antinode_1 not in antenna_pos):
                POS[potential_antinode_1[0]][potential_antinode_1[1]] = "#"
                ANTINODE_POS.add(potential_antinode_1)
            if (potential_antinode_0[0] >= 0 and potential_antinode_0[0] < num_rows
                and potential_antinode_0[1] >= 0 and potential_antinode_0[1] < num_cols
                and potential_antinode_0 not in antenna_pos
                ):
                POS[potential_antinode_0[0]][potential_antinode_0[1]] = "#"
                ANTINODE_POS.add(potential_antinode_0)


def part01(data: list[list[str]]):
    print(f"{'*'*20}  PART 01  {'*'*20}")
    global POS
    num_rows = len(data)
    num_cols = len(data[0])
    POS = deepcopy(data)
    antenna_coordinates = get_antenna_coordinates(data)
    antenna_pos = set()
    for _, v in antenna_coordinates.items():
        antenna_pos.update(*v)
    for antenna, coordinates in antenna_coordinates.items():
        calculate_antinode_pos(coordinates, num_cols=num_cols, num_rows=num_rows, antenna_pos=antenna_pos)
    ## uncomment below to print final grid status
    # POS = ["".join(i) for i in POS]
    # pprint({POS})
    print(f"unique num_antinodes={len(ANTINODE_POS)}")


def coordinate_out_of_bounds(c: tuple[int, int], row_size: int, col_size: int) -> bool:
    if c[0] >= 0 and c[0] < row_size and c[1] >= 0 and c[1] < col_size:
        return False
    return True

def calculate_linear_coords(
    c0: tuple[int, int], delta: tuple[int, int], row_limit: int, col_limit: int, antenna_pos: list, antinode_holder: set
) -> int:
    global POS
    multiplier = 1
    num_antinodes = 0
    while True:
        potential_antinode_0 = coordinates_add(c0, (delta[0] * multiplier, delta[1] * multiplier))
        potential_antinode_1 = coordinates_difference(
            c0, (delta[0] * (multiplier + 1), delta[1] * (multiplier + 1))
        )
        p01oob = coordinate_out_of_bounds(potential_antinode_0, row_limit, col_limit)
        p02oob = coordinate_out_of_bounds(potential_antinode_1, row_limit, col_limit)
        if not p01oob and potential_antinode_0:
            antinode_holder.add(potential_antinode_0)
            POS[potential_antinode_0[0]][potential_antinode_0[1]] = "#"
            num_antinodes += 1
        if not p02oob and potential_antinode_1:
            antinode_holder.add(potential_antinode_1)
            POS[potential_antinode_1[0]][potential_antinode_1[1]] = "#"
            num_antinodes += 1
        multiplier += 1
        if p01oob and p02oob:
            return num_antinodes


def pt02_calculate_antinode_pos(
    coordinates: list[tuple[int, int]], num_rows: int, num_cols: int, antenna_pos: list, antinode_holder: set
) -> int:
    num_antinodes = 0
    for index, c0 in enumerate(coordinates):
        for c1 in coordinates[index+1:]:
            antinode_holder.add(c0)
            antinode_holder.add(c1)
            delta = coordinates_difference(c0, c1)
            num_antinodes += calculate_linear_coords(
                c0, delta=delta, row_limit=num_rows, col_limit=num_cols, antenna_pos=antenna_pos, antinode_holder=antinode_holder
            )
    return num_antinodes


def part02(data):
    print(f"{'*'*20}  PART 02  {'*'*20}")
    global ANTINODE_POS, POS
    POS = deepcopy(data)
    antinode_pos = set()
    num_rows = len(data)
    num_cols = len(data[0])
    antenna_coordinates = get_antenna_coordinates(data)
    antenna_pos = set()
    for _, v in antenna_coordinates.items():
        antenna_pos.update(*v)
    for _, coordinates in antenna_coordinates.items():
        pt02_calculate_antinode_pos(
            coordinates, num_cols=num_cols, num_rows=num_rows, antenna_pos=antenna_pos, antinode_holder=antinode_pos
        )
    ## uncomment below to print final grid status
    # POS = ["".join(i) for i in POS]
    # pprint(POS)
    print(f"unique num antinodes: {len(antinode_pos)}")


def main():
    data = process_input()
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()
