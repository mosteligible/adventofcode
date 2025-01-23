from collections import deque

SAMPLE = "2333133121414131402"
SPACE_ID = -1
SPACE_NOT_FOUND = -100
FILE_NOT_FOUND = -2
MOVED_FILE_ID = -5


def process_input() -> str:
    with open("./day09/input.txt", "r") as fp:
        content = fp.read()
    # content = SAMPLE
    content = [int(i) for i in content.strip()]
    return content


def calculate_checksum(compact_str: str) -> int:
    total = 0
    for index, ch in enumerate(compact_str):
        # if ch != ".":
        total += index * int(ch)
    return total


def get_right_index(dotted_list: list[int], right_index: int) -> int:
    right_index -= 1
    while True:
        if dotted_list[right_index] == -1:
            right_index -= 1
        else:
            return right_index


def part01(data: list[int]):
    dotted = []
    for idx, num in enumerate(data):
        if idx % 2 == 0:
            curr = [idx // 2] * num
        else:
            curr = [SPACE_ID] * num
        dotted.extend(curr)

    left = 0
    right = get_right_index(dotted_list=dotted, right_index=len(dotted))
    moved = []
    while left <= right:
        if dotted[left] != -1:
            moved.append(dotted[left])
        else:
            moved.append(dotted[right])
            right = get_right_index(dotted_list=dotted, right_index=right)
        left += 1
    print(f"PART 01: {sum([i*n for i, n in enumerate(moved)])}")


def get_dotted(data):
    dotted = []
    for idx, num in enumerate(data):
        if idx % 2 == 0:
            curr = [idx // 2] * num
        else:
            curr = [SPACE_ID] * num
        dotted.extend(curr)
    return dotted


def get_occurences(arr, idx):
    curr_num = arr[idx]
    occurences = 1
    while curr_num == arr[idx]:
        occurences += 1
        idx += 1
    return occurences - 1


def get_first_occuring_spaces(unfragmented: list[int]):
    try:
        first_space_index = unfragmented.index(-1)
    except ValueError:
        print(f"{SPACE_ID} not found!")
        return SPACE_NOT_FOUND, SPACE_NOT_FOUND
    num_space = get_occurences(unfragmented, first_space_index)
    return first_space_index, num_space


def get_matching_file_id(unfragmented: list[int], curr_idx: int, num_spaces: int):
    """
    unfragmented file list is reverse of original unfragmented list
    so we can operate as is
    """
    size = len(unfragmented)
    idx = 0
    while idx < size - curr_idx:
        file_id = unfragmented[idx]
        if file_id > 0:
            occurences = get_occurences(unfragmented, idx)
            if occurences <= num_spaces:
                return file_id, occurences, size - idx
            idx += occurences
            continue
        idx += 1
    return FILE_NOT_FOUND, FILE_NOT_FOUND, FILE_NOT_FOUND


def part02(data):
    dotted = get_dotted(data)

    while True:
        space_idx, num_spaces = get_first_occuring_spaces(dotted)
        # if space index is -100, solution is found
        if space_idx == -100:
            break
        right_file_id, occurences, index = get_matching_file_id(
            dotted[::-1], space_idx + num_spaces, num_spaces
        )
        if right_file_id == -2:
            dotted[space_idx : space_idx + num_spaces] = [FILE_NOT_FOUND] * num_spaces
        else:
            replacer = [right_file_id] * occurences
            dotted[space_idx : space_idx + occurences] = replacer
            dotted[index - occurences : index] = [MOVED_FILE_ID] * occurences
    checksum = 0
    for idx, num in enumerate(dotted):
        if num >= 0:
            checksum += idx * num
    print(f"checksum: {checksum}")


def main():
    data = process_input()
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()
