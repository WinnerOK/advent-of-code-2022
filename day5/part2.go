package day5

func Part2(input []string) string {
	stacks, moves := parseInput(input)
	for _, move := range moves {
		var items []uint8
		for i := 0; i < move.count; i++ {
			items = append(items, stacks[move.from].Pop().(uint8))
		}
		for i := move.count - 1; i >= 0; i-- {
			stacks[move.to].Push(items[i])
		}

	}
	var answer string
	for _, s := range stacks {
		top := s.Peek().(uint8)
		answer += string(top)
	}
	return answer
}
