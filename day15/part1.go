package day15

import (
	"errors"
	"fmt"
	"github.com/WinnerOK/advent-of-code-2022/utils"
	"sort"
)

const TargetRow = 2000000

type point = utils.IntVector

func sensorCoverageAtRow(sensor, beacon point, row int) (i interval, err error) {
	radius := sensor.ManhattanDist(beacon)
	rowDist := utils.IntAbs(row - sensor.Y)

	if radius-rowDist < 0 {
		return interval{}, errors.New("sensor can't cover target row")
	}

	coverageStrength := radius - rowDist
	coverageCenterAtRow := sensor.X

	return interval{coverageCenterAtRow - coverageStrength, coverageCenterAtRow + coverageStrength}, nil
}

func mergedSensorCoverageAtRow(measurements map[point]point, row int) []interval {
	var intervals []interval
	for sensor, beacon := range measurements {
		rowInterval, err := sensorCoverageAtRow(sensor, beacon, row)
		if err == nil {
			intervals = append(intervals, rowInterval)
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return less(intervals[i], intervals[j])
	})

	return mergeSortedIntervals(intervals)
}

type interval struct {
	start, end int
}

func less(a, b interval) bool {
	if a.start != b.start {
		return a.start < b.start
	}
	return a.end < b.end
}

func areOverlap(first, second interval) bool {
	return first.start <= second.end && second.start <= first.end
}

func mergeSortedIntervals(intervals []interval) (res []interval) {
	currentInterval := intervals[0]

	for _, nextInterval := range intervals[1:] {
		if areOverlap(currentInterval, nextInterval) {
			maxEnd := nextInterval.end
			if maxEnd < currentInterval.end {
				maxEnd = currentInterval.end
			}
			currentInterval = interval{currentInterval.start, maxEnd}
		} else {
			res = append(res, currentInterval)
			currentInterval = nextInterval
		}
	}
	res = append(res, currentInterval)

	return res
}

func parseInput(input []string) (map[point]point, map[point]bool) {
	measurements := map[point]point{}
	beacons := map[point]bool{}
	for _, line := range input {
		var sensorX, sensorY, beaconX, beaconY int
		_, _ = fmt.Sscanf(
			line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensorX,
			&sensorY,
			&beaconX,
			&beaconY,
		)

		sensor := point{X: sensorX, Y: sensorY}
		beacon := point{X: beaconX, Y: beaconY}
		measurements[sensor] = beacon
		beacons[beacon] = true
	}

	return measurements, beacons
}

func Part1(input []string) int {
	measurements, beacons := parseInput(input)

	merged := mergedSensorCoverageAtRow(measurements, TargetRow)
	coverageAtRow := 0
	for _, currentInterval := range merged {
		coverageAtRow += currentInterval.end - currentInterval.start + 1
	}

	beaconsAtRow := 0
	for beacon := range beacons {
		if beacon.Y == TargetRow {
			beaconsAtRow++
		}
	}

	return coverageAtRow - beaconsAtRow
}
