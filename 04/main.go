package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func parse(path string) []string {
	return utils.ReadLines(path)
}

func neighbors(i0, j0 int, grid []string) (result [][]int) {
	H, W := len(grid), len(grid[0])
	for _, delta := range [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}} {
		i, j := i0+delta[0], j0+delta[1]

		if i >= 0 && i < H && j >= 0 && j < W {
			result = append(result, []int{i, j})
		}
	}
	return result
}

func canAccess(i, j int, grid []string) bool {
	count := 0

	for _, position := range neighbors(i, j, grid) {
		if grid[position[0]][position[1]] == '@' {
			count++
		}
	}

	return count < 4
}

func countCanAccess(grid []string) (result int) {
	for i, row := range grid {
		for j, val := range row {
			if val == '.' {
				continue
			}

			if canAccess(i, j, grid) {
				result++
			}
		}
	}
	return result
}

func main() {
	grid := parse(utils.FilePath())

	fmt.Println("p1:", countCanAccess(grid))
}
