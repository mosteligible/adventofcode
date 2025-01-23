from graphlib import TopologicalSorter


def process_input():
    with open("./day05/input_rules.txt", "r") as fp:
        content = fp.readlines()
    # content = RULES.splitlines()
    # dictionary: key is number that should be printed before
    # value is dictionary of values that should come after
    rules = {}
    for line in content:
        line = line.strip()
        before, after = line.split("|")
        if rules.get(before, False):
            rules[before][after] = True
        else:
            rules[before] = {after: True}
    with open("./day05/input_ordering.txt", "r") as fp:
        ordering = fp.readlines()
    # ordering = ORDERS.splitlines()
    ordering = [line.strip().split(",") for line in ordering]
    return rules, ordering


def part01(rules: dict[str, dict[str, bool]], orders):
    print(f"{'*'*20}  part 01  {'*'*20}")
    valid_orders = []
    invalid_orders = []
    for order in orders:
        # keep track of page numbers that have come through
        seen_pages = []
        invalid_order = False
        for pg_no in order:
            # get the numbers that should come after current pg_no
            after_pg_no = rules.get(pg_no, {})
            match = [i for i in seen_pages if after_pg_no.get(i) is not None]
            seen_pages.append(pg_no)
            if match:
                invalid_order = True
                invalid_orders.append(order)
                break
        if not invalid_order:
            valid_orders.append(order)

    mid_sum = 0
    for order in valid_orders:
        mid = order[len(order) // 2]
        mid_sum += int(mid)
    print(f"{mid_sum=}")
    return invalid_orders


def part02(rules, order):
    invalid_orders = part01(rules, order)
    print(f"{'*'*20}  part 02  {'*'*20}")
    print(f"{invalid_orders=}")
    reordered = []
    rules = {k: set(v) for k, v in rules.items()}
    for order in invalid_orders:
        ruled_items = {}
        for pg_no in order:
            item = rules.get(pg_no, {})
            ruled_items[pg_no] = item
        ts = TopologicalSorter(ruled_items)
        sorted_items = tuple(ts.static_order())
        sorted_items = sorted_items[::-1]
        sorted_items = [i for i in sorted_items if ruled_items.get(i) is not None]
        reordered.append(sorted_items)

    # print(f"{reordered=}")
    mid_sum = 0
    for items in reordered:
        mid = items[len(items) // 2]
        mid_sum += int(mid)
    print(f"{mid_sum=}")


def main():
    rules, order = process_input()
    part02(rules, order)


RULES = """47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13"""

ORDERS = """75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47"""


if __name__ == "__main__":
    main()
