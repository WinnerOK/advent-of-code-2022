package day10

import "fmt"

const (
	lineWidth   = 40
	finishCycle = 240
)

func emptyScreen() [][]uint8 {
	screen := make([][]uint8, finishCycle/lineWidth)
	for colIdx, _ := range screen {
		screen[colIdx] = make([]uint8, lineWidth)
		for rowIdx, _ := range screen[colIdx] {
			screen[colIdx][rowIdx] = '.'
		}
	}
	return screen
}

func display(screen [][]uint8) {
	for _, row := range screen {
		for _, pixel := range row {
			fmt.Printf("%c", pixel)
		}
		println()
	}
}

func Part2(input []string) string {
	screen := emptyScreen()

	execute(input, func(clock int, register int) {
		drawingRow := (clock - 1) / lineWidth
		drawingCol := (clock - 1) % lineWidth

		if register-1 <= drawingCol && drawingCol <= register+1 {
			screen[drawingRow][drawingCol] = '#'
		}
	}, unitHook)

	display(screen)
	return ""
}
