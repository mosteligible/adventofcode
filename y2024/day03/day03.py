import re

MUL_PATTERN = r"mul\((\d{1,3},\d{1,3})\)"
MUL_REGEX = re.compile(MUL_PATTERN)

COND_PATTERN = r"(do\(\))|(don't\(\))|mul\((\d{1,3},\d{1,3})\)"
COND_REGEX = re.compile(COND_PATTERN)


def process_input01():
    with open("./day03/input.txt", "r") as fp:
        content = fp.readlines()
    for index, line in enumerate(content):
        line_nums: list[str] = MUL_REGEX.findall(line)
        for i, nums in enumerate(line_nums):
            n1, n2 = nums.split(",")
            line_nums[i] = (int(n1), int(n2))
        content[index] = line_nums
    return content


def process_ipnut_01():
    with open("./day03/input.txt", "r") as fp:
        content = fp.readlines()
    for index, line in enumerate(content):
        line_nums = COND_REGEX.findall(line)
        for i, chunk in enumerate(line_nums):
            do, dont, nums = chunk
            if do:
                line_nums[i] = do
            elif dont:
                line_nums[i] = dont
            else:
                n1, n2 = nums.split(",")
                line_nums[i] = (int(n1), int(n2))
        content[index] = line_nums
    return content


def part_01():
    print(f"{'*'*20}  part_01  {'*'*20}")
    data = process_ipnut_01()
    total = 0
    for line in data:
        for d in line:
            if isinstance(d, tuple):
                total += d[0] * d[1]
    print(f"{total=}")


def part01():
    data = process_input01()
    print(f"{'*'*20}  part 01  {'*'*20}")
    total = 0
    for line in data:
        for nums in line:
            total += nums[0] * nums[1]
    print(f"{total=}")


def part02():
    print(f"{'*'*20}  part 02  {'*'*20}")
    data = process_ipnut_01()
    do = True
    total = 0
    for line in data:
        for d in line:
            if d == "do()":
                do = True
            elif d == "don't()":
                do = False
            elif do:
                total += d[0] * d[1]
    print(f"\n{total=}")


def main():
    part01()
    part_01()
    part02()


if __name__ == "__main__":
    main()
