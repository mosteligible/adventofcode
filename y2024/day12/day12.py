import os
import traceback
from enum import Enum, auto
from pprint import pprint

SAMPLE = """
AAAA
BBCD
BBCC
EEEC
"""
SAMPLE = """
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
"""


class Direction(Enum):
    up = auto()
    down = auto()
    left = auto()
    right = auto()


def process_input():
    with open("./day12/input.txt", "r") as fp:
        content = fp.readlines()
    # content = SAMPLE.splitlines()
    content = [list(i.strip()) for i in content if i != ""]
    return content


def count_edges(
    positions: list[list[str]],
    row: int,
    col: int,
    curr_plant: str,
    rowlim: int,
    collim: int,
) -> int:
    edges = 0
    # up
    try:
        if row == 0:
            edges += 1
        elif row - 1 >= 0 and positions[row - 1][col] != curr_plant:
            edges += 1
        # down
        if row == rowlim - 1:
            edges += 1
        elif row + 1 < rowlim and positions[row + 1][col] != curr_plant:
            edges += 1
        # right
        if col == collim - 1:
            edges += 1
        elif col + 1 < collim and positions[row][col + 1] != curr_plant:
            edges += 1
        # left
        if col == 0:
            edges += 1
        elif col - 1 >= 0 and positions[row][col - 1] != curr_plant:
            edges += 1
    except IndexError:
        print(f"index error at: {row=}, {col=}")
        print(traceback.format_exc())
        os._exit(1)
    return edges


def get_edges(
    positions: list[list[str]],
    coord: tuple[int, int],
    edge_tracker: dict[tuple[int, int], int],
    visited_coords: set,
    row_lim: int,
    col_lim: int,
    # plant_nodes: dict[str, list[tuple[int, int]]]
) -> tuple[int, int]:
    if coord in visited_coords:
        return 0, 0
    visited_coords.add(coord)
    curr_plant = positions[coord[0]][coord[1]]
    # print(f"\n [X] {curr_plant} - adding coord: {coord}")
    num_edges = edge_tracker[coord]
    total_area = 1
    if num_edges == 4:
        return num_edges, total_area

    # can you go up:
    if coord[0] - 1 >= 0 and positions[coord[0] - 1][coord[1]] == curr_plant:
        edge, area = get_edges(
            positions=positions,
            coord=(coord[0] - 1, coord[1]),
            edge_tracker=edge_tracker,
            visited_coords=visited_coords,
            row_lim=row_lim,
            col_lim=col_lim,
        )
        num_edges += edge
        total_area += area
    # can you go down:
    if coord[0] + 1 < row_lim and positions[coord[0] + 1][coord[1]] == curr_plant:
        edge, area = get_edges(
            positions,
            (coord[0] + 1, coord[1]),
            edge_tracker,
            visited_coords,
            row_lim,
            col_lim,
        )
        num_edges += edge
        total_area += area
    # can you go left:
    if coord[1] - 1 >= 0 and positions[coord[0]][coord[1] - 1] == curr_plant:
        edge, area = get_edges(
            positions,
            (coord[0], coord[1] - 1),
            edge_tracker,
            visited_coords,
            row_lim,
            col_lim,
        )
        num_edges += edge
        total_area += area
    # can you go right:
    if coord[1] + 1 < col_lim and positions[coord[0]][coord[1] + 1] == curr_plant:
        edge, area = get_edges(
            positions,
            (coord[0], coord[1] + 1),
            edge_tracker,
            visited_coords,
            row_lim,
            col_lim,
        )
        num_edges += edge
        total_area += area

    return num_edges, total_area


def make_edge_count(
    positions: list[list[str]], rowlim: int, collim: int
) -> dict[tuple[int, int], int]:
    edge_map = {}
    for row_num, row in enumerate(positions):
        for col_num, col in enumerate(row):
            coordinate = (row_num, col_num)
            edge_map[coordinate] = count_edges(
                positions, row_num, col_num, col, rowlim, collim
            )
    return edge_map


