package main

import (
	day "github.com/WinnerOK/advent-of-code-2022/day10"
)

func main() {
	input, err := readLines("./in.txt")
	if err != nil {
		panic(err)
	}
	println(day.Part1(input))
	println(day.Part2(input))
}
