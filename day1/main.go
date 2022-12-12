package main

import (
	"sort"
	"strconv"
)

func main() {
	input, err := readLines("day1/in.txt")
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}

func part2(input []string) {
	calories := make([]int, len(input))
	curCalories := 0
	for _, calStr := range input {
		if len(calStr) > 0 {
			foodCalorie, err := strconv.Atoi(calStr)
			if err != nil {
				panic(err)
			}
			curCalories += foodCalorie
		} else {
			calories = append(calories, curCalories)
			curCalories = 0
		}
	}
	calories = append(calories, curCalories)

	sort.Ints(calories)
	sum := 0
	for _, v := range calories[len(calories)-3:] {
		sum += v
	}
	println(sum)
}

func part1(input []string) {
	maxCalories := 0
	curCalories := 0
	for _, calStr := range input {
		if len(calStr) == 0 {
			if curCalories > maxCalories {
				maxCalories = curCalories
			}
			curCalories = 0
		} else {
			foodCalorie, err := strconv.Atoi(calStr)
			if err != nil {
				panic(err)
			}
			curCalories += foodCalorie
		}
	}
	if curCalories > maxCalories {
		maxCalories = curCalories
	}
	curCalories = 0
	println(maxCalories)
}
