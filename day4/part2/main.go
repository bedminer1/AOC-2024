package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Open file
	f, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	res := checkMatches(string(data))

	fmt.Println("Actual Result: ", res)

	// testing
	testInput := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	testRes := checkMatches(testInput)
	fmt.Println("Test Result: ", testRes)
}

func checkMatches(input string) int {
	res := 0
	matchStr := "MAS"
	reverseMatchStr := "SAM"

	rows := strings.Split(input, "\n")
	numRows := len(rows)
	numCols := len(rows[0])

	for r := 0; r <= numRows-3; r++ {
		for c := 0; c <= numCols-3; c++ {
			// Check the diagonals
			topLeftBottomRight := string([]byte{rows[r][c], rows[r+1][c+1], rows[r+2][c+2]})
			topRightBottomLeft := string([]byte{rows[r][c+2], rows[r+1][c+1], rows[r+2][c]})

			if (topLeftBottomRight == matchStr || topLeftBottomRight == reverseMatchStr) && (topRightBottomLeft == matchStr || topRightBottomLeft == reverseMatchStr) {
				res++
			}
		}
	}

	return res
}
