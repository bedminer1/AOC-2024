package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data := fetch("../input.txt")
	locationMap, rowLen, _ := parseLocations(data)
	res := tallyAntinodes(locationMap, rowLen)
	fmt.Println("Actual Result: ", res)

	testInput := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
	testLocationMap, testRowLen, _ := parseLocations(testInput)
	testRes := 0
	testRes += tallyAntinodes(testLocationMap, testRowLen)

	fmt.Println("Test Result: ", testRes)
}

type coords struct {
	x, y int
}

func fetch(fileName string) string {
	f, _ := os.Open(fileName)
	defer f.Close()

	data, _ := io.ReadAll(f)
	return string(data)
}

func parseLocations(data string) (map[rune][]coords, int, int) {
	res := make(map[rune][]coords)

	lines := strings.Split(data, "\n")
	for x, line := range lines {
		for y, c := range line {
			if c == '.' {
				continue
			}

			if _, found := res[c]; !found {
				res[c] = []coords{}
			}
			res[c] = append(res[c], coords{x: x, y: y})
		}
	}

	return res, len(lines), len(lines[0])
}

func tallyAntinodes(locationMap map[rune][]coords, mapBoundary int) int {
	uniqueAntinodes := make(map[coords]struct{})

	isWithinMap := func(x, y int) bool {
		return x >= 0 && x < mapBoundary && y >= 0 && y < mapBoundary
	}

	for _, antennaLocations := range locationMap {
		if len(antennaLocations) < 2 {
			continue
		}

		for i := 0; i < len(antennaLocations)-1; i++ {
			for j := i + 1; j < len(antennaLocations); j++ {
				a1 := antennaLocations[i]
				a2 := antennaLocations[j]

				dx := a2.x - a1.x
				dy := a2.y - a1.y

				extendLine := func(start coords, dirX, dirY int) {
					x, y := start.x, start.y
					for isWithinMap(x, y) {
						uniqueAntinodes[coords{x, y}] = struct{}{}
						x += dirX
						y += dirY
					}
				}

				extendLine(a1, dx, dy)
				extendLine(a2, -dx, -dy)
			}
		}
	}

	return len(uniqueAntinodes)
}
