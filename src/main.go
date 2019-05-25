package main

import (
	"time"
)

func main() {
	// Set parameters.
	segSize := 2
	offset := 4
	rows := segSize*2 + 1

	PrintClock(NowAsMatrix(), segSize, offset)
	for {
		// Sleep for a small time just for make sure that nothing (Ctrl-C for
		// example) will mess up with screen
		time.Sleep(time.Millisecond * 10)
		MoveCursorPreviousLine(rows + 1)
		PrintClock(NowAsMatrix(), segSize, offset)
	}
}
