from heapq import heapify, heappop, heappush


SAMPLE = """5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0"""


class Coordinate:
    def __init__(self, row: int, col: int, value: str):
        self.row = row
        self.col = col
        self.value = value

    def next_nodes(self, r_limit: int, c_limit: int, grid: list[list[str]]) -> list["Coordinate"]:
        retval = []
        # north
        if self.row - 1 >= 0:
            retval.append(grid[self.row - 1][self.col])
        # south
        if self.row + 1 < r_limit:
            retval.append(grid[self.row + 1][self.col])
        # west
        if self.col - 1 >= 0:
            retval.append(grid[self.row][self.col - 1])
        # east
        if self.col + 1 < c_limit:
            retval.append(grid[self.row][self.col + 1])
        return retval

    def __repr__(self) -> str:
        return f"(row:{self.row}, col:{self.col}, value:{self.value})"

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

    def from_aoc_input(self, grid: list[list[Coordinate]]) -> None:
        self.graph = {}
        row_limit = len(grid)
        col_limit = len(grid[0])
        for row in (grid):
            for col in (row):
                n1 = col
                if n1.value == "#":
                    continue
                elif n1.value == "S":
                    self.start_coordinate = n1
                elif n1.value == "E":
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
            current_distance, node = heappop(priority_queue)
            if node in visited_nodes:
                continue
            visited_nodes.add(node)
            for neighbor_node, _ in self.graph[node].items():
                tentative_distance = current_distance + 1
                if neighbor_node.value == "E":
                    print(f" [o] Reached E, current tentative dist: {tentative_distance}")
                if tentative_distance <= distances[neighbor_node]:
                    self.best_tiles.add(node)
                    distances[neighbor_node] = tentative_distance
                    heappush(priority_queue, (tentative_distance, neighbor_node))
        return distances


def show_grid(grid: list[list[Coordinate]]) -> None:
    print('-'*60)
    for row in grid:
        print("".join([i.value for i in row]))
    print('-'*60)


def get_grid(size: int) -> list[list[Coordinate]]:
    grid = []
    for row in range(size):
        grid.append([Coordinate(row, col, ".") for col in range(size)])
    grid[0][0].value = "S"
    grid[size - 1][size - 1].value = "E"
    return grid


def process_input() -> Graph:
    with open("./day18/input.txt", "r") as fp:
        content = fp.readlines()
    # content = SAMPLE.splitlines()
    coordinates = []
    for index, line in enumerate(content):
        line = line.strip()
        col, row = line.split(",")
        coordinates.append((int(row), int(col)))
    return coordinates


def part01(obstructions: list[tuple[int, int]], size: int = 71) -> None:
    grid = get_grid(size)
    for index in range(1024):
        row, col = obstructions[index]
        grid[row][col].value = "#"
    graph = Graph().from_aoc_input(grid)
    distances = graph.shortest_path()
    distance_to_end = distances[graph.end_coordinate]
    print(f"obs_{index=}, distance: {distance_to_end}")
    print("*" * 60)


def part02(obsreuctions: list[tuple[int, int]], size: int = 71) -> None:
    grid = get_grid(71)
    for index, coordinate in enumerate(obsreuctions):
        row, col = coordinate
        grid[row][col].value = "#"
        if index < 3000:
            continue
        graph = Graph().from_aoc_input(grid)
        distances = graph.shortest_path()
        distance_to_end = distances[graph.end_coordinate]
        if distance_to_end == float("inf"):
            print(" [x] No path to end")
            print(f"obs_{index=}, distance: {distance_to_end}")
            break


def main():
    obstructions = process_input()
    part01(obstructions)
    part02(obstructions)


if __name__ == "__main__":
    main()
