package day9

import "strconv"

type intVector struct {
	x, y int
}

func (v intVector) Add(other intVector) intVector {
	return intVector{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v intVector) Rev() intVector {
	return intVector{
		x: -v.x,
		y: -v.y,
	}
}

func (v intVector) Sub(other intVector) intVector {
	return v.Add(other.Rev())
}

func (v intVector) Normalize() intVector {
	newX := v.x
	if newX != 0 {
		newX = v.x / abs(v.x)
	}
	newY := v.y
	if newY != 0 {
		newY = v.y / abs(v.y)
	}
	return intVector{
		x: newX,
		y: newY,
	}
}

var directionMap = map[uint8]intVector{
	'R': {
		x: 1,
		y: 0,
	},
	'U': {
		x: 0,
		y: 1,
	},
	'D': {
		x: 0,
		y: -1,
	},
	'L': {
		x: -1,
		y: 0,
	},
}

type move struct {
	dir   intVector
	steps int
}

func fromString(s string) move {
	steps, err := strconv.Atoi(s[2:])
	if err != nil {
		panic(err)
	}

	return move{
		dir:   directionMap[s[0]],
		steps: steps,
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Part1(input []string) int {
	moves := make([]move, len(input))
	for i := 0; i < len(moves); i++ {
		moves[i] = fromString(input[i])
	}

	tailVisits := map[intVector]bool{}
	headPos := intVector{0, 0}
	tailPos := intVector{0, 0}
	tailVisits[tailPos] = true

	for _, currentMove := range moves {
		for stepsLeft := currentMove.steps; stepsLeft > 0; stepsLeft-- {
			headPos, tailPos = simulate(headPos, tailPos, currentMove.dir)
			tailVisits[tailPos] = true
		}
	}

	return len(tailVisits)
}

func simulate(headPos intVector, tailPos intVector, dir intVector) (intVector, intVector) {
	newHeadPos := headPos.Add(dir)
	diff := newHeadPos.Sub(tailPos)

	// is Touching horizontally, vertically, diagonally or overlap
	isTouching := abs(diff.x) == 1 && diff.y == 0 ||
		diff.x == 0 && abs(diff.y) == 1 ||
		abs(diff.x) == 1 && abs(diff.y) == 1 ||
		diff == intVector{0, 0}
	if isTouching {
		return newHeadPos, tailPos
	}

	newTailPos := tailPos.Add(diff.Normalize())
	return newHeadPos, newTailPos
}
