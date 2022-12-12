package day3

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func Part2(input []string) int {
	cumulativePriority := 0
	uniques := make([]mapset.Set, 3)
	for idx, rucksack := range input {
		groupIdx := idx % 3
		uniques[groupIdx] = mapset.NewSet()
		for _, item := range rucksack {
			uniques[groupIdx].Add(item)
		}

		if (idx+1)%3 == 0 {
			commonItemSet := uniques[0].Intersect(uniques[1]).Intersect(uniques[2])
			if commonLen := commonItemSet.Cardinality(); commonLen != 1 {
				panic(fmt.Sprintf("Found %d common items, expected 1", commonLen))
			}
			commonItem := commonItemSet.Pop().(rune)
			cumulativePriority += priority(commonItem)
		}

	}

	return cumulativePriority
}
