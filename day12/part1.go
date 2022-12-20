package day12

import (
	"errors"
	"fmt"
)

const (
	lowestElevation  = 'a'
	highestElevation = 'z'
	startMarker      = 'S'
	endMarker        = 'E'
)

// FindShortestPath function is written by https://chat.openai.com/ and then optimized by me and adapted to solve the problem
// AI comment: FindShortestPath performs a BFS to find the length of the shortest path from source to destination
// on a 2D grid represented by the elevationMap slice
func FindShortestPath(elevationMap [][]int32, source, destination coordinate) (int, error) {
	// Create a queue to store the nodes to visit
	queue := []coordinate{source}

	// Create a map to store the distances from the source node
	distances := make(map[coordinate]int)
	distances[source] = 0

	// Create a set to store the visited nodes
	visited := make(map[coordinate]bool)

	// Perform the BFS
	for len(queue) > 0 {
		// Dequeue the next node
		current := queue[0]
		queue = queue[1:]

		// Check if the current node is the destination
		if current == destination {
			return distances[current], nil
		}

		// Mark the current node as visited
		visited[current] = true

		// Enqueue the unvisited neighbors of the current node
		// AI got this cycle totally wrong
		for dRow := int32(-1); dRow <= 1; dRow++ {
			for dCol := int32(-1); dCol <= 1; dCol++ {
				if (dRow == 0 && dCol == 0) ||
					(dRow != 0 && dCol != 0) {
					continue
				}
				neighborRow, neighborCol := current.row+dRow, current.col+dCol
				if neighborRow >= 0 && neighborRow < int32(len(elevationMap)) &&
					neighborCol >= 0 && neighborCol < int32(len(elevationMap[0])) {
					heightDiff := elevationMap[neighborRow][neighborCol] - elevationMap[current.row][current.col]
					if heightDiff <= 1 {
						neighbor := coordinate{row: neighborRow, col: neighborCol}
						if !visited[neighbor] {
							queue = append(queue, neighbor)
							distances[neighbor] = distances[current] + 1
							visited[neighbor] = true // this line speeds up the code a lot, because we early cut repeated neighbours
						}
					}
				}
			}
		}
	}

	return -1, errors.New("path does not exists")
}

type coordinate struct {
	row, col int32
}

func parseInputPart1(input []string) ([][]int32, coordinate, coordinate) {
	elevationMap := make([][]int32, len(input))
	for rowIdx := range elevationMap {
		elevationMap[rowIdx] = make([]int32, len(input[0]))
	}

	var start, end coordinate
	for rowIdx, row := range input {
		for colIdx, el := range row {
			var elevationNum int32
			switch el {
			case startMarker:
				start = coordinate{int32(rowIdx), int32(colIdx)}
				elevationNum = lowestElevation - lowestElevation
			case endMarker:
				end = coordinate{int32(rowIdx), int32(colIdx)}
				elevationNum = highestElevation - lowestElevation
			default:
				elevationNum = el - lowestElevation
			}
			elevationMap[rowIdx][colIdx] = elevationNum
		}
	}
	return elevationMap, start, end
}

func Part1(input []string) int {
	elevationMap, start, end := parseInputPart1(input)
	fmt.Printf("%v -> %v = ", start, end)
	pathLen, err := FindShortestPath(elevationMap, start, end)
	if err != nil {
		panic(err)
	}
	return pathLen
}
