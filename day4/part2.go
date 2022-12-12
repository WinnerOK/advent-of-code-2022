package day4

import "strings"

func isOverlap(start1 int, end1 int, start2 int, end2 int) bool {
	return start1 <= end2 && start2 <= end1
}

func Part2(input []string) int {
	answer := 0
	for _, line := range input {
		ranges := strings.Split(line, ",")
		start1, end1 := parseRange(ranges[0])
		start2, end2 := parseRange(ranges[1])
		if isOverlap(start1, end1, start2, end2) {
			answer += 1
		}
	}
	return answer
}
