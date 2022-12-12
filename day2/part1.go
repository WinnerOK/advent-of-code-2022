package day2

var (
	meMap = map[uint8]Shape{
		'X': rock,
		'Y': paper,
		'Z': scissors,
	}
)

func Part1(input []string) int {
	result := 0
	for _, line := range input {
		opponent := opponentMap[line[0]]
		me := meMap[line[2]]
		result += outcomeScore(opponent, me) + shapeScore(me)
	}
	return result
}
