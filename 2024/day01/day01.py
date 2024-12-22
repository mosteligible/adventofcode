from collections import Counter


def process_input():
    with open("./day01/input.txt", "r") as fp:
        content = fp.readlines()
    left, right = [], []
    for line in content:
        l, r = line.split()
        left.append(int(l))
        right.append(int(r))
    left = sorted(left)
    right = sorted(right)
    return left, right


def part01(left, right) -> None:
    distance = 0
    for index, l in enumerate(left):
        distance += abs(l - right[index])
    print(f"total {distance=}")


def part02(left, right) -> None:
    similarity = 0
    right_count = Counter(right)

    for l in left:
        similarity += l * right_count.get(l, 0)
    print(f"total {similarity=}")


def main() -> None:
    left, right = process_input()
    part01(left, right)
    part02(left, right)


if __name__ == "__main__":
    main()
