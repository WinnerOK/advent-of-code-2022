package day13

import (
	"encoding/json"
)

func parseList(lst string) []interface{} {
	res := []interface{}{}
	err := json.Unmarshal([]byte(lst), &res)
	if err != nil {
		panic(err)
	}
	return res
}

func parseInput(input []string) (parsed [][2][]interface{}) {
	for i := 0; i <= len(input)/3; i++ {
		left := input[3*i]
		right := input[3*i+1]
		parsed = append(parsed, [2][]interface{}{
			parseList(left),
			parseList(right),
		})
	}
	return parsed
}

func Part1(input []string) int {
	data := parseInput(input)
	inOrderIndexSum := 0
	for idx, pair := range data {
		left, right := pair[0], pair[1]
		if isInOrder(left, right) == LOWER {
			inOrderIndexSum += idx + 1
		}
	}
	return inOrderIndexSum
}

type comparisonResult int

const (
	LOWER  comparisonResult = -1
	EQUAL  comparisonResult = 0
	HIGHER comparisonResult = 1
)

func isInOrder(left, right []interface{}) comparisonResult {
	for i := 0; i < len(left); i++ {
		if i > len(right)-1 {
			return HIGHER
		}

		leftNum, isLeftNum := left[i].(float64)
		rightNum, isRightNum := right[i].(float64)

		leftLst, isLeftLst := left[i].([]interface{})
		rightLst, isRightLst := right[i].([]interface{})

		if isLeftNum && isRightNum {
			if leftNum != rightNum {
				if leftNum < rightNum {
					return LOWER
				} else {
					return HIGHER
				}
			}
		} else if isLeftNum || isRightNum {
			if isLeftNum {
				leftLst = []interface{}{leftNum}
			} else if isRightNum {
				rightLst = []interface{}{rightNum}
			} else {
				panic("Should never happen")
			}
			cmp := isInOrder(leftLst, rightLst)
			if cmp != EQUAL {
				return cmp
			}
		} else if isLeftLst && isRightLst {
			cmp := isInOrder(leftLst, rightLst)
			if cmp != EQUAL {
				return cmp
			}
		} else {
			panic("Unexpected type combination")
		}
	}

	if len(left) < len(right) {
		return LOWER
	}

	return EQUAL
}
