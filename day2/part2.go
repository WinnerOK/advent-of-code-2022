package day2

var outcomeMap = map[uint8]Outcome{
	'X': LOSE,
	'Y': DRAW,
	'Z': WIN,
}

func Part2(input []string) int {
	result := 0
	for _, line := range input {
		opponent := opponentMap[line[0]]
		desiredOutcome := outcomeMap[line[2]]
		me := calculateMe(opponent, desiredOutcome)
		result += outcomeScore(opponent, me) + shapeScore(me)
	}
	return result
}

func calculateMe(opponent Shape, outcome Outcome) Shape {
	switch outcome {
	case DRAW:
		return opponent
	case WIN:
		switch opponent {
		case rock:
			return paper
		case paper:
			return scissors
		default:
			//case scissors:
			return rock
		}
	default: // LOSE
		switch opponent {
		case rock:
			return scissors
		case paper:
			return rock
		default:
			//case scissors:
			return paper
		}
	}
}
