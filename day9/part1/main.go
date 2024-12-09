package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input := fetch("../input.txt")

	uncompressedInput := uncompress(input)
	flushedData := flush(uncompressedInput)
	res := calculateCheckSum(flushedData)
	fmt.Println("Actual Result: ", res)

	testInput := `2333133121414131402`

	uncompressedTestInput := uncompress(testInput)
	testFlushedData := flush(uncompressedTestInput)
	testRes := calculateCheckSum(testFlushedData)
	fmt.Println("Test Result: ", testRes)
}

func fetch(fileName string) string {
	f, _ := os.Open(fileName)
	defer f.Close()
	data, _ := io.ReadAll(f)
	return string(data)
}

func uncompress(input string) []int {
	id := 0
	res := []int{}

	for i, c := range input {
		size := int(c - '0')
		toPrint := id
		if i%2 != 0 {
			id++
			toPrint = -1
		}

		for j := 0; j < size; j++ {
			res = append(res, toPrint)
		}
	}

	return res
}

func flush(input []int) []int {
	left, right := 0, len(input)-1

	res := []int{}

	for left <= right {
		for input[left] != -1 {
			res = append(res, input[left])
			left++
			if left > right {
				break
			}
		}

		for input[left] == -1 {
			for input[right] == -1 {
				right--
			}
			res = append(res, input[right])
			left++
			right--
			if left > right {
				break
			}
		}
	}

	return res
}

func calculateCheckSum(input []int) int {
	res := 0
	for i, num := range input {
		res += i * num
	}

	return res
}
