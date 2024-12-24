from typing import List


SAMPLE = """7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"""

def process_input() -> str:
    with open("./day02/input.txt", "r") as fp:
        content = fp.readlines()
    
    content = SAMPLE.splitlines()
    for index, line in enumerate(content):
        levels = line.strip().split()
        levels = [int(i) for i in levels]
        content[index] = levels
    return content


def is_level_safe(levels: List[int]) -> bool:
    decreasing = False
    if levels[0] > levels[1]:
        decreasing = True
    else:
        decreasing = False
    index = 0
    for l in levels[1:]:
        if levels[index] == l:
            return False, index, index + 1
        if (decreasing and levels[index] < l) or (not decreasing and levels[index] > l):
            return False, index, index + 1
        if (decreasing and (levels[index] - l > 3)) or (not decreasing and (l - levels[index] > 3)):
            return False, index, index + 1
        index += 1
    return True, -1, -1


def part01(data) -> None:
    print(f"{'*'*20}  part01  {'*'*20}")
    unsafe = 0
    safe = 0
    for levels in data:
        is_safe, _, _ = is_level_safe(levels)
        if is_safe:
            safe += 1
            continue
        unsafe += 1

    print(f"{safe=}")
    print(f"{unsafe=}")


def part02(data) -> None:
    print(f"{'*'*20}  part02  {'*'*20}")
    unsafe = 0
    safe = 0
    for levels in data:
        is_safe, idx1, idx2 = is_level_safe(levels)
        if is_safe:
            safe += 1
            continue

        l1 = levels[:idx1] + levels[idx1+1:]
        l2 = levels[:idx2] + levels[idx2+1:]
        is_safe_1, _, _ = is_level_safe(l1)
        is_safe_2, _, _ = is_level_safe(l2)
        if is_safe_1 or is_safe_2:
            safe += 1
            continue
        print(f"{levels=}")
        unsafe += 1
    print(f"{safe=}")
    print(f"{unsafe=}")


def main() -> None:
    data = process_input()
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()