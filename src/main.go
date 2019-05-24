package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Set parameters.
	segSize := 2
	offset := 4
	rows := segSize*2 + 1

	// Although Flush method handles Ctrl-C interrupt we still need to handle
	// it for the case when cursor was moved up in the loop below. We also need
	// to guarantee that cursor will be showed again.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	moved := make(chan bool, 1)

	go func() {
		// Receive value from channel.
		<-sig

		// Receive value saying that cursor move function was called.
		if <-moved {
			MoveCursorNextLine(rows + 1)
		}

		// Show cursor again
		ShowCursor()

		// Now program can be stopped.
		os.Exit(0)
	}()

	for {
		// Hide cursor to make a pretty print.
		HideCursor()

		PrintClock(NowAsMatrix(), segSize, offset)
		MoveCursorPreviousLine(rows + 1)
		moved <- true

		ShowCursor()
	}
}
