from dataclasses import dataclass
from enum import Enum


class Direction(Enum):
    up = "^"
    down = "v"
    left = "<"
    right = ">"


@dataclass
class Coordinate:
    row: int
    col: int

    def move(self, direction: Direction) -> "Coordinate":
        ...


class Warehouse:
    def __init__(self, layout: list[list[str]]):
        self.layout: list[list[str]] = layout
        self.robot_pos: Coordinate = None

    def get_robot_pos(self) -> Coordinate:
        if self.robot_pos is not None:
            return self.robot_pos
        for row_num, row in self.layout:
            try:
                robot_col = row.index("@")
                self.robot_pos = Coordinate(row=row_num, col=robot_col)
                return self.robot_pos
            except IndexError:
                continue
        raise ValueError("robot not found!")

    def move_robot(self, direction: str) -> None:
        ...


def process_input() -> tuple[Warehouse, list[Direction]]:
    # warehouse
    with open("./day15/input_warehouse.txt", "r") as fp:
        warehouse = fp.readlines()

    warehouse = [list(i.strip()) for i in warehouse]
    warehouse = Warehouse(layout=warehouse)

    # movements
    with open("./day15/input_robot_movements.txt", "r") as fp:
        movements = fp.readlines()
    movements = [i.strip() for i in movements]
    directions = []
    for i in "".join(movements):
        if i == "^":
            directions.append(Direction.up)
        elif i == "v":
            directions.append(Direction.down)
        elif i == "<":
            directions.append(Direction.left)
        else:
            directions.append(Direction.right)

    return warehouse, directions


def part01(warehouse: Warehouse, movements: str):
    ...


def part02(warehouse: Warehouse, movements: str):
    ...


def main():
    warehouse, movements = process_input()
    part01(warehouse, movements)
    part02(warehouse, movements)


if __name__ == "__main__":
    main()
