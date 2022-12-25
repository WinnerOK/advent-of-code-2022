package day14

import (
	"github.com/WinnerOK/advent-of-code-2022/utils"
)

const extendSize = 170

func Part2(input []string) int {
	var rockPaths []path
	for _, rockPathStr := range input {
		rockPaths = append(rockPaths, parsePath(rockPathStr))
	}

	minPoint, maxPoint := getActiveBoundaries(rockPaths)
	maxPoint.Y += 2
	maxPoint.X += extendSize
	minPoint.X -= extendSize
	rockPaths = append(rockPaths, path{point{X: minPoint.X, Y: maxPoint.Y}, point{X: maxPoint.X, Y: maxPoint.Y}})

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
					panic("Simulation width is not enough. Please increase `extendSize` const")
				}
				if itemAtDir == Air {
					moved = true
					sandCoords = tryCoords
					break SandMove
				}
			}
			if !moved {
				_ = sim.Set(sandCoords, Sand)
				if sandCoords == sandSourceCoords {
					sandDropped += 1
					break SimulationLoop
				} else {
					break SandLoop
				}
			}
		}
		sandDropped += 1
	}

	return sandDropped
}
