package day8

import (
	"fmt"
	"strconv"
)

const (
	top = iota
	bottom
	left
	right
)

func Part1(input []string) int {
	ylen := len(input)
	xlen := len(input[0])
	visibilityMap := make([][]bool, ylen)
	heightMap := make([][]int, ylen)
	maxHighFromDir := make([][][]int, 4)
	for idx, _ := range maxHighFromDir {
		maxHighFromDir[idx] = make([][]int, ylen)
		for y := 0; y < len(input); y++ {
			maxHighFromDir[idx][y] = make([]int, xlen)
		}
	}

	for y := 0; y < len(input); y++ {
		visibilityMap[y] = make([]bool, xlen)
		heightMap[y] = make([]int, xlen)
		for x, heightChar := range input[y] {
			height, err := strconv.Atoi(string(heightChar))
			if err != nil {
				panic(err)
			}
			heightMap[y][x] = height
			visibilityMap[y][x] = x == 0 || y == 0 || x == xlen-1 || y == ylen-1
			if y == 0 {
				maxHighFromDir[top][y][x] = heightMap[y][x]
			}
			if x == 0 {
				maxHighFromDir[left][y][x] = heightMap[y][x]
			}
			if y == ylen-1 {
				maxHighFromDir[bottom][y][x] = heightMap[y][x]
			}
			if x == xlen-1 {
				maxHighFromDir[right][y][x] = heightMap[y][x]
			}
		}
	}

	// -------

	for y := 1; y < ylen-1; y += 1 {
		for x := 1; x < xlen-1; x += 1 {
			if maxHighFromDir[top][y-1][x] < heightMap[y][x] {
				visibilityMap[y][x] = true
				maxHighFromDir[top][y][x] = heightMap[y][x]
			} else {
				maxHighFromDir[top][y][x] = maxHighFromDir[top][y-1][x]
			}
			// -------
			if maxHighFromDir[left][y][x-1] < heightMap[y][x] {
				visibilityMap[y][x] = true
				maxHighFromDir[left][y][x] = heightMap[y][x]
			} else {
				maxHighFromDir[left][y][x] = maxHighFromDir[left][y][x-1]
			}
		}
	}

	for y := ylen - 2; y > 0; y-- {
		for x := xlen - 2; x > 0; x-- {
			if maxHighFromDir[bottom][y+1][x] < heightMap[y][x] {
				visibilityMap[y][x] = true
				maxHighFromDir[bottom][y][x] = heightMap[y][x]
			} else {
				maxHighFromDir[bottom][y][x] = maxHighFromDir[bottom][y+1][x]
			}
			// ---------
			if maxHighFromDir[right][y][x+1] < heightMap[y][x] {
				visibilityMap[y][x] = true
				maxHighFromDir[right][y][x] = heightMap[y][x]
			} else {
				maxHighFromDir[right][y][x] = maxHighFromDir[right][y][x+1]
			}
		}
	}

	answer := 0
	for _, row := range visibilityMap {
		for _, isVisible := range row {
			if isVisible {
				answer += 1
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	return answer
}
