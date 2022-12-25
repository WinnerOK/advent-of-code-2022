package day15

import (
	"testing"
)

type baseCoverData struct {
	sensor, beacon point
	targetRow      int
}

type testInCoverData struct {
	input    baseCoverData
	expected interval
}

func TestSensorCovers(t *testing.T) {
	testInput := []testInCoverData{
		{
			input: baseCoverData{
				sensor:    point{X: 8, Y: 7},
				beacon:    point{X: 2, Y: 10},
				targetRow: 10,
			},

			expected: interval{start: 2, end: 14},
		},
		{
			input: baseCoverData{
				sensor:    point{X: 8, Y: 7},
				beacon:    point{X: 2, Y: 10},
				targetRow: 7,
			},
			expected: interval{start: -1, end: 17},
		},
		{
			input: baseCoverData{
				sensor:    point{X: 8, Y: 7},
				beacon:    point{X: 2, Y: 10},
				targetRow: 1,
			},
			expected: interval{start: 5, end: 11},
		},
	}

	for _, test := range testInput {
		actualInterval, err := sensorCoverageAtRow(test.input.sensor, test.input.beacon, test.input.targetRow)
		if err != nil {
			t.Errorf("failed with %s, but expected to cover", err.Error())
		}

		if actualInterval != test.expected {
			t.Errorf("Expected: %v\nGot: %v", test.expected, actualInterval)
		}
	}
}

func TestMergeSortedIntervals(t *testing.T) {

}
