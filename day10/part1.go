package day10

import (
	"strconv"
)

const (
	noOp = iota
	addx
)

func parseInstruction(s string) (int, int) {
	if s == "noop" {
		return noOp, 0
	}
	arg, err := strconv.Atoi(s[5:])
	if err != nil {
		panic(err)
	}
	return addx, arg
}

func execute(input []string, beforeCycleHook, afterCycleHook func(int, int)) {
	register := 1
	clock := 1

	programCounter := 0
	executeAdd := false
	for programCounter < len(input) {
		beforeCycleHook(clock, register)

		op, arg := parseInstruction(input[programCounter])

		if executeAdd {
			register += arg
			programCounter += 1
			executeAdd = false
		} else {
			switch op {
			case noOp:
				programCounter += 1
				break
			case addx:
				executeAdd = true
				break
			}
		}

		clock += 1
		afterCycleHook(clock, register)
	}
}

func unitHook(_, _ int) {}

func Part1(input []string) int {
	interestingCycles := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}

	answer := 0
	execute(input, unitHook, func(clock int, register int) {
		if _, isInteresting := interestingCycles[clock]; isInteresting {
			answer += clock * register
		}
	})
	return answer
}
