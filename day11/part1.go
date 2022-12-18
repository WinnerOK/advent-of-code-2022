package day11

import (
	"sort"
)

// Monkey represents a monkey in the simulation.
type monkey struct {
	items                            []int
	operation                        func(int) int
	testArg, trueGoToID, falseGoToID int
}

func solve(inputMonkeys []monkey, stopRound int, isPart2 bool) int {
	inspections := make([]int, len(inputMonkeys))
	modulo := 1
	for _, curMonkey := range inputMonkeys {
		modulo *= curMonkey.testArg
	}

	for roundsPassed := 0; roundsPassed < stopRound; roundsPassed++ {
		for monkeyIdx, curMonkey := range inputMonkeys {
			inspections[monkeyIdx] += len(curMonkey.items)
			for _, itemStress := range curMonkey.items {
				newStress := curMonkey.operation(itemStress)
				if isPart2 {
					newStress %= modulo
				} else {
					newStress /= 3
				}

				var nextMonkeyIdx int
				if newStress%curMonkey.testArg == 0 {
					nextMonkeyIdx = curMonkey.trueGoToID
				} else {
					nextMonkeyIdx = curMonkey.falseGoToID
				}
				inputMonkeys[nextMonkeyIdx].items = append(inputMonkeys[nextMonkeyIdx].items, newStress)
			}
			inputMonkeys[monkeyIdx].items = []int{}
		}
	}

	sort.Slice(inspections, func(i, j int) bool { return inspections[i] > inspections[j] })

	return inspections[0] * inspections[1]
}

func Part1(_ []string) int {
	return solve(realInput(), 20, false)
}
