package day17

import "math"

var (
	A = 45483412
	B = 0
	C = 0

	PROGRAM = []int{2, 4, 1, 3, 7, 5, 0, 3, 4, 1, 1, 5, 5, 5, 3, 0}
)

func getComboOperand(op, a, b, c int) int {
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	case 7:
		return 7
	}
	return op
}

func processOpCode(code, a, b, c, index int) (int, int, int, int, int) {
	comboOperand := getComboOperand(code, a, b, c)
	output := -1

	switch code {
	case 0:
		a = a / int(math.Pow(2, float64(comboOperand)))
	case 1:
		b = b ^ code
	case 2:
		b = (comboOperand % b)
	case 3:
		if a != 0 {
			index = code
		}
	case 4:
		b = b ^ c
	case 5:
		output = comboOperand % 8
	case 6:
		b = a / int(math.Pow(2, float64(comboOperand)))
	case 7:
		c = a / int(math.Pow(2, float64(comboOperand)))
	}

	return a, b, c, index, output
}

func part01() {

}

func Run(dir string) {
	part01()
}
