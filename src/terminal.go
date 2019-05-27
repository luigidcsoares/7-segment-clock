package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	// ESC[ --> ESC = \033 (ASCII octal)
	esc = "\033"

	// ANSI Control Sequence Introducer used to move cursor on the screen:
	// ESC
	csi = esc + "["
)

// Basic 3/4 bit ANSI colors.
const (
	ColorBlack   = csi + "1;30m"
	ColorRed     = csi + "1;31m"
	ColorGreen   = csi + "1;32m"
	ColorYellow  = csi + "1;33m"
	ColorBlue    = csi + "1;34m"
	ColorMagenta = csi + "1;35m"
	ColorCyan    = csi + "1;36m"
	ColorWhite   = csi + "1;37m"
)

// A new buffered writer to write to stdout.
var out = bufio.NewWriter(os.Stdout)

// A buffer to be used before writing to stdout.
var screen = new(bytes.Buffer)

// Print writes params on screen buffer.
func Print(args ...interface{}) (n int, err error) {
	return fmt.Fprint(screen, args...)
}

// Println write params on screen buffer adding a new line.
func Println(args ...interface{}) (n int, err error) {
	return fmt.Fprintln(screen, args...)
}

// Printf format params and write on screen buffer.
func Printf(format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(screen, format, args...)
}

// Flush writes everything that is on the screen buffer, reset it and flush
// stdout.
func Flush() {
	out.WriteString(screen.String())
	screen.Reset()
	out.Flush()
}

// MoveCursorUp moves the cursor <n> cells up.
func MoveCursorUp(n int) {
	fmt.Fprintf(screen, "%s%dA", csi, n)
}

// MoveCursorDown moves the cursor <n> cells down.
func MoveCursorDown(n int) {
	fmt.Fprintf(screen, "%s%dB", csi, n)
}

// MoveCursorForward moves the cursor <n> cells forward.
func MoveCursorForward(n int) {
	fmt.Fprintf(screen, "%s%dC", csi, n)
}

// MoveCursorBack moves the cursor <n> cells back.
func MoveCursorBack(n int) {
	fmt.Fprintf(screen, "%s%dD", csi, n)
}

// MoveCursorPreviousLine moves the cursor to beginning of <n> lines up.
func MoveCursorPreviousLine(n int) {
	fmt.Fprintf(screen, "%s%dF", csi, n)
}

// MoveCursorNextLine moves the cursor to beginning of <n> lines down.
func MoveCursorNextLine(n int) {
	fmt.Fprintf(screen, "%s%dE", csi, n)
}
