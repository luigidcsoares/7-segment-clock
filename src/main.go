package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Set parameters.
	segSize := 2
	offset := 4
	rows := segSize*2 + 1

	// Handle interrupt signal (Ctrl-C) properly.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		// Receive the OS interrupt signal.
		<-sig

		// Print exit message using the carriage return to erase ^C caused by
		// Ctrl-C signal.
		fmt.Print("\rSee ya ;)")

		os.Exit(0)
	}()

	PrintClock(NowAsMatrix(), segSize, offset)
	for {
		// Sleep for a small time just for make sure that nothing (Ctrl-C for
		// example) will mess up with screen
		time.Sleep(time.Millisecond * 10)
		MoveCursorPreviousLine(rows + 1)
		PrintClock(NowAsMatrix(), segSize, offset)
	}
}
