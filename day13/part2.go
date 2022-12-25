package day13

import "sort"

func Part2(input []string) int {
	div1 := parseList("[[2]]")
	div2 := parseList("[[6]]")
	data := [][]interface{}{
		div1, div2,
	}
	for _, line := range input {
		if len(line) > 0 {
			data = append(data, parseList(line))
		}
	}

	sort.Slice(data, func(i, j int) bool {
		return isInOrder(data[i], data[j]) == LOWER
	})

	answer := 1
	equalityCount := 0
	for idx, elem := range data {
		if isInOrder(elem, div1) == EQUAL || isInOrder(elem, div2) == EQUAL {
			answer *= idx + 1
			equalityCount++
		}
	}
	if equalityCount != 2 {
		panic("Found unexpected number of packets equal to dividers")
	}
	return answer
}
