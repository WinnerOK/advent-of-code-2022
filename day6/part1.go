package day6

func solve(input []rune, markerSize int) int {
	window := make(map[rune]int, markerSize)

	var answer int
	for idx, char := range input {
		if len(window) == markerSize {
			answer = idx
			break
		}
		if pos, ok := window[char]; ok {
			// introduces x^2 time complexity
			for k, v := range window {
				if v < pos {
					delete(window, k)
				}
			}
		}

		window[char] = idx
	}

	return answer
}

func Part1(input []string) int {
	inputStr := []rune(input[0])
	return solve(inputStr, 4)
}
