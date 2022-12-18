package day11

func sampleMonkeys() []monkey {
	return []monkey{
		{
			items:       []int{79, 98},
			operation:   func(x int) int { return x * 19 },
			testArg:     23,
			trueGoToID:  2,
			falseGoToID: 3,
		},
		{
			items:       []int{54, 65, 75, 74},
			operation:   func(x int) int { return x + 6 },
			testArg:     19,
			trueGoToID:  2,
			falseGoToID: 0,
		},
		{
			items:       []int{79, 60, 97},
			operation:   func(x int) int { return x * x },
			testArg:     13,
			trueGoToID:  1,
			falseGoToID: 3,
		},
		{
			items:       []int{74},
			operation:   func(x int) int { return x + 3 },
			testArg:     17,
			trueGoToID:  0,
			falseGoToID: 1,
		},
	}
}

func realInput() []monkey {
	return []monkey{
		{
			items: []int{62, 92, 50, 63, 62, 93, 73, 50},
			operation: func(x int) int {
				return x * 7
			},
			testArg:     2,
			trueGoToID:  7,
			falseGoToID: 1,
		},
		{
			items: []int{51, 97, 74, 84, 99},
			operation: func(x int) int {
				return x + 3
			},
			testArg:     7,
			trueGoToID:  2,
			falseGoToID: 4,
		},
		{
			items: []int{98, 86, 62, 76, 51, 81, 95},
			operation: func(x int) int {
				return x + 4
			},
			testArg:     13,
			trueGoToID:  5,
			falseGoToID: 4,
		},
		{
			items: []int{53, 95, 50, 85, 83, 72},
			operation: func(x int) int {
				return x + 5
			},
			testArg:     19,
			trueGoToID:  6,
			falseGoToID: 0,
		},
		{
			items: []int{59, 60, 63, 71},
			operation: func(x int) int {
				return x * 5
			},
			testArg:     11,
			trueGoToID:  5,
			falseGoToID: 3,
		},
		{
			items: []int{92, 65},
			operation: func(x int) int {
				return x * x
			},
			testArg:     5,
			trueGoToID:  6,
			falseGoToID: 3,
		},
		{
			items: []int{78},
			operation: func(x int) int {
				return x + 8
			},
			testArg:     3,
			trueGoToID:  0,
			falseGoToID: 7,
		},
		{
			items: []int{84, 93, 54},
			operation: func(x int) int {
				return x + 1
			},
			testArg:     17,
			trueGoToID:  2,
			falseGoToID: 1,
		},
	}

}
