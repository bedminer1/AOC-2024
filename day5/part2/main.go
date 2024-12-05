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

		for before, afters := range adjList {
			beforePos, exists := position[before]
			if !exists {
				continue
			}

			for _, after := range afters {
				afterPos, exists := position[after]
				if exists && afterPos <= beforePos {
					valid = false
					break
				}
			}
		}

		if !valid {
			// If not valid, attempt to reorder
			reordered := reorder(update, adjList)
			if len(reordered) > 0 {
				res += reordered[len(reordered)/2]
			}
		}
	}

	return res
}

func reorder(update []int, adjList map[int][]int) []int {
	inDegree := make(map[int]int)
	graph := make(map[int][]int)

	for _, num := range update {
		if _, exists := graph[num]; !exists {
			graph[num] = []int{}
		}
	}

	for center, dependents := range adjList {
		for _, dependent := range dependents {
			// Add edges only if both center and dependent are in the update list
			if contains(update, center) && contains(update, dependent) {
				graph[center] = append(graph[center], dependent)
				inDegree[dependent]++
			}
		}
	}

	for _, num := range update {
		if _, exists := inDegree[num]; !exists {
			inDegree[num] = 0
		}
	}

	queue := []int{}
	for _, num := range update {
		if inDegree[num] == 0 {
			queue = append(queue, num)
		}
	}

	reordered := []int{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // remove from queue

		reordered = append(reordered, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check if all nodes in the update are used in the reordered result
	if len(reordered) == len(update) {
		return reordered
	}
	
	return []int{} // Return empty if reordering fails
}

// Helper function to check if a slice contains a given value
func contains(slice []int, value int) bool {
	for _, num := range slice {
		if num == value {
			return true
		}
	}
	return false
}
