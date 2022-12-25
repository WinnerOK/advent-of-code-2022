package day13

import "testing"

type testData struct {
	left, right string
}

func TestInOfOrder(t *testing.T) {
	var inOrderData = []testData{
		{"[1, [], 0, 1]", "[1, [], 1, 1]"},
		{"[1,1,3,1,1]", "[1,1,5,1,1]"},
		{"[[1],[2,3,4]]", "[[1],4]"},
		{"[[4,4],4,4]", "[[4,4],4,4,4]"},
		{"[]", "[3]"},
	}

	for _, test := range inOrderData {
		left := parseList(test.left)
		right := parseList(test.right)

		if cmp := isInOrder(left, right); cmp != LOWER {
			t.Errorf("Failed %v, %v. Actual: %d", test.left, test.right, cmp)
		}
	}
}

func TestOutOfOrder(t *testing.T) {
	var outOfOrderData = []testData{
		{"[1, [], 5, 1]", "[1, [], 1, 1]"},
		{"[9]", "[[8,7,6]]"},
		{"[7,7,7,7]", "[7,7,7]"},
		{"[[[]]]", "[[]]"},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]"},
	}

	for _, test := range outOfOrderData {
		left := parseList(test.left)
		right := parseList(test.right)

		if cmp := isInOrder(left, right); cmp != HIGHER {
			t.Errorf("Failed %v, %v. Actual: %d", test.left, test.right, cmp)
		}
	}
}
