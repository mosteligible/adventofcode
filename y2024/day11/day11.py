import math
from copy import deepcopy


INPUT = "17639 47 3858 0 470624 9467423 5 188"


def process_input():
    with open("./day11/input.txt", "r") as fp:
        content = fp.read()
    content = content.strip()
    content = {int(i):1 for i in content.split()}
    if content.get(1) is None:
        content[1] = 0
    if content.get(0) is None:
        content[0] = 0
    return content


def get_num_digits(n: int) -> int:
    return math.floor(math.log10(n)+ 1)


def part01(data: dict[int, int], num_blinks: int):
    rocks = deepcopy(data)
    for i in range(num_blinks):
        new_rocks = {}
        for stone, count in rocks.items():
            if stone == 0:
                if new_rocks.get(1) is None:
                    new_rocks[1] = 0
                new_rocks[1] += count
            elif get_num_digits(stone)%2==0:
                num_digits = get_num_digits(stone)
                ten_pow = num_digits // 2
                left = stone // 10**ten_pow
                right = stone % 10**ten_pow
                if new_rocks.get(left) is None:
                    new_rocks[left] = 0
                if new_rocks.get(right) is None:
                    new_rocks[right] = 0
                new_rocks[left] += count
                new_rocks[right] += count
            else:
                if new_rocks.get(stone*2024) is None:
                    new_rocks[stone*2024] = 0
                new_rocks[stone * 2024] += count
        rocks = new_rocks
        print(f"blinking: {i} - num_rocks = {sum(rocks.values())}")
    print(f"num rocks: {sum(rocks.values())}")


def main():
    data = process_input()
    part01(data, 25)
    part01(data, 75)


if __name__ == "__main__":
    main()

