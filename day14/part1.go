package day14

import (
	"fmt"
	"github.com/WinnerOK/advent-of-code-2022/utils"
	"math"
	"strconv"
	"strings"
)

type point = utils.IntVector

func pointFromStr(str string) point {
	nums := strings.Split(str, ",")
	if len(nums) > 2 {
		panic(fmt.Sprintf("Could not parse a point from \"%s\"", str))
	}
	x, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	return point{X: x, Y: y}
}

type path []point

func expandPath(p path) (newPath path) {
	newPath = append(newPath, p[0])
	for i := 0; i <= len(p)-2; i++ {
		from := p[i]
		to := p[i+1]

		dir := to.Sub(from).Normalize()
		for point := from.Add(dir); point != to; point = point.Add(dir) {
			newPath = append(newPath, point)
		}

		newPath = append(newPath, to)
	}
	return newPath
}

func parsePath(str string) (newPath path) {
	points := strings.Split(str, " -> ")
	for _, pointStr := range points {
		newPath = append(newPath, pointFromStr(pointStr))
	}
	return newPath
}

func getActiveBoundaries(rockPaths []path) (point, point) {
	minX := math.MaxInt
	minY := math.MaxInt

	maxX := math.MinInt
	maxY := math.MinInt

	for _, rockpath := range append(rockPaths, path{sandSourceCoords}) {
		for _, p := range rockpath {
			if p.X < minX {
				minX = p.X
			} else if p.X > maxX {
				maxX = p.X
			}

			if p.Y < minY {
				minY = p.Y
			} else if p.Y > maxY {
				maxY = p.Y
			}
		}
	}

	minPoint := point{X: minX, Y: minY}
	maxPoint := point{X: maxX, Y: maxY}
	return minPoint, maxPoint
}

var (
	sandSourceCoords = point{X: 500, Y: 0}
	below            = utils.IntVector{
		X: 0,
		Y: 1,
	}
	downLeft = utils.IntVector{
		X: -1,
		Y: 1,
	}
	downRight = utils.IntVector{
		X: 1,
		Y: 1,
	}
)

const (
	Rock       = '#'
	Air        = '.'
	Sand       = 'o'
	SandSource = '+'
)

func Part1(input []string) int {
	var rockPaths []path
	for _, rockPathStr := range input {
		rockPaths = append(rockPaths, parsePath(rockPathStr))
	}

	minPoint, maxPoint := getActiveBoundaries(rockPaths)
	sim := utils.InitSimulation(minPoint, maxPoint, Air)
	_ = sim.Set(sandSourceCoords, SandSource)
	for _, rockPath := range rockPaths {
		for _, p := range expandPath(rockPath) {
			_ = sim.Set(p, Rock)
		}
	}

	sandDropped := 0
SimulationLoop:
	for {
		sandCoords := sandSourceCoords

	SandLoop:
		for {
			moved := false
		SandMove:
			for _, dir := range []point{below, downLeft, downRight} {
				tryCoords := sandCoords.Add(dir)
				itemAtDir, err := sim.Get(tryCoords)
				if err != nil {
					break SimulationLoop
				}
				if itemAtDir == Air {
					moved = true
					sandCoords = tryCoords
					break SandMove
				}
			}
			if !moved {
				_ = sim.Set(sandCoords, Sand)
				break SandLoop
			}
		}
		sandDropped += 1
	}

	sim.Visualize(func(item int32) {
		fmt.Printf("%c", item)
	})

	return sandDropped
}
