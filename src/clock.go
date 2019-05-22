package main

import "fmt"

// Each position of this array describes how a digit will be displayed as a
// 7-segment clock digit. For each digit, a segment could be false or true.
// Segments are shown below:
//    __1__
// 0 |     | 2
//   |__3__|
// 4 |     | 6
//   |__5__|
var digitSegments = [][]bool{
	{true, true, true, false, true, true, true},     // 0
	{false, false, true, false, false, false, true}, // 1
	{false, true, true, true, true, true, false},    // 2
	{false, true, true, true, false, true, true},    // 3
	{true, false, true, true, false, false, true},   // 4
	{true, true, false, true, false, true, true},    // 5
	{true, true, false, true, true, true, true},     // 6
	{false, true, true, false, false, false, true},  // 7
	{true, true, true, true, true, true, true},      // 8
	{true, true, true, true, false, true, true},     // 9
}

func buildDigit(digitValue, segSize int) (digit [][]byte) {
	// We have two segments in vertical axis, so the actual number of lines
	// should be 2 * segment size + 1 (first row has only the top segment).
	digit = make([][]byte, 2*segSize+1)

	for i := range digit {
		// Vertical segments are made by underscores with size equal to segment
		// size. On the other hand, horizontal segments (top, mid, bottom) are
		// made by double underscores. So, for segment size = 1 we have 2
		// underscores; for segment size = 2, 4 undescores. The first and last
		// column must be taken into account (the + 2 element).
		digit[i] = make([]byte, 2+2*segSize)

		for j := range digit[i] {
			// Init each position with blank space.
			digit[i][j] = ' '
		}
	}

	// Horizontal segments (1, 3, 5)
	segment := 1
	for i := 0; i < 3; i++ {
		row := i * segSize

		// Ex.: row 0 --> segment 1
		for j := 1; j < len(digit[row])-1; j++ {
			if digitSegments[digitValue][segment] {
				digit[row][j] = '_'
			}
		}

		segment += 2
	}

	// Vertical segments (0, 2, 4, 6)
	// Left: 0, 4 (col = 0)
	// Right: 2, 6 (col = last)
	for i := 0; i <= 6; i += 2 {
		// Segments 0, 2 --> top rows
		// Segments 4, 6 --> bottom rows
		add := 0
		if i >= 4 {
			add = segSize
		}

		row := 1 + add

		// 0; 4 % 4 = 0 --> Left segments
		// 2; 6 % 4 = 2 --> Right segments
		// 0 >> 1 = 0; 2 >> 1 = 1
		col := (len(digit[row]) - 1) * (i % 4 >> 1)

		for j := row; j < row+segSize; j++ {
			if digitSegments[digitValue][i] {
				digit[j][col] = '|'
			}
		}
	}

	return
}

func printDigit(digit [][]byte) {
	for i := range digit {
		for j := range digit[i] {
			fmt.Printf("%c", digit[i][j])
		}

		fmt.Println()
	}

	fmt.Println()
}

func main() {
	for d := 0; d <= 9; d++ {
		digit := buildDigit(d, 4)
		printDigit(digit)
	}
}
