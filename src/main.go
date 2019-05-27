package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var mapColors = map[string]string{
	"black":   ColorBlack,
	"red":     ColorRed,
	"green":   ColorGreen,
	"yellow":  ColorYellow,
	"blue":    ColorBlue,
	"magenta": ColorMagenta,
	"cyan":    ColorCyan,
	"white":   ColorWhite,
}

func main() {
	// Get arguments from command line.
	// First position is the path to the program, so we're starting
	// from position 1.
	color := mapColors["magenta"]

	if len(os.Args) > 1 {
		color = mapColors[os.Args[1]]
	}

	// Changing foreground color by using ANSI colors.
	fmt.Print(color)

	// Set parameters.
	min := 1
	max := 5
	var segSize int

	for {
		fmt.Print("## Entre com o tamanho do segmento [1..5]: ")
		fmt.Scan(&segSize)
		fmt.Scan("\n")

		// Segment size is in the specified range.
		if segSize >= min && segSize <= max {
			fmt.Println()
			break
		}

		// Else, ask the segment size again.
		fmt.Println("## O valor deve ser entre 1 e 5!!!")
	}

	// Space between digits.
	margin := 4

	// Number of rows needed to print a digit.
	rows := segSize*2 + 1

	// Handle interrupt signal (Ctrl-C) properly.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		// Receive the OS interrupt signal.
		<-sig

		// Print exit message using the carriage return to erase ^C caused by
		// Ctrl-C signal.
		fmt.Println("\rSee ya ;)")

		os.Exit(0)
	}()

	PrintClock(NowAsMatrix(), segSize, margin)
	for {
		// Sleep for a small time just for make sure that nothing (Ctrl-C for
		// example) will mess up with screen
		time.Sleep(time.Millisecond * 10)
		MoveCursorPreviousLine(rows + 1)
		PrintClock(NowAsMatrix(), segSize, margin)
	}
}
