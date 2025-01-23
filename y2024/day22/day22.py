from collections import deque

REQUIRED_SEQ_LENGTH = 4
SAMPLE = """1
10
100
2024"""


def process_input() -> list[int]:
    with open("./day22/input.txt", "r") as fp:
        content = fp.readlines()
    # content = SAMPLE.splitlines()
    content = [int(i.strip()) for i in content]
    return content


def mix(secret: int, value: int) -> int:
    return secret ^ value


def prune(secret: int) -> int:
    return secret % 16777216


def evolve(secret: int) -> int:
    prod = secret * 64
    secret = mix(secret, prod)
    secret = prune(secret)
    div = secret // 32
    secret = mix(secret, div)
    secret = prune(secret)
    prod = secret * 2048
    return prune(mix(secret, prod))


def nth_evolved_secret(secret: int, n: int = 2000) -> int:
    for _ in range(n):
        secret = evolve(secret)
    return secret


def nth_evolve_max(secret: int, n: int = 2000) -> int:
    """
    first secret, skip
    """
    prev = 0
    max_diff = -float("inf")
    temp_hold_diff = deque(maxlen=4)
    temp_hold_prices = deque(
        [
            secret,
        ],
        maxlen=4,
    )
    final_four_seq_diff = []
    final_four_seq_prices = []
    max_price = -1
    for _ in range(n):
        secret = evolve(secret)
        to_compare = secret % 10
        diff = to_compare - prev
        # print(f"{secret=}, {to_compare=}, diff={diff}")
        temp_hold_diff.append(diff)
        temp_hold_prices.append(to_compare)
        if diff > max_diff and len(temp_hold_diff) == REQUIRED_SEQ_LENGTH:
            max_diff = diff
            final_four_seq_diff = list(temp_hold_diff)
            final_four_seq_prices = list(temp_hold_prices)
            max_price = to_compare
        prev = to_compare
    print(f"max_price: {max_price}")
    return final_four_seq_diff, final_four_seq_prices


def part01(data: list[int]):
    total = 0
    for secret in data:
        total += nth_evolved_secret(secret)
    print(f"total: {total}")


def part02(data: list[int]):
    print(nth_evolve_max(2024))


def main():
    data = process_input()
    # part01(data)
    part02(data)


if __name__ == "__main__":
    main()
