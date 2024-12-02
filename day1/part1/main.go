package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums1, nums2 := readCSV("../input.csv")

	sort.Ints(nums1)
	sort.Ints(nums2)

	fmt.Println("Col 1: ", nums1[:10])
	fmt.Println("Col 2: ", nums2[:10])

	res := 0
	for i := range nums1 {
		num1 := nums1[i]
		num2 := nums2[i]

		diff := num1 - num2
		if diff < 0 {
			diff *= -1
		}

		res += diff
	}

	fmt.Println(res)
}

// takes in filename and returns numbers in the 2 columns
func readCSV(fileName string) ([]int, []int) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil, nil
	}

	var (
		nums1 []int
		nums2 []int
	)

	for _, line := range records {
		record := strings.Split(line[0], "   ")

		val1, err1 := strconv.Atoi(record[0])
		val2, err2 := strconv.Atoi(record[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting row to integers:", record, err1, err2)
			continue
		}

		nums1 = append(nums1, val1)
		nums2 = append(nums2, val2)
	}

	return nums1, nums2
}