package day5

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

type Move struct {
	count int
	from  int
	to    int
}

func parseMove(shiftStr string) Move {
	var (
		count int
		from  int
		to    int
	)
	_, err := fmt.Sscanf(shiftStr, "move %d from %d to %d", &count, &from, &to)
	if err != nil {
		panic(err)
	}

	return Move{
		count: count,
		from:  from - 1,
		to:    to - 1,
	}
}

func initStack(stackDescription []string) []*stack.Stack {
	stackNumberLine := stackDescription[len(stackDescription)-1]
	var stacksTotal int
	for i := len(stackNumberLine) - 1; i >= 0; i-- {
		char := stackNumberLine[i]
		if '0' <= char && char <= '9' {
			stacksTotal = int(char - '0')
			break
		}
	}

	stacks := make([]*stack.Stack, stacksTotal)
	for stackIdx := 0; stackIdx < stacksTotal; stackIdx++ {
		stacks[stackIdx] = stack.New()
		stackColumn := 4*stackIdx + 1
		for lineIdx := len(stackDescription) - 2; lineIdx >= 0; lineIdx-- {
			if len(stackDescription[lineIdx]) > stackColumn {
				stackItem := stackDescription[lineIdx][stackColumn]
				if stackItem != ' ' {
					stacks[stackIdx].Push(stackItem)
				}
			}
		}
	}

	return stacks
}

func parseInput(input []string) ([]*stack.Stack, []Move) {
	var stacks []*stack.Stack
	var moves []Move
	readMoves := false
	for idx, line := range input {
		if readMoves {
			moves = append(moves, parseMove(line))
		}

		if line == "" {
			stacks = initStack(input[0:idx])
			readMoves = true

		}
	}

	return stacks, moves
}

func Part1(input []string) string {
	stacks, moves := parseInput(input)
	for _, move := range moves {
		for i := 0; i < move.count; i++ {
			item := stacks[move.from].Pop()
			stacks[move.to].Push(item)
		}
	}
	var answer string
	for _, s := range stacks {
		top := s.Peek().(uint8)
		answer += string(top)
	}
	return answer
}
