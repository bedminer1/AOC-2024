package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := fetch("../input.txt")
	parsedInput := parse(input)
	res := tallyScores(parsedInput)
	fmt.Println("Actual Results: ", res)

	testInput := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	testParsedInput := parse(testInput)
	testRes := tallyScores(testParsedInput)
	fmt.Println("Test Result: ", testRes)
}

func fetch(fileName string) string {
	f, _ := os.Open(fileName)
	defer f.Close()

	data, _ := io.ReadAll(f)
	return string(data)
}

func parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	res := make([][]int, len(lines))
	for i := range res {
		res[i] = make([]int, len(lines[0]))
	}

	for i, line := range lines {
		for j, c := range line {
			res[i][j] = int(c - '0')
		}
	}

	return res
}

func tallyScores(input [][]int) int {
	directions := [][2]int{{-1,0},{1,0},{0,-1},{0,1}}
	rows, cols := len(input), len(input[0])
	visited := make([][]bool, cols)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	counted := make([][]bool, rows)
    for i := range counted {
        counted[i] = make([]bool, cols)
    }

	var dfs func(x, y, current int) int
	dfs = func(x, y, current int) int {
		if x < 0 || x >= rows || y < 0 || y >= cols || visited[x][y] || input[x][y] != current {
			return 0
		}

		if current == 9 {
			if counted[x][y] == false {
				counted[x][y] = true
				return 1
			}
			return 0
		}

		visited[x][y] = true
		trails := 0

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			trails += dfs(nx, ny, current+1)
		}

		visited[x][y] = false
		return trails
	}

	res := 0

	for x, row := range input {
		for y, value := range row {
			if value == 0 {
				for i := range counted {
                    for j := range counted[i] {
                        counted[i][j] = false
                    }
                }
				// dfs for trails
				res += dfs(x, y, 0)
			}
		}
	}

	return res
}