package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	targets, numbers := fetch("../input.txt")

	res := parallelCheck(targets, numbers)
	fmt.Println("Actual Result: ", res)

	var testTargets = []int{
		190, 3267, 83, 156, 7290, 161011, 192, 21037, 292,
	}
	var testNumbers = [][]int{
		{10, 19},
		{81, 40, 27},
		{17, 5},
		{15, 6},
		{6, 8, 6, 15},
		{16, 10, 13},
		{17, 8, 14},
		{9, 7, 18, 13},
		{11, 6, 16, 20},
	}
	testRes := parallelCheck(testTargets, testNumbers)
	fmt.Println("Test Result: ", testRes)
}

func parallelCheck(targets []int, numbers [][]int) int {
	var wg sync.WaitGroup
	resChan := make(chan int, len(targets))

	// Spawn a Goroutine for each target
	for i := 0; i < len(targets); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if checkValid(numbers[i], targets[i]) {
				resChan <- targets[i] // Send result to channel
			} else {
				resChan <- 0 // Send 0 if invalid
			}
		}(i)
	}

	// Close the channel after all Goroutines complete
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// Collect results
	res := 0
	for val := range resChan {
		res += val
	}

	return res
}

func checkValid(nums []int, target int) bool {
	return helper(nums, target, nums[0], 1)
}

func helper(nums []int, target, current int, index int) bool {
	if current > target {
		return false
	}

	// Base case: if we've processed all numbers
	if index == len(nums) {
		return current == target
	}

	nextNum := nums[index]

	// Try adding the next number
	if helper(nums, target, current+nextNum, index+1) {
		return true
	}

	// Try multiplying by the next number
	if helper(nums, target, current*nextNum, index+1) {
		return true
	}

	// Try concatenating the current and next numbers
	concatNum, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, nextNum))
	if helper(nums, target, concatNum, index+1) {
		return true
	}

	// If none of the operations work, return false
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
