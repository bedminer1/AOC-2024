package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	gameMap := fetch("../input.txt")

	var x, y int
	// 1=up, 2=right, 3=down, 4=left
	direction := 1
	visited := make(map[string]struct{})

	for i, row := range gameMap {
		for j := range row {
			if gameMap[i][j] == byte('^') {
				x = i
				y = j
			}
		}
	}

	for {
		// out of bounds
		if x < 0 || x >= len(gameMap) || y < 0 || y >= len(gameMap[0]) {
			break
		}

		switch direction {
		case 1:
			if x-1 >= 0 && gameMap[x-1][y] == byte('#') {
				direction++
				continue
			}

			// go up
			x--
		case 2:
			if y+1 < len(gameMap[0]) && gameMap[x][y+1] == byte('#') {
				direction++
				continue
			}

			// go right
			y++
		case 3:
			if x+1 < len(gameMap) && gameMap[x+1][y] == byte('#') {
				direction++
				continue
			}

			// go down
			x++
		case 4:
			if y-1 >= 0 && gameMap[x][y-1] == byte('#') {
				direction = 1
				continue
			}

			// go left
			y--
		}
		coords := fmt.Sprintf("%d,%d", x, y)
		visited[coords] = struct{}{}
	}

	// number of travelled places
	fmt.Println(len(visited))
}

func fetch(fileName string) [][]byte {
	f, _ := os.Open(fileName)
	defer f.Close()

	data, _ := io.ReadAll(f)
	m := bytes.Split(data, []byte("\n"))

	return m
}