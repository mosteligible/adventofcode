import json
import time
from collections import deque
from heapq import heapify, heappop, heappush
from pprint import pprint


SAMPLE = """###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############"""


class Coordinate:
    def __init__(self, row: int, col: int, value: str = None) -> None:
        self.row: int = row
        self.col: int = col
        self.value: str = value

    def next_nodes(self, row_lim: int, col_lim: int, grid: list[list[str]], obstacle: str = "#") -> list["Coordinate"]:
        retval = []
        # up
        if self.row - 1 >= 0 and grid[self.row - 1][self.col] != obstacle:
            retval.append(Coordinate(self.row - 1, self.col, grid[self.row - 1][self.col]))
        # down
        if self.row + 1 < row_lim and grid[self.row + 1][self.col] != obstacle:
            retval.append(Coordinate(self.row + 1, self.col, grid[self.row + 1][self.col]))
        # left
        if self.col - 1 >= 0 and grid[self.row][self.col - 1] != obstacle:
            retval.append(Coordinate(self.row, self.col - 1, grid[self.row][self.col - 1]))
        # right
        if self.col + 1 < col_lim and grid[self.row][self.col + 1] != obstacle:
            retval.append(Coordinate(self.row, self.col + 1, grid[self.row][self.col + 1]))
        return retval

    def __hash__(self) -> int:
        return hash((self.row, self.col))

    def __lt__(self, value: "Coordinate") -> bool:
        return (self.row, self.col) < (value.row, value.col)

    def __gt__(self, value: "Coordinate") -> bool:
        return (self.row, self.col) > (value.row, value.col)

    def __eq__(self, value: "Coordinate") -> bool:
        return (self.row, self.col) == (value.row, value.col)

    def __repr__(self) -> str:
        return f"Coordinate({self.row}, {self.col}, {self.value})"

class Input:
    def __init__(self, filepath: str):
        self.row_limit: int
        self.col_limit: int
        self.start_coordinate: Coordinate = None
        self.end_coordinate: Coordinate = None
        self.obstacles: set[Coordinate] = set()
        self.grid: list[list[str]] = self.process_aoc_input(filepath)

    def process_aoc_input(self, filepath: str = "./day20/input.txt") -> list[list[str]]:
        with open(filepath, "r") as fp:
            content = fp.readlines()
        # content = SAMPLE.splitlines()
        self.grid = [list(x.strip()) for x in content]
        self.get_start_end_coordinates()
        return self

    def get_start_end_coordinates(self) -> tuple[Coordinate, Coordinate]:
        maxRow = len(self.grid)
        maxCol = len(self.grid[0])
        for row in range(len(self.grid)):
            for col in range(len(self.grid[0])):
                if self.grid[row][col] == "S":
                    self.start_coordinate = Coordinate(row, col)
                elif self.grid[row][col] == "E":
                    self.end_coordinate = Coordinate(row, col)
                elif (row not in (0, maxRow-1) or col not in (0, maxCol-1)) and self.grid[row][col] == "#":
                    self.obstacles.add(Coordinate(row, col))
        if self.start_coordinate is None:
            raise ValueError("Start coordinate not found")
        if self.end_coordinate is None:
            raise ValueError("End coordinate not found")


class Graph:
    def __init__(self):
        self.graph: dict[Coordinate, dict[Coordinate, int]] = {}
        self.start_cordinate: Coordinate
        self.end_coordinate: Coordinate

    def add_edge(self, n1: Coordinate, n2: Coordinate) -> None:
        if n1 not in self.graph:
            self.graph[n1] = []
        self.graph[n1].append(n2)

    def from_grid(self, grid: list[list[str]], obstacle_chr: str = "#") -> "Graph":
        self.graph = {}
        for row_num, row in enumerate(grid):
            for col_num, col in enumerate(row):
                if col == obstacle_chr:
                    continue
                curr_node = Coordinate(row_num, col_num, value=col)
                if curr_node.value == "S":
                    self.start_cordinate = curr_node
                elif curr_node.value == "E":
                    self.end_coordinate = curr_node
                row_lim = len(grid)
                col_lim = len(grid[0])
                for next_node in curr_node.next_nodes(row_lim, col_lim, grid, obstacle="#"):
                    if self.graph.get(curr_node) is None:
                        self.graph[curr_node] = {}
                    self.graph[curr_node][next_node] = 1

        return self

    def calculate_distances(self, start: Coordinate, end: Coordinate) -> int:
        visited = set()
        priority_queue = [(0, start)]
        heapify(priority_queue)
        distances: dict[Coordinate, int] = {coord: float("inf") for coord in self.graph.keys()}
        distances[start] = 0
        while priority_queue:
            distance, curr_node = heappop(priority_queue)
            if curr_node in visited:
                continue
            visited.add(curr_node)
            for node, displacement in self.graph[curr_node].items():
                dist_to_node = distance + displacement
                if dist_to_node < distances[node]:
                    distances[node] = dist_to_node
                    heappush(priority_queue, (dist_to_node, node))
        return distances


def part01(data: Input):
    print(f"num-obstacles: {len(data.obstacles)}")
    graph = Graph().from_grid(data.grid)
    distances = graph.calculate_distances(graph.start_cordinate, graph.end_coordinate)
    original_dist = distances[graph.end_coordinate]
    start = time.time()
    saves = {}
    for i, obstacle in enumerate(data.obstacles):
        data.grid[obstacle.row][obstacle.col] = "."
        graph = Graph().from_grid(data.grid)
        data.grid[obstacle.row][obstacle.col] = "#"
        distances = graph.calculate_distances(graph.start_cordinate, graph.end_coordinate)
        cheat_dist = distances[graph.end_coordinate]
        delta = original_dist - cheat_dist
        if saves.get(delta) is None:
            saves[delta] = 1
        else:
            saves[delta] += 1
        print(
            f"{i+1}: obstacle: {obstacle}, distance: {distances[graph.end_coordinate]}"
        )
    hundos = 0
    for k, v in saves.items():
        if k >= 100:
            hundos += v
    print(f"100s: {hundos}")
    with open("saved.json", "w") as fp:
        json.dump(saves, fp, indent=2)
    print(f"time taken: {time.time() - start}")


def part02(data):
    ...


def main():
    data = Input("./day20/input.txt").process_aoc_input()
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()