def part01(data: list[list[str]]):
    visited_plant_coordinates: set = set()
    row_lim = len(data)
    col_lim = len(data[0])
    edge_tracker: dict[tuple[int, int], int] = make_edge_count(data, row_lim, col_lim)
    price = 0
    for row_num, row in enumerate(data):
        for col_num, plant in enumerate(row):
            coord = (row_num, col_num)
            if coord in visited_plant_coordinates:
                continue
            perimeter, area = get_edges(
                positions=data,
                coord=coord,
                edge_tracker=edge_tracker,
                visited_coords=visited_plant_coordinates,
                row_lim=row_lim,
                col_lim=col_lim,
            )
            price += perimeter * area
            # print(f"{plant=}, {perimeter=}, {area=}")

    print(f"price: {price}")


def get_continuous_edges(
    positions: list[list[str]],
    coord: tuple[int, int],
    edge_tracker: dict[tuple[int, int], int],
    visited_coords: set,
    row_lim: int,
    col_lim: int,
    positional_coords: set,
) -> tuple[int, int]:
    if coord in visited_coords:
        return 0, 0
    visited_coords.add(coord)
    curr_plant = positions[coord[0]][coord[1]]
    # print(f"\n [X] {curr_plant} - adding coord: {coord}")
    num_edges = edge_tracker[coord]
    if num_edges == 0:
        return 0, 1
    total_area = 1
    if num_edges == 4:
        return num_edges, total_area

    # can you go up:
    if coord[0] - 1 >= 0 and positions[coord[0] - 1][coord[1]] == curr_plant:
        edge, area = get_edges(
            positions=positions,
            coord=(coord[0] - 1, coord[1]),
            edge_tracker=edge_tracker,
            visited_coords=visited_coords,
            row_lim=row_lim,
            col_lim=col_lim,
        )
        num_edges += edge
        total_area += area
    # can you go down:
    if coord[0] + 1 < row_lim and positions[coord[0] + 1][coord[1]] == curr_plant:
        edge, area = get_edges(
            positions,
            (coord[0] + 1, coord[1]),
            edge_tracker,
            visited_coords,
            row_lim,
            col_lim,
        )
        num_edges += edge
        total_area += area
    # can you go left:
    if coord[1] - 1 >= 0 and positions[coord[0]][coord[1] - 1] == curr_plant:
        edge, area = get_edges(
            positions,
            (coord[0], coord[1] - 1),
            edge_tracker,
            visited_coords,
            row_lim,
            col_lim,
        )
        num_edges += edge
        total_area += area
    # can you go right:
    if coord[1] + 1 < col_lim and positions[coord[0]][coord[1] + 1] == curr_plant:
        edge, area = get_edges(
            positions,
            (coord[0], coord[1] + 1),
            edge_tracker,
            visited_coords,
            row_lim,
            col_lim,
        )
        num_edges += edge
        total_area += area

    return num_edges, total_area


def part02(data: list[list[str]]):
    visited_plant_coordinates: set = set()
    row_lim = len(data)
    col_lim = len(data[0])
    edge_tracker: dict[tuple[int, int], int] = make_edge_count(data, row_lim, col_lim)
    price = 0
    letter_tracker: dict = {}
    for row_num, row in enumerate(data):
        for col_num, col in enumerate(row):
            positional_coords = set()
            area = get_continuous_edges(
                positions=data,
                coord=(row_num, col_num),
                edge_tracker=edge_tracker,
                visited_coords=visited_plant_coordinates,
                row_lim=row_lim,
                col_lim=col_lim,
                positional_coords=positional_coords,
            )


def main():
    import time

    data = process_input()
    # pprint(data)
    start = time.time()
    part01(data)
    print(f"time taken part 01: {(time.time() - start)*1000}")
    part02(data)


if __name__ == "__main__":
    main()
