package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	gameMap := fetch("../input.txt")
	res := findObstacleSpots(gameMap)
	fmt.Println("Actual Results: ", res)

	testInput := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	lines := strings.Split(testInput, "\n")
	testGameMap := make([][]byte, len(lines))
	for i, line := range lines {
		testGameMap[i] = []byte(line)
	}
	testRes := findObstacleSpots(testGameMap)
	fmt.Println("Test Results: ", testRes)
}

func findObstacleSpots(gameMap [][]byte) int {
	var result int
	var guardX, guardY int

	// Locate the guard's initial position
	for i, row := range gameMap {
		for j, cell := range row {
			if cell == '^' {
				guardX, guardY = i, j
				break
			}
		}
	}

	// Precompute the guard's movement path without obstacles
	initialPath := precomputeGuardPath(gameMap, guardX, guardY, 1)

	// Test each empty spot for placing an obstacle
	for i, row := range gameMap {
		for j, cell := range row {
			if cell == '#' { // Skip existing obstacles
				continue
			}

			// Skip positions not in the initial path
			key := fmt.Sprintf("%d,%d", i, j)
			foundInPath := false
			for _, p := range initialPath {
				if len(p) >= len(key) && p[:len(key)] == key { // Ensure p is long enough
					foundInPath = true
					break
				}
			}
			if !foundInPath {
				continue
			}

			// Place temporary obstacle
			gameMap[i][j] = '#'

			// Check if the guard is stuck in a loop
			if hasCycle(gameMap, guardX, guardY, 1) {
				result++
			}

			// Remove the temporary obstacle
			gameMap[i][j] = cell
		}
	}

	return result
}


func precomputeGuardPath(gameMap [][]byte, startX, startY, startDirection int) []string {
	visited := make(map[string]struct{})
	path := []string{}
	x, y, direction := startX, startY, startDirection

	for {
		// Check if out of bounds
		if x < 0 || x >= len(gameMap) || y < 0 || y >= len(gameMap[0]) {
			break // Guard has left the map
		}

		// Encode the position and direction as a unique key
		key := fmt.Sprintf("%d,%d,%d", x, y, direction)
		if _, ok := visited[key]; ok {
			path = append(path, key)
			break // Cycle detected
		}

		// Mark current position and direction as visited
		visited[key] = struct{}{}
		path = append(path, key)

		// Move based on the current direction
		switch direction {
		case 1: // Up
			if x-1 >= 0 && gameMap[x-1][y] == byte('#') {
				direction = 2 // Turn right
			} else {
				x-- // Move up
			}
		case 2: // Right
			if y+1 < len(gameMap[0]) && gameMap[x][y+1] == byte('#') {
				direction = 3 // Turn down
			} else {
				y++ // Move right
			}
		case 3: // Down
			if x+1 < len(gameMap) && gameMap[x+1][y] == byte('#') {
				direction = 4 // Turn left
			} else {
				x++ // Move down
			}
		case 4: // Left
			if y-1 >= 0 && gameMap[x][y-1] == byte('#') {
				direction = 1 // Turn up
			} else {
				y-- // Move left
			}
		}
	}

	return path
}


func hasCycle(gameMap [][]byte, startX, startY, startDirection int) bool {
	visited := make(map[string]struct{})
	x, y, direction := startX, startY, startDirection

	for {
		// Check if out of bounds
		if x < 0 || x >= len(gameMap) || y < 0 || y >= len(gameMap[0]) {
			return false // Guard has left the map; no cycle
		}

		// Encode the position and direction as a unique key
		key := fmt.Sprintf("%d,%d,%d", x, y, direction)
		if _, ok := visited[key]; ok {
			return true // Cycle detected
		}

		// Mark current position and direction as visited
		visited[key] = struct{}{}

		// Move based on the current direction
		switch direction {
		case 1: // Up
			if x-1 >= 0 && gameMap[x-1][y] == byte('#') {
				direction = 2 // Turn right
			} else {
				x-- // Move up
			}
		case 2: // Right
			if y+1 < len(gameMap[0]) && gameMap[x][y+1] == byte('#') {
				direction = 3 // Turn down
			} else {
				y++ // Move right
			}
		case 3: // Down
			if x+1 < len(gameMap) && gameMap[x+1][y] == byte('#') {
				direction = 4 // Turn left
			} else {
				x++ // Move down
			}
		case 4: // Left
			if y-1 >= 0 && gameMap[x][y-1] == byte('#') {
				direction = 1 // Turn up
			} else {
				y-- // Move left
			}
		}
	}
}


func fetch(fileName string) [][]byte {
	f, _ := os.Open(fileName)
	defer f.Close()

	data, _ := io.ReadAll(f)
	m := bytes.Split(data, []byte("\n"))

	return m
}