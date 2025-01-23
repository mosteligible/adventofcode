package day23

import "adventofcode/y2024/utils"

const SAMPLE = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

func processInput(daydir string) map[string]map[string]bool {
	connectionInput := utils.ReadLines(daydir + "/input.txt")
	connections := map[string]map[string]bool{}
	for _, line := range connectionInput {
		comp1 := line[:2]
		comp2 := line[3:]
		if _, ok := connections[comp1]; !ok {
			connections[comp1] = map[string]bool{}
		}
		connections[comp1][comp2] = true
	}
	return connections
}

func part01(connections map[string]map[string]bool) {

}

func Run(dir string) {
	connections := processInput(dir + "/day23")
	part01(connections)
}
