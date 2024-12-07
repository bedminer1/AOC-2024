package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	targets, numbers := fetch("../input.txt")
	res := 0

	for i := 0; i < len(targets); i++ {
		if checkValid(numbers[i], targets[i]) {
			res += targets[i]
		}
	}

	fmt.Println("Actual Result: ", res)
}

func checkValid(nums []int, target int) bool {
	return helper(nums, target, nums[0], 1)
}

func helper(nums []int, target, current int, index int) bool {
	// Base case: if we've processed all numbers
	if index == len(nums) {
		return current == target
	}

	// Try adding the current number
	if helper(nums, target, current+nums[index], index+1) {
		return true
	}

	// Try multiplying the current number
	if helper(nums, target, current*nums[index], index+1) {
		return true
	}

	// If neither operation works, return false
	return false
}

func fetch(fileName string) ([]int, [][]int) {
	targets := []int{}
	numbers := [][]int{}

	f, _ := os.Open(fileName)
	data, _ := io.ReadAll(f)

	equations := bytes.Split(data, []byte("\n"))

	for _, equation := range equations {
		parts := bytes.Split(equation, []byte(":"))
		target, _ := strconv.Atoi(string(parts[0]))
		targets = append(targets, target)

		numbersStr := strings.Split(string(parts[1]), " ")
		newNums := []int{}
		for _, numStr := range numbersStr {
			num, _ := strconv.Atoi(numStr)
			newNums = append(newNums, num)
		}
		numbers = append(numbers, newNums)
	}

	return targets, numbers
}
