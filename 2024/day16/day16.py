from pprint import pprint
from enum import Enum
from heapq import heapify, heappop, heappush


SAMPLE = """###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############"""
# SAMPLE = """#################
# #...#...#...#..E#
# #.#.#.#.#.#.#.#.#
# #.#.#.#...#...#.#
# #.#.#.#.###.#.#.#
# #...#.#.#.....#.#
# #.#.#.#.#.#####.#
# #.#...#.#.#.....#
# #.#.#####.#.###.#
# #.#.#.......#...#
# #.#.###.#####.###
# #.#.#...#.....#.#
# #.#.#.#####.###.#
# #.#.#.........#.#
# #.#.#.#########.#
# #S#.............#
# #################"""

class Direction(Enum):
    north = 0
    east = 1
    south = 2
    west = 3
    not_set = 4


class Coordinate:
    def __init__(self, row: int, col: int, value: str, direction: Direction = Direction.not_set):
        self.row = row
        self.col = col
        self.value = value
        self.direction = direction

    def next_nodes(self, r_limit: int, c_limit: int, grid: list[list[str]]) -> list["Coordinate"]:
        retval = []
        # north
        if self.row - 1 >= 0:
            retval.append(Coordinate(self.row - 1, self.col, grid[self.row - 1][self.col]))
        # south
        if self.row + 1 < r_limit:
            retval.append(Coordinate(self.row + 1, self.col, grid[self.row + 1][self.col]))
        # west
        if self.col - 1 >= 0:
            retval.append(Coordinate(self.row, self.col - 1, grid[self.row][self.col - 1]))
        # east
        if self.col + 1 < c_limit:
            retval.append(Coordinate(self.row, self.col + 1, grid[self.row][self.col + 1]))
        return retval

    def get_score_to_neighbor(self, neighbor: "Coordinate") -> tuple[int, Direction]:
        # direction from current coordinate to neighbor
        if neighbor.row == self.row + 1:
            self_to_neighbor = Direction.south
        elif neighbor.row == self.row - 1:
            self_to_neighbor = Direction.north
        elif neighbor.col == self.col - 1:
            self_to_neighbor = Direction.west
        else:
            self_to_neighbor = Direction.east
        difference = abs(self.direction.value - self_to_neighbor.value)
        if difference == 0:
            return 1, self_to_neighbor
        elif difference == 1 or difference == 2:
            return 1000 * difference + 1, self_to_neighbor
        elif difference == 3:
            return 1000 + 1, self_to_neighbor

    def __repr__(self) -> str:
        return f"(row:{self.row}, col:{self.col}, value:{self.value}, dir:{self.direction})"

    def __hash__(self) -> int:
        return hash((self.row, self.col))

    def __eq__(self, value: "Coordinate"):
        if self.row == value.row and self.col == value.col:
            return True
        return False

    def __lt__(self, value: "Coordinate") -> bool:
        return (self.row, self.col) < (value.row, value.col)

    def __gt__(self, value: "Coordinate") -> bool:
        return (self.row, self.col) > (value.row, value.col)


class Graph:
    def __init__(self):
        self.graph : dict[Coordinate, dict[Coordinate, int]] = {}
        self.start_coordinate: Coordinate
        self.end_coordinate: Coordinate
        self.best_tiles: set = set()

    def add_edge(self, n1: Coordinate, n2: Coordinate, weight: int = 1):
        if self.graph.get(n1) is None:
            self.graph[n1] = {}
        self.graph[n1][n2] = weight

    def from_aoc_input(self, grid: list[list[str]]) -> None:
        row_limit = len(grid)
        col_limit = len(grid[0])
        for row_num, row in enumerate(grid):
            for col_num, col in enumerate(row):
                n1 = Coordinate(row_num, col_num, col)
                if col == "#":
                    continue
                elif col == "S":
                    self.start_coordinate = n1
                    self.start_coordinate.direction = Direction.east
                elif col == "E":
                    self.end_coordinate = n1
                surrounding_nodes = n1.next_nodes(row_limit, col_limit, grid)
                for i in surrounding_nodes:
                    if i.value != "#":
                        self.add_edge(n1, i)
        return self

    def shortest_path(self) -> dict[Coordinate, int]:
        if not self.graph:
            raise ValueError("graph has not been initialized!")
        visited_nodes: set = set()
        distances = {node: float("inf") for node in self.graph}
        distances[self.start_coordinate] = 0
        priority_queue = [(0, self.start_coordinate)]
        heapify(priority_queue)

        while priority_queue:
            # print("-"*72)
            current_distance, node = heappop(priority_queue)
            if node in visited_nodes:
                continue
            visited_nodes.add(node)
            for neighbor_node, _ in self.graph[node].items():
                distance, neighbor_direction = node.get_score_to_neighbor(neighbor=neighbor_node)
                neighbor_node.direction = neighbor_direction
                tentative_distance = current_distance + distance
                if neighbor_node.value == "E":
                    print(f" [o] Reached E, current tentative dist: {tentative_distance}")
                # print(f"- {node} -> {neighbor_node}  --  [{tentative_distance}]")
                if tentative_distance <= distances[neighbor_node]:
                    # node.direction = neighbor_direction
                    self.best_tiles.add(node)
                    distances[neighbor_node] = tentative_distance
                    heappush(priority_queue, (tentative_distance, neighbor_node))
        return distances


def process_input() -> Graph:
    with open("./day16/input.txt", "r") as fp:
        content = fp.readlines()
    content = SAMPLE.splitlines()
    content = [list(i.strip()) for i in content]
    # pprint(content)
    grid_graph = Graph().from_aoc_input(content)
    return grid_graph


def part01(graph: Graph):
    # pprint(graph.graph)
    distances = graph.shortest_path()
    # pprint(distances)
    pprint(graph.best_tiles)
    print(f"Shortest Distance to E{graph.end_coordinate}: {distances[graph.end_coordinate]}")
    print(f"shortest tile count: {len(graph.best_tiles)}")


def part02(data):
    ...


def main():
    data = process_input()
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()
