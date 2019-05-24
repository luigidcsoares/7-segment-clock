package main

func main() {
	segSize := 2
	offset := 4
	rows := segSize*2 + 1

	for {
		// Hide cursor to make a pretty print.
		HideCursor()

		PrintClock(NowAsMatrix(), segSize, offset)
		MoveCursorPreviousLine(rows + 1)

		// Show cursor again.
		ShowCursor()
	}
}
