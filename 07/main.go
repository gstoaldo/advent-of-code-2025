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

func timeline(lines []string, start int) int {
	type key struct{ i, j int }
	cache := map[key]int{}
	var run func(i, j int) int

	run = func(i, j int) int {
		if val, ok := cache[key{i, j}]; ok {
			return val
		}

		if i == len(lines)-1 {
			return 1
		}

		if lines[i][j] == '^' {
			result := run(i+1, j-1) + run(i+1, j+1)
			cache[key{i, j}] = result
			return result
		}

		result := run(i+1, j)
		cache[key{i, j}] = result

		return result
	}

	return run(0, start)
}

func main() {
	lines, start := parse(utils.FilePath())
	fmt.Println("p1:", run(lines, start))
	fmt.Println("p2:", timeline(lines, start))
}
