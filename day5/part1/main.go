package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := fetch("../input.txt")
	pairs, updates := parse(input)
	adjList := createGraph(pairs)

	res := checkUpdates(adjList, updates)
	fmt.Println("Acutal: ", res)

	testInput := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	testPairs, testUpdates := parse(testInput)
	testAdjList := createGraph(testPairs)
	testRes := checkUpdates(testAdjList, testUpdates)
	fmt.Println("Test Results: ", testRes)
}

func fetch(fileName string) string {
	// Open file
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer f.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
		sb.WriteString("\n")
	}

	return sb.String()
}

func parse(input string) ([][]int, [][]int) {
	pairs := [][]int{}
	updates := [][]int{}

	currentSection := &pairs
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			currentSection = &updates
			continue
		}

		numbers := []int{}
		for _, part := range strings.Split(line, "|") {
			for _, n := range strings.Split(part, ",") {
				if n == "" {
					continue
				}
				val, err := strconv.Atoi(strings.TrimSpace(n))
				if err != nil {
					fmt.Println("Error parsing number: ", err)
					continue
				}
				numbers = append(numbers, val)
			}
		}

		*currentSection = append(*currentSection, numbers)
	}


	return pairs, updates
}

func createGraph(pairs [][]int) map[int][]int {
	graph := make(map[int][]int)

	for _, pair := range pairs {
		if _, found := graph[pair[0]]; !found {
			graph[pair[0]] = []int{}
		}
		graph[pair[0]] = append(graph[pair[0]], pair[1])
	}

	return graph
}

func checkUpdates(adjList map[int][]int, updates [][]int) int {
	res := 0

	for _, update := range updates {
		valid := true
		position := make(map[int]int)
		for i, num := range update {
			position[num] = i
		}

		for center, dependents := range adjList {
			// Ensure the center number exists in the update
			centerPos, exists := position[center]
			if !exists {
				continue
			}

			// Verify that all dependents appear after the center
			for _, dependent := range dependents {
				dependentPos, exists := position[dependent]
				if exists && dependentPos <= centerPos {
					valid = false
					break
				}
			}
		}

		if valid && len(update) > 0 {
			res += update[len(update)/2] // Assuming the first number is the "center"
		}
	}

	return res
}