package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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

	res := tallyMatches(string(data))
	fmt.Println(res)
}

func tallyMatches(input string) int {
	res := 0
	prefix := "mul("
	for i := 0; i <= len(input)-len(prefix); i++ {
		if input[i:i+len(prefix)] != prefix {
			continue
		}

		start := i + len(prefix)
		end := strings.Index(input[start:], ")")
		if end == -1 {
			continue
		}

		end += start
		sub := input[start:end]
		parts := strings.Split(sub, ",")
		if len(parts) != 2 {
			continue
		}

		if !isValidNumber(parts[0]) || !isValidNumber(parts[1]) {
			continue
		}

		n1, _ := strconv.Atoi(parts[0])
		n2, _ := strconv.Atoi(parts[1])
		res += n1*n2
	}

	return res
}

func isValidNumber(s string) bool {
	if len(s) < 1 || len(s) > 3 {
		return false
	}
	_, err := strconv.Atoi(s)
	return err == nil
}
