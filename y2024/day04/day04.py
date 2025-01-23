def process_input():
    with open("./day04/input.txt", "r") as fp:
        content = fp.readlines()
    print(f"{len(content)=}")
    #     content = """MMMSXXMASM
    # MSAMXMSMSA
    # AMXSXMAAMM
    # MSAMASMSMX
    # XMASAMXAMM
    # XXAMMXXAMA
    # SMSMSASXSS
    # SAXAMASAAA
    # MAMMMXMMMM
    # MXMXAXMASX"""
    #     content = content.splitlines()
    content = [list(line) for line in content]
    return content


def get_words(
    characters: list[list[str]], origin: tuple[int, int], size: tuple[int, int]
):
    # from given origin
    # can xmas be up
    x, y = origin
    x_limit, y_limit = size
    word_down, word_up, word_right, word_left = "", "", "", ""
    left_up, left_down, right_up, right_down = "", "", "", ""
    if x - 3 >= 0:
        word_up = (
            characters[x][y]
            + characters[x - 1][y]
            + characters[x - 2][y]
            + characters[x - 3][y]
        )
    # can xmas be down
    if x + 3 < x_limit:
        word_down = (
            characters[x][y]
            + characters[x + 1][y]
            + characters[x + 2][y]
            + characters[x + 3][y]
        )
    # can xmas be to left
    if y - 3 >= 0:
        word_left = "".join(characters[x][y - 3 : y + 1])[::-1]
    # can xmas be to right
    if y + 3 < y_limit:
        word_right = "".join(characters[x][y : y + 4])

    # diagonally, left up
    if x - 3 >= 0 and y - 3 >= 0:
        left_up = (
            characters[x][y]
            + characters[x - 1][y - 1]
            + characters[x - 2][y - 2]
            + characters[x - 3][y - 3]
        )
    # diagonally, left down
    if x + 3 < x_limit and y - 3 >= 0:
        left_down = (
            characters[x][y]
            + characters[x + 1][y - 1]
            + characters[x + 2][y - 2]
            + characters[x + 3][y - 3]
        )
    # diagonally, right up
    if x - 3 >= 0 and y + 3 < y_limit:
        right_up = (
            characters[x][y]
            + characters[x - 1][y + 1]
            + characters[x - 2][y + 2]
            + characters[x - 3][y + 3]
        )
    # diagonally, right down"
    if x + 3 < x_limit and y + 3 < y_limit:
        try:
            right_down = (
                characters[x][y]
                + characters[x + 1][y + 1]
                + characters[x + 2][y + 2]
                + characters[x + 3][y + 3]
            )
        except IndexError:
            print(f"{origin=}")
            print(f"{size=}")
            raise IndexError(f"{x+3=}, {y+3=}")

    return (
        word_left,
        word_right,
        word_up,
        word_down,
        left_up,
        left_down,
        right_up,
        right_down,
    )


def part01(data):
    rows = len(data)
    columns = len(data[0])
    print(f"xlimit: {rows}, ylimit: {columns}")
    found = 0
    for row in range(rows):
        for col in range(columns):
            if data[row][col] != "X":
                continue
            origin = (row, col)
            words = get_words(data, origin, (rows, columns))
            for w in words:
                if w == "XMAS":
                    found += 1
    print(f"{found=}")


def is_xmas(characters: list[list[str]], origin: tuple[int, int]):
    x, y = origin
    left_top_right_down = (
        characters[x - 1][y - 1] + characters[x][y] + characters[x + 1][y + 1]
    )
    right_top_left_bottom = (
        characters[x - 1][y + 1] + characters[x][y] + characters[x + 1][y - 1]
    )
    if (left_top_right_down == "MAS" or left_top_right_down[::-1] == "MAS") and (
        right_top_left_bottom == "MAS" or right_top_left_bottom[::-1] == "MAS"
    ):
        return True
    return False


def part02(data):
    print(f"{'*'*20}  part 02  {'*'*20}")
    rows = len(data)
    columns = len(data[0])
    found = 0
    for row in range(1, rows - 1):
        for col in range(1, columns - 1):
            if data[row][col] != "A":
                continue
            if is_xmas(data, (row, col)):
                found += 1
    print(f"{found=}")


def main():
    data = process_input()
    part01(data)
    part02(data)


if __name__ == "__main__":
    main()
