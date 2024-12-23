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


def part01(data: list[int]):
    total = 0
    for secret in data:
        total += nth_evolved_secret(secret)
    print(f"total: {total}")


def part02(data):
    ...


def main():
    data = process_input()
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()
