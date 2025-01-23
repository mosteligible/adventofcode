import re
from dataclasses import dataclass

SAMPLE = """Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279"""


X_MOVE_REGEX = re.compile(r"X\+(\d+),")
Y_MOVE_REGEX = re.compile(f"Y\+(\d+)")
X_PRIZE_MOVE = re.compile(r"X=(\d+),")
Y_PRIZE_MOVE = re.compile(r"Y=(\d+)")


@dataclass
class Button:
    x: int = -1
    y: int = -1


@dataclass
class Prize:
    x: int = -1
    y: int = -1


@dataclass
class ArcadeMachine:
    a: Button = Button()
    b: Button = Button()
    prize: Prize = Prize()


def process_input():
    data = []
    with open("./day13/input.txt", "r") as fp:
        content = fp.readlines()

    # content = SAMPLE.splitlines()
    an_arcade_machine = ArcadeMachine()
    for line in content:
        if line.startswith("Button"):
            button_name = line[7]
            x_move = int(X_MOVE_REGEX.search(line).group(1))
            y_move = int(Y_MOVE_REGEX.search(line).group(1))
            if button_name == "A":
                an_arcade_machine.a = Button(x=x_move, y=y_move)
            else:
                an_arcade_machine.b = Button(x=x_move, y=y_move)
        elif line.startswith("Prize:"):
            x_move = int(X_PRIZE_MOVE.search(line).group(1))
            y_move = int(Y_PRIZE_MOVE.search(line).group(1))
            an_arcade_machine.prize = Prize(x=x_move, y=y_move)
            data.append(an_arcade_machine)
            an_arcade_machine = ArcadeMachine()

    return data


def num_tokens(l1: Button, l2: Button, prize: Prize) -> int:
    divisible = (l2.y * prize.x - l1.y * prize.y) % (l2.y * l1.x - l1.y * l2.x)
    if divisible == 0:
        x_moves = (l2.y * prize.x - l1.y * prize.y) // (l2.y * l1.x - l1.y * l2.x)
    else:
        return 0
    y_moves = (prize.x * l2.x - prize.y * l1.x) // (l1.y * l2.x - l2.y * l1.x)
    tokens = x_moves * 3 + y_moves
    # print(f'{x_moves=}, {y_moves=}, {tokens=} | {prize.x=}')
    return tokens


def part01(data: list[ArcadeMachine]):
    tokens = 0
    for machine in data:
        l1 = Button(x=machine.a.x, y=machine.b.x)
        l2 = Button(x=machine.a.y, y=machine.b.y)
        t = num_tokens(l1, l2, machine.prize)
        tokens += t
        # print(f" [X] {machine.a=}, {machine.b=}, tokens: {t}")
    print()
    print(f"most prizes = {tokens}")


def part02(data: list[ArcadeMachine]):
    print(f" -----------------  PART 02  ----------------- ")
    tokens = 0
    for machine in data:
        machine.prize.x += 10000000000000
        machine.prize.y += 10000000000000
        l1 = Button(x=machine.a.x, y=machine.b.x)
        l2 = Button(x=machine.a.y, y=machine.b.y)
        t = num_tokens(l1, l2, machine.prize)
        tokens += t
        # print(f" [X] {machine=}\ntokens: {t}")
    print()
    print(f"most prizes = {tokens}")


def other():
    with open("./day13/input.txt") as f:
        scenarios = f.read().split("\n\n")

    def parse(scenario):
        output = {}
        a, b, prize = scenario.splitlines()
        output["A"] = [int(item.split("+")[1]) for item in a.split(":")[1].split(", ")]
        output["B"] = [int(item.split("+")[1]) for item in b.split(":")[1].split(", ")]
        output["Prize"] = [
            10000000000000 + int(item.split("=")[1])
            for item in prize.split(":")[1].split(", ")
        ]
        return output

    scenarios = [parse(scenario) for scenario in scenarios]

    def solve(scenario):
        ax, ay = scenario["A"]
        bx, by = scenario["B"]
        tx, ty = scenario["Prize"]
        b = (tx * ay - ty * ax) // (ay * bx - by * ax)
        a = (tx * by - ty * bx) // (by * ax - bx * ay)
        if ax * a + bx * b == tx and ay * a + by * b == ty:
            return 3 * a + b
        else:
            return 0

    answer = sum(solve(scenario) for scenario in scenarios)
    print(answer)


def main():
    data = process_input()
    part01(data)
    part02(data)
    other()


if __name__ == "__main__":
    main()
