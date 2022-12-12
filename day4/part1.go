package day4

import (
	"strconv"
	"strings"
)

func parseRange(rangeStr string) (int, int) {
	splitRange := strings.Split(rangeStr, "-")
	start, err := strconv.Atoi(splitRange[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(splitRange[1])
	if err != nil {
		panic(err)
	}

	return start, end
}

func isInside(start1 int, end1 int, start2 int, end2 int) bool {
	// is 2nd inside 1st
	return start1 <= start2 && end2 <= end1
}

func Part1(input []string) int {
	answer := 0
	for _, line := range input {
		ranges := strings.Split(line, ",")
		start1, end1 := parseRange(ranges[0])
		start2, end2 := parseRange(ranges[1])
		if isInside(start1, end1, start2, end2) ||
			isInside(start2, end2, start1, end1) {
			answer += 1
		}
	}
	return answer
}
