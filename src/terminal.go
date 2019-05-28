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

	// ANSI CSI (Control Sequence Introducer).
	// Utilizado para mover o cursor na tela.
	// ESC
	csi = esc + "["
)

// Cores básicas ANSI (3/4 bits).
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

// Criamos um novo buffer para escrever no stdout no método Flush.
var out = bufio.NewWriter(os.Stdout)

// Criamos, também, um buffer para ser utilizado antes da escrita ao stdout.
var screen = new(bytes.Buffer)

// Print escreve os parâmetros no buffer.
func Print(args ...interface{}) (n int, err error) {
	return fmt.Fprint(screen, args...)
}

// Println escreve os parâmetros no buffer, adicionando uma nova linha.
func Println(args ...interface{}) (n int, err error) {
	return fmt.Fprintln(screen, args...)
}

// Printf realiza a formatação e escreve a saída no buffer.
func Printf(format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(screen, format, args...)
}

// Flush escreve tudo que exite no buffer da tela, o reseta e chama o método
// flush para que a escrita no stdout seja realizada.
func Flush() {
	out.WriteString(screen.String())
	screen.Reset()
	out.Flush()
}

// MoveCursorUp move o cursor <n> células para cima.
func MoveCursorUp(n int) {
	fmt.Fprintf(screen, "%s%dA", csi, n)
}

// MoveCursorDown move o cursor <n> células para baixo.
func MoveCursorDown(n int) {
	fmt.Fprintf(screen, "%s%dB", csi, n)
}

// MoveCursorForward move o cursor <n> células para a direita.
func MoveCursorForward(n int) {
	fmt.Fprintf(screen, "%s%dC", csi, n)
}

// MoveCursorBack move o cursor <n> células para a esquerda.
func MoveCursorBack(n int) {
	fmt.Fprintf(screen, "%s%dD", csi, n)
}

// MoveCursorPreviousLine move o cursor para o inicío da n-ésima linha a cima.
func MoveCursorPreviousLine(n int) {
	fmt.Fprintf(screen, "%s%dF", csi, n)
}

// MoveCursorNextLine move o cursor para o início da n-ésima linha abaixo.
func MoveCursorNextLine(n int) {
	fmt.Fprintf(screen, "%s%dE", csi, n)
}

// ClearScreenEnd limpa a tela a partir do cursor até o final.
func ClearScreenEnd() {
	fmt.Fprintf(screen, "%s%dJ", csi, 0)
}
