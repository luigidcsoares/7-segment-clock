package main

import (
	"time"
)

// NowAsMatrix returns the current time in the form of a matrix of 3 rows and 2
// columns, where each row refer to a piece of the time (hour, minute or
// second) and each column indicates one algarism (HH:MM:SS).
func NowAsMatrix() (now [3][2]int) {
	h, m, s := time.Now().Clock()
	now[0] = [2]int{h / 10, h % 10}
	now[1] = [2]int{m / 10, m % 10}
	now[2] = [2]int{s / 10, s % 10}
	return
}
