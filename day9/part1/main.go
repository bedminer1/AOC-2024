package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := fetch("../input.txt")

	uncompressedInput := uncompress(input)
	res := flush(uncompressedInput)
	f, _ := os.Create("output.txt")
	defer f.Close()
	f.WriteString(res)

	testInput := `2333133121414131402`
	uncompressedTestInput := uncompress(testInput)
	testRes := flush(uncompressedTestInput)
	fmt.Println("Test Result: ", testRes)

}

func fetch(fileName string) string {
	f, _ := os.Open(fileName)
	defer f.Close()
	data, _ := io.ReadAll(f)
	return string(data)
}

func uncompress(input string) string {
	id := 0
	var sb strings.Builder

	for i, c := range input {
		size := int(c - '0')
		toPrint := strconv.Itoa(id)
		if i%2 != 0 {
			id++
			toPrint = "."
		}

		for j := 0; j < size; j++ {
			sb.WriteString(toPrint)
		}
	}

	return sb.String()
}

func flush(input string) string {
	left, right := 0, len(input)-1

	var sb strings.Builder
	var empty strings.Builder

	for left < right {
		for input[left] != '.' {
			sb.WriteByte(input[left])
			left++
			if left >= right { break }
		}

		for input[left] == '.' {
			for input[right] == '.' {
				empty.WriteByte('.')
				right--
			}
			sb.WriteByte(input[right])
			empty.WriteByte('.')
			left++
			right--
			if left >= right { break }
		}
	}
	sb.WriteString(empty.String())

	return sb.String()
}