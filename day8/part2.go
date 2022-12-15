package day8

import "strconv"

func Part2(input []string) int {
	ylen := len(input)
	xlen := len(input[0])
	heightMap := make([][]int, ylen)
	scenicScores := make([][]int, ylen)

	for y := 0; y < len(input); y++ {
		heightMap[y] = make([]int, xlen)
		scenicScores[y] = make([]int, xlen)
		for x, heightChar := range input[y] {
			height, err := strconv.Atoi(string(heightChar))
			if err != nil {
				panic(err)
			}
			heightMap[y][x] = height
			if x == 0 || y == 0 || x == xlen-1 || y == ylen-1 {
				scenicScores[y][x] = 0
			} else {
				scenicScores[y][x] = 1
			}
		}
	}

	// -------
	maxAnswer := 0
	// for a forest of size n*n
	// for each tree we at most look at all row and column
	// ie for each tree we look at n+n trees
	// total complexity: n*n*(n+n) = O(n^3)
	for y := 1; y < ylen-1; y += 1 {
		for x := 1; x < xlen-1; x += 1 {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx != 0 && dy != 0 ||
						dx == 0 && dy == 0 {
						continue
					}
					visibleTrees := 1
					currentHeight := heightMap[y][x]
					nextX := x + dx
					nextY := y + dy
					for 0 < nextX && nextX < xlen-1 &&
						0 < nextY && nextY < ylen-1 &&
						heightMap[nextY][nextX] < currentHeight {
						visibleTrees += 1
						nextX += dx
						nextY += dy
					}

					scenicScores[y][x] *= visibleTrees
				}
			}
			if scenicScores[y][x] > maxAnswer {
				maxAnswer = scenicScores[y][x]
			}
		}
	}
	return maxAnswer
}
