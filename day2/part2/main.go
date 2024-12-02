package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	records := readCSV("../input.csv")
	res := 0

	for _, record := range records {
		if dampenedCheck(record) {
			res++
		}
	}

	fmt.Println(res)
}

func dampenedCheck(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if checkAscending(nums, i) || checkDescending(nums, i) {
			return true
		}
	}
	return false
}

func checkAscending(nums []int, skippedIndex int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if i == skippedIndex {
			continue
		}
		nextIndex := i+1
		if nextIndex == skippedIndex {
			if nextIndex == len(nums)-1 {
				continue
			}
			nextIndex = i+2
		}
		if nums[i] >= nums[nextIndex] || nums[nextIndex]-nums[i] > 3 {
			return false
		}
	}

	return true
}

func checkDescending(nums []int, skippedIndex int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if i == skippedIndex {
			continue
		}
		nextIndex := i+1
		if nextIndex == skippedIndex {
			if nextIndex == len(nums)-1 {
				continue
			}
			nextIndex = i+2
		}
		if nums[i] <= nums[nextIndex] || nums[i]-nums[nextIndex] > 3 {
			return false
		}
	}

	return true
}

// takes in filename and returns numbers grouped by lines
func readCSV(fileName string) ([][]int) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	res := [][]int{}
	for _, line := range records {
		recordStr := strings.Split(line[0], " ")
		record := []int{}

		for _, numberStr := range recordStr {
			number, _ := strconv.Atoi(numberStr)
			record = append(record, number)
		}

		res = append(res, record)
	}

	return res
}
