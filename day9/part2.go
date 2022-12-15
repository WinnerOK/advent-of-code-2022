package day9

const ropeSize = 10

func Part2(input []string) int {
	moves := make([]move, len(input))
	for i := 0; i < len(moves); i++ {
		moves[i] = fromString(input[i])
	}

	tailVisits := map[intVector]bool{}
	var rope [ropeSize]intVector
	tailVisits[rope[len(rope)-1]] = true

	for _, currentMove := range moves {
		for stepsLeft := currentMove.steps; stepsLeft > 0; stepsLeft-- {
			for tailIdx := 1; tailIdx < ropeSize; tailIdx++ {
				headPos := rope[tailIdx-1]
				tailPos := rope[tailIdx]

				moveDir := intVector{0, 0}
				if tailIdx == 1 {
					moveDir = currentMove.dir
				}
				rope[tailIdx-1], rope[tailIdx] = simulate(headPos, tailPos, moveDir)
			}
			tailVisits[rope[len(rope)-1]] = true
		}
	}

	return len(tailVisits)
}
