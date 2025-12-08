package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func parse(path string) ([]string, int) {
	lines := utils.ReadLines(path)

	for i, char := range lines[0] {
		if char == 'S' {
			return lines, i
		}
	}

	panic("start not found")
}

func run(lines []string, start int) int {
	beams := map[int]bool{
		start: true,
	}

	result := 0

	for i := range lines {
		for j := range beams {
			if lines[i][j] == '^' {
				result++
				beams[j-1] = true
				beams[j+1] = true
				delete(beams, j)
			}
		}
	}

	return result
}

func main() {
	lines, start := parse(utils.FilePath())
	fmt.Println("p1:", run(lines, start))
}
