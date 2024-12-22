REGISTERS = {
    "A": 729,
    "B": 0,
    "C": 0,
}
PROG = [0, 1, 5, 4, 3, 0]

 
def process_input():
    with open("./day17/input.txt", "r") as fp:
        content = fp.readlines()
    reg_index = len("Register ")
    registers = {}
    for line in content:
        if line.startswith("Register"):
            registers[line[reg_index]] = int(line[reg_index+2:].strip())
        elif line.startswith("Program"):
            prog = line.replace("Program: ", "")
            prog = [int(i) for i in prog.strip().split(",")]
    # registers = REGISTERS
    # prog = PROG
    return registers, prog


def get_operand(instruction: int, registers: dict[str, int]) -> int:
    if 0 <= instruction <= 3:
        return instruction
    elif instruction == 4:
        return registers["A"]
    elif instruction == 5:
        return registers["B"]
    elif instruction == 6:
        return registers["C"]
    elif instruction == 7:
        raise ValueError("Invalid operand")


def part01(registers: dict[str, int], program: list[int]):
    output = []
    pointer = 0
    prog_size = len(program)
    while pointer < prog_size:
        opcode = program[pointer]
        operand = program[pointer + 1]
        match str(opcode):
            case "0":
                registers["A"] = registers["A"]//2**get_operand(operand, registers)
            case "1":
                registers["B"] = registers["B"] ^ opcode
            case "2":
                registers["B"] = 2 % 8
            case "3":
                if registers["A"] == 0:
                    continue
                print(f"-- {opcode=}, {operand=}, {registers=}")
                pointer += 2
                continue
            case "4":
                registers["B"] = registers["B"] ^ registers["C"]
            case "5":
                val = get_operand(operand, registers) % 8
                print(f"output: {val=} | {registers=}")
            case "6":
                registers["B"] = registers["A"]//2**get_operand(operand, registers)
            case "7":
                registers["C"] = registers["A"]//2**get_operand(operand, registers)
        print(f"-- {opcode=}, {operand=}, {registers=}")
        pointer += 2
    print()


def part02(data):
    ...


def main():
    data = process_input()
    print(data)
    part01(*data)
    part02(data)


if __name__ == "__main__":
    main()

