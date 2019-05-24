package main

import (
	"fmt"
)

const (
	// ESC[ --> ESC = \033 (ASCII octal)
	esc = "\033"

	// ANSI Control Sequence Introducer used to move cursor on the screen:
	// ESC
	csi = esc + "["
)

// MoveCursorUp moves the cursor <n> cells up.
func MoveCursorUp(n int) {
	fmt.Printf("%s%dA", csi, n)
}

// MoveCursorDown moves the cursor <n> cells down.
func MoveCursorDown(n int) {
	fmt.Printf("%s%dB", csi, n)
}

// MoveCursorForward moves the cursor <n> cells forward.
func MoveCursorForward(n int) {
	fmt.Printf("%s%dC", csi, n)
}

// MoveCursorBack moves the cursor <n> cells back.
func MoveCursorBack(n int) {
	fmt.Printf("%s%dD", csi, n)
}

// MoveCursorPreviousLine moves the cursor to beginning of <n> lines up.
func MoveCursorPreviousLine(n int) {
	fmt.Printf("%s%dF", csi, n)
}

// SaveCursorPos saves the current position.
func SaveCursorPos() {
	fmt.Printf("%ss", csi)
}

// RestoreCursorPos restores the current position.
func RestoreCursorPos() {
	fmt.Printf("%su", csi)
}

// HideCursor hides the cursor.
func HideCursor() {
	fmt.Printf("%s?25l", csi)
}

// ShowCursor shows the cursor.
func ShowCursor() {
	fmt.Printf("%s?25h", csi)
}
