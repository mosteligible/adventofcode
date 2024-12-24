from typing import Literal


def process_input() -> dict[str, dict[str, Literal[True]]]:
    connections: dict[str, dict[str, Literal[True]]] = {}
    with open("./day23/input.txt", "r") as fp:
        content = fp.readlines()
    content = [i.strip().split("-") for i in content]
    for conn in content:
        if connections.get(conn[0]) is None:
            connections[conn[0]] = {conn[1]: True}
        else:
            connections[conn[0]][conn[1]] = True
    return content


def get_num_connections(data: dict[str, dict[str, Literal[True]]], conns: dict[str, Literal[True]]) -> int:
    total = 0
    for conn in conns.keys():
        total += len(data[conn])
    return total


def part01(data: dict[str, dict[str, Literal[True]]]) -> None:
    count = 0
    for first, first_conns in data.items():
        if first.startswith("t"):
            num_first_conn = len(first_conns)


def part02(data):
    ...


def main():
    data = process_input()
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()
