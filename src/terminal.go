package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	// ESC[ --> ESC = \033 (ASCII octal)
	esc = "\033"

	// ANSI Control Sequence Introducer used to move cursor on the screen:
	// ESC
	csi = esc + "["
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
	// When stopping program with Ctrl-C output may not be written fully. Thus,
	// we need to handle the interrupt signal to flush everything before
	// exiting program.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	flushed := make(chan bool, 1)

	go func() {
		// Receive value from channel.
		<-sig

		// Receive value saying that everything was flushed.
		<-flushed

		// Now program can be stopped.
		os.Exit(0)
	}()

	for _, str := range screen.String() {
		out.WriteRune(str)
	}

	screen.Reset()
	out.Flush()
	flushed <- true
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

// HideCursor hides the cursor.
func HideCursor() {
	fmt.Fprintf(screen, "%s?25l", csi)
}

// ShowCursor shows the cursor.
func ShowCursor() {
	fmt.Fprintf(screen, "%s?25h", csi)
}
