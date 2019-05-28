package main

import (
	"bufio"
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

func printRealTime(segment chan int, pause chan bool) {
	// Initial value.
	segSize := <-segment

	// Space between digits.
	margin := 4

	// Number of rows needed to print a digit.
	rows := segSize*2 + 1

	PrintClock(NowAsMatrix(), segSize, margin)
	for {
		select {
		case <-pause:
			segSize = <-segment
			rows = segSize*2 + 1
			PrintClock(NowAsMatrix(), segSize, margin)
		default:
			// Sleep for a small time just for make sure that nothing (Ctrl-C for
			// example) will mess up with screen
			time.Sleep(time.Millisecond * 10)
			MoveCursorPreviousLine(rows + 1)
			PrintClock(NowAsMatrix(), segSize, margin)
		}
	}
}

func askParams(min, max int) (n, segSize int) {
	// Number of printed lines.
	n = 0

	for {
		fmt.Print("## Entre com o tamanho do segmento [1..5]: ")
		n++

		fmt.Scanln(&segSize)
		fmt.Scan("\n")

		// Segment size is in the specified range.
		if segSize >= min && segSize <= max {
			break
		}

		// Else, ask the segment size again.
		fmt.Println("## O valor deve ser entre 1 e 5!!!")
		n++
	}

	return
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
	n, segSize := askParams(min, max)

	// If user pressed s, stop to read the new segment size.
	fmt.Println("## Pressione enter para redefinir os parâmetros")
	fmt.Println("## Pressione Ctrl-C para finalizar")

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

	// Create channels to communicate with print routine.
	pause := make(chan bool)
	segment := make(chan int)

	// Run real time printing on another go routine.
	go printRealTime(segment, pause)

	// Send initial segment size.
	segment <- segSize

	// New buffered read to read from stdin.
	// We're using this to read until reach a newline character, so we know
	// user pressed ENTER.
	reader := bufio.NewReader(os.Stdin)
	for {
		reader.ReadString('\n')

		// If enter was pressed, stop clock; clean everything and ask again for
		// input.
		pause <- true

		// Go back the total number of rows printed until then.
		// nrows is the number of rows between the first print and the curret
		// line of cursor when pressed ENTER.
		nrows := n + (segSize*2 + 1) + 4

		MoveCursorPreviousLine(nrows)
		ClearScreenEnd()
		Flush()

		n, segSize = askParams(min, max)

		// If user pressed s, stop to read the new segment size.
		fmt.Println("## Pressione enter para redefinir os parâmetros")
		fmt.Println("## Pressione Ctrl-C para finalizar")

		segment <- segSize
	}
}
