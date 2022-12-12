package day3

import "unicode"

func Part1(input []string) int {
	cumulativePriority := 0
	for _, rucksack := range input {
		compartmentSize := len(rucksack) / 2
		firstCompartment := map[int32]bool{}
		for idx, item := range rucksack {
			if idx < compartmentSize {
				firstCompartment[item] = true
			} else {
				if ok, _ := firstCompartment[item]; ok {
					cumulativePriority += priority(item)
					break // don't fall for repetition
				}
			}
		}
	}
	return cumulativePriority
}

func priority(item int32) int {
	if unicode.IsUpper(item) {
		return 26 + priority(unicode.ToLower(item))
	} else {
		return int(item - 'a' + 1)
	}
}
