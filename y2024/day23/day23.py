from typing import Literal

SAMPLE = """kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn"""


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
    return connections


def get_num_connections(
    data: dict[str, dict[str, Literal[True]]], conns: dict[str, Literal[True]]
) -> int:
    total = 0
    for conn in conns.keys():
        total += len(data[conn])
    return total


def num_connections(
    conns: dict[str, dict[str, Literal[True]]], curr_level: int, target: int
) -> int:
    if curr_level < target:
        ...


def part01(data: dict[str, dict[str, Literal[True]]]) -> None:
    count = 0
    for first, first_conns in data.items():
        num_first_conn = 0
        if first.startswith("t"):
            num_first_conn = len(first_conns)
        if num_first_conn == 0:
            continue
        num_second_conn = 0
        for second in first_conns:
            second_conns = data.get(second)


def part02(data):
    ...


def main():
    data = process_input()
    print(f"num connections: {len(data)}")
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()
