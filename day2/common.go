package day2

type Shape int

const (
	rock = iota
	paper
	scissors
)

type Outcome int

const (
	WIN  = 6
	DRAW = 3
	LOSE = 0
)

func shapeScore(s Shape) int {
	return map[Shape]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}[s]
}

var opponentMap = map[uint8]Shape{
	'A': rock,
	'B': paper,
	'C': scissors,
}

func outcomeScore(opponent Shape, me Shape) int {
	if opponent == rock {
		if me == paper {
			return WIN
		}
		if me == scissors {
			return LOSE
		}
		return DRAW
	} else if opponent == paper {
		if me == rock {
			return LOSE
		}
		if me == scissors {
			return WIN
		}
		return DRAW
	} else { // opponent == scissors
		if me == rock {
			return WIN
		}
		if me == paper {
			return LOSE
		}
		return DRAW
	}
}
