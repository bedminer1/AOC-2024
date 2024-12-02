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
		isAscending := false
		if record[0] < record[1] {
			isAscending = true
		} 

		if isAscending && checkAscending(record) {
			res++
		} else if checkDescending(record){
			res++
		}
	}

	fmt.Println(res)
}

func checkAscending(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] || nums[i+1]-nums[i] > 3 {
			return false
		}
	}

	return true
}

func checkDescending(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] || nums[i]-nums[i+1] > 3 {
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
