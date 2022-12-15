package day9

import "testing"

type stepTest struct {
	head, tail, dir, headExpected, tailExpected intVector
}

var tests = []stepTest{
	{
		head:         intVector{0, 0},
		tail:         intVector{0, 0},
		dir:          intVector{1, 0},
		headExpected: intVector{1, 0},
		tailExpected: intVector{0, 0},
	},
	{
		head:         intVector{0, 1},
		tail:         intVector{0, 0},
		dir:          intVector{0, -1},
		headExpected: intVector{0, 0},
		tailExpected: intVector{0, 0},
	},
	{
		head:         intVector{4, 4},
		tail:         intVector{4, 3},
		dir:          intVector{-1, 0},
		headExpected: intVector{3, 4},
		tailExpected: intVector{4, 3},
	},
	{
		head:         intVector{3, 4},
		tail:         intVector{4, 3},
		dir:          intVector{-1, 0},
		headExpected: intVector{2, 4},
		tailExpected: intVector{3, 4},
	},
	{
		head:         intVector{-4, -3},
		tail:         intVector{-5, -2},
		dir:          intVector{1, 0},
		headExpected: intVector{-3, -3},
		tailExpected: intVector{-4, -3},
	},
}

func TestSameRow(t *testing.T) {
	for _, test := range tests {
		if headOut, tailOut := simulate(test.head, test.tail, test.dir); headOut != test.headExpected || tailOut != test.tailExpected {
			t.Errorf("Expected head %+v, tail %+v\nGot head %+v, tail %+v", test.headExpected, test.tailExpected, headOut, tailOut)
		}
	}
}
