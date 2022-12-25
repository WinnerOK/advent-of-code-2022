package day14

import (
	"reflect"
	"testing"
)

func TestExpandPath(t *testing.T) {
	type testData struct {
		sourcePath   path
		expectedPath path
	}

	testInput := []testData{
		{
			sourcePath: parsePath("498,4 -> 498,6 -> 496,6"),
			expectedPath: parsePath(
				"498,4 -> 498,5 -> 498,6 -> " +
					"497,6 -> 496,6",
			),
		},
		{
			sourcePath: parsePath("498,6 -> 495,6"),
			expectedPath: parsePath(
				"498,6 -> 497,6 -> 496,6 -> 495,6",
			),
		},
	}

	for _, test := range testInput {
		expanded := expandPath(test.sourcePath)
		if !reflect.DeepEqual(expanded, test.expectedPath) {
			t.Errorf(
				"Incorrect expansion:\n\tSource: %v\n\tExpected: %v\n\tActual: %v",
				test.sourcePath,
				test.expectedPath,
				expanded,
			)
		}
	}
}
