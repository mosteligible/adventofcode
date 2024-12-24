from pprint import pprint


SAMPLE = """190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20"""


def process_input():
    with open("./day07/input.txt", "r") as fp:
        content = fp.readlines()
    # content = SAMPLE.splitlines()
    equations = []
    for line in content:
        total, nums = line.split(": ")
        total = int(total)
        nums = [int(i) for i in nums.split()]
        equations.append((total, nums))
    return equations


def ten_power(num: int) -> int:
    ten = 10
    while num // ten != 0:
        ten *= 10
    return ten


def operate(nums: list[int], curr_index: int, res: int, target: int) -> bool:
    """
        77, 77, 16, 765, 9

            77
         /+     \*
       77         77
    /+     \*   /+   \*
  16        16 16      16

    """
    # print(f"\n{nums=}, {curr_index=}, {res=}, {target=}", end = " ")
    if curr_index >= len(nums) and res != target:
        return False
    if curr_index >= len(nums) and res == target:
        return True

    sum_res = nums[curr_index] + res
    # print("sum", end = " ")
    sum_ok = operate(nums, curr_index+1, res=sum_res, target=target)
    # print(f"{sum_ok=}")
    if sum_ok:
        return True
    prod_res = nums[curr_index] * res
    # print("product", end=" ")
    prod_ok = operate(nums, curr_index+1, res=prod_res, target=target)
    # print(f"{prod_ok=}")
    return prod_ok


def operate_extra(nums: list[int], curr_index: int, res: int, target: int) -> bool:
    """
        77, 77, 16, 765, 9

            77
         /+     \*
       77         77
    /+     \*   /+   \*
  16        16 16      16

    """
    # print(f"\n{nums=}, {curr_index=}, {res=}, {target=}", end = " ")
    if curr_index >= len(nums) and res != target:
        return False
    if curr_index >= len(nums) and res == target:
        return True

    sum_res = nums[curr_index] + res
    # print("sum", end = " ")
    sum_ok = operate_extra(nums, curr_index+1, res=sum_res, target=target)
    # print(f"{sum_ok=}")
    if sum_ok:
        return True
    pow_10 = ten_power(nums[curr_index])
    concatenate_res = res * pow_10 + nums[curr_index]
    concatenate_ok = operate_extra(nums, curr_index+1, res=concatenate_res, target=target)
    if concatenate_ok:
        return True
    prod_res = nums[curr_index] * res
    # print("product", end=" ")
    prod_ok = operate_extra(nums, curr_index+1, res=prod_res, target=target)
    # print(f"{prod_ok=}")
    if prod_ok:
        return True
    return False


def part01(data):
    print(f"{'X'*30}  PART 01  {'X'*30}")
    total = 0
    for target, nums in data:
        result = operate(nums[1:], curr_index=0, res=nums[0], target=target)
        if result:
            print(f"=> Success with, {target=}, {nums=}, {result=}")
            total += target
        else:
            print(f"XXXX FAILURE: {target=}, {nums=}, {result=}")
    print(f"{total=}")


def part02(data):
    print(f"{'X'*30}  PART 02  {'X'*30}")
    total = 0
    for target, nums in data:
        result = operate_extra(nums[1:], curr_index=0, res=nums[0], target=target)
        if result:
            print(f"=> Success with, {target=}, {nums=}, {result=}")
            total += target
        else:
            print(f"XXXX FAILURE: {target=}, {nums=}, {result=}")
    print(f"{total=}")


def main():
    data = process_input()
    # part01(data)
    part02(data)
    # to_concatenate = 12345
    # power = (ten_power(to_concatenate))
    # num = 78
    # print(f"{num=}, {power=} => {num * power + to_concatenate}")


if __name__ == "__main__":
    main()

