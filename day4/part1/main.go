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
	matchStr := "XMAS"
	reverseMatchStr := "SAMX"

	rows := strings.Split(input, "\n")
	numRows := len(rows)
	numCols := len(rows[0])

	// Check horizontal
	for _, row := range rows {
		for i := 0; i <= numRows-len(matchStr); i++ {
			if row[i:i+len(matchStr)] == matchStr || row[i:i+len(matchStr)] == reverseMatchStr {
				res++
			}
		}
	}
	// Check vertical
	for c := 0; c < numCols; c++ {
		for r := 0; r <= numRows-len(matchStr); r++ {
			var sb strings.Builder
			for k := 0; k < len(matchStr); k++ {
				sb.WriteByte(rows[r+k][c])
			}
			curr := sb.String()
			if curr == matchStr || curr == reverseMatchStr {
				res++
			}
		}
	}

	// Check top-left to bottom-right diagonal matches
	for r := 0; r <= numRows-len(matchStr); r++ {
		for c := 0; c <= numCols-len(matchStr); c++ {
			var sb strings.Builder
			for k := 0; k < len(matchStr); k++ {
				sb.WriteByte(rows[r+k][c+k])
			}
			curr := sb.String()
			if curr == matchStr || curr == reverseMatchStr {
				res++
			}
		}
	}

	// Check top-right to bottom-left diagonal matches
	for r := 0; r <= numRows-len(matchStr); r++ {
		for c := len(matchStr) - 1; c < numCols; c++ {
			var sb strings.Builder
			for k := 0; k < len(matchStr); k++ {
				sb.WriteByte(rows[r+k][c-k])
			}
			curr := sb.String()
			if curr == matchStr || curr == reverseMatchStr {
				res++
			}
		}
	}

	return res
}
