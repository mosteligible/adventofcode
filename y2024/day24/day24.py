import time
from collections import deque
from enum import Enum, auto
from pprint import pprint


SAMPLE_SIGNALS = """x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1"""
SAMPLE_CONNS = """ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj"""


SIGNAL_MAP: dict[str, int] = {}

class Operator(Enum):
    AND = auto()
    OR = auto()
    XOR = auto()


class Connection:
    def __init__(self, sig1: str, sig2: str, gate: Operator, result: str):
        self.sig1: str = sig1
        self.sig2: str = sig2
        self.gate: Operator = gate
        self.result: str = result

    def operate(self, signals: dict[str, int]) -> int | None:
        sig1 = signals.get(self.sig1)
        sig2 = signals.get(self.sig2)
        if sig1 is None or sig2 is None:
            return None
        if self.gate == Operator.AND:
            return sig1 & sig2
        elif self.gate == Operator.OR:
            return sig1 | sig2
        elif self.gate == Operator.XOR:
            return sig1 ^ sig2
        raise ValueError(f"Invalid gate operator: {self.gate}")

    def __repr__(self):
        return f"{self.sig1} {self.gate.name} {self.sig2} -> {self.result}"


def set_signals(signals: dict[str, int]) -> None:
    global SIGNAL_MAP
    SIGNAL_MAP = signals


def process_input():
    signals: dict[str, int] = {}
    connections: list = []
    with open("./day24/input_signals.txt", "r") as fp:
        content = fp.readlines()
    # content = SAMPLE_SIGNALS.splitlines()
    for sig in content:
        indicator, num = sig.strip().split(": ")
        signals[indicator] = int(num)
    with open("./day24/input_wires.txt", "r") as fp:
        content = fp.readlines()
    # content = SAMPLE_CONNS.splitlines()
    for conn in content:
        wires, result = conn.strip().split(" -> ")
        sig1, gate, sig2 = wires.split()
        if gate == "AND":
            gate = Operator.AND
        elif gate == "OR":
            gate = Operator.OR
        elif gate == "XOR":
            gate = Operator.XOR
        else:
            raise ValueError(f"Invalid gate operator: {gate}")
        connections.append(Connection(sig1, sig2, gate, result))
    return signals, connections


def filter_zs(signals: dict[str, int]) -> str:
    start = 99
    retval = ""
    while start >= 0:
        key = f"z{start:02}"
        signal = signals.get(key)
        if signal is not None:
            retval = f"{retval}{signal}"
        start -= 1
    return retval


def part01(connections: list[Connection]):
    global SIGNAL_MAP
    conns: deque[Connection] = deque(connections)
    while conns:
        curr_conn = conns.popleft()
        res = curr_conn.operate(SIGNAL_MAP)
        if res is None:
            conns.append(curr_conn)
        else:
            SIGNAL_MAP[curr_conn.result] = res

    # pprint(SIGNAL_MAP)
    bin_rep = filter_zs(signals=SIGNAL_MAP)
    print(f"{bin_rep=}")
    print(f"{int(bin_rep, 2)=}")


def get_conn_from_result(connections: list[Connection], result: str) -> Connection:
    ...


def swap_results(
    connections: list[Connection], sig1: str, sig2: str, res: str,
) -> tuple[Connection, Connection, list[str]]:
    left_match = None
    result_match = None
    for conn in connections:
        if conn.sig1 == sig1 and conn.sig2 == sig2:
            left_match = conn
        if conn.result == res:
            result_match = conn
    left_match.result, result_match.result = result_match.result, left_match.result

    return left_match, result_match


def sort_connections(connections: list[Connection]) -> list[Connection]:
    start = 99
    while start >= 0:
        sig1 = f"x{start:02}"
        sig2 = f"y{start:02}"
        res = f"z{start:02}"
        l, r = swap_results(connections, sig1, sig2, res)
        start -= 1


def part02(signals: dict[str, int], connections: list[Connection]):
    ...


def main():
    signals, connections = process_input()
    set_signals(signals.copy())
    part01(connections)
    part02(signals, connections)


if __name__ == "__main__":
    main()
