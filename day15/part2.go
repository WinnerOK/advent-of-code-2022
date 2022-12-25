package day15

const DistressBoundary = 4000000

func tuningFrequency(p point) int {
	return p.X*4000000 + p.Y
}

func Part2(input []string) int {
	measurements, _ := parseInput(input)

	for sensor, beacon := range measurements {
		beaconDist := sensor.ManhattanDist(beacon)
		for dx := 0; dx <= beaconDist+1; dx++ {
			dy := beaconDist + 1 - dx

			for _, distVec := range []point{
				{-1, -1},
				{-1, 1},
				{1, -1},
				{1, 1},
			} {
				checkPoint := point{
					X: sensor.X + dx*distVec.X,
					Y: sensor.Y + dy*distVec.Y,
				}

				if !(0 <= checkPoint.X && checkPoint.X <= DistressBoundary &&
					0 <= checkPoint.Y && checkPoint.Y <= DistressBoundary) {
					continue
				}
				rowCoverageIntervals := mergedSensorCoverageAtRow(measurements, checkPoint.Y)
				if len(rowCoverageIntervals) == 2 { // 2 intervals with gap 1
					if rowCoverageIntervals[0].end+2 == rowCoverageIntervals[1].start {
						return tuningFrequency(point{X: rowCoverageIntervals[0].end + 1, Y: checkPoint.Y})

					}
				}
			}
		}
	}
	panic("Could not find any answer for Part 2")
}
