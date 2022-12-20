package day12

import "fmt"

func parseInputPart2(input []string) ([][]int32, []coordinate, coordinate) {
	elevationMap := make([][]int32, len(input))
	for rowIdx := range elevationMap {
		elevationMap[rowIdx] = make([]int32, len(input[0]))
	}

	starts := []coordinate{}
	var end coordinate
	for rowIdx, row := range input {
		for colIdx, el := range row {
			var elevationNum int32
			switch el {
			case startMarker, lowestElevation:
				start := coordinate{int32(rowIdx), int32(colIdx)}
				starts = append(starts, start)
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
	return elevationMap, starts, end
}

func Part2(input []string) int {
	elevationMap, starts, end := parseInputPart2(input)
	minDist := len(input)*len(input[0]) + 1
	var minStart coordinate
	for _, start := range starts {
		pathLen, err := FindShortestPath(elevationMap, start, end)
		if err == nil {
			if pathLen < minDist {
				minStart = start
				minDist = pathLen
			}
		}
	}
	fmt.Printf("%v -> %v = ", minStart, end)
	return minDist
}
