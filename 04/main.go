package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func parse(path string) []string {
	return utils.ReadLines(path)
}

func neighbors(i0, j0 int, grid []string) (result [][2]int) {
	H, W := len(grid), len(grid[0])
	for _, delta := range [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}} {
		i, j := i0+delta[0], j0+delta[1]

		if i >= 0 && i < H && j >= 0 && j < W {
			result = append(result, [2]int{i, j})
		}
	}
	return result
}

func canAccess(i, j int, grid []string, removed map[[2]int]bool) bool {
	count := 0

	for _, position := range neighbors(i, j, grid) {
		if removed[position] {
			continue
		}

		if grid[position[0]][position[1]] == '@' {
			count++
		}
	}

	return count < 4
}

func accessedPositions(grid []string, removed map[[2]int]bool) (positions [][2]int) {
	for i, row := range grid {
		for j, val := range row {
			if val == '.' || removed[[2]int{i, j}] {
				continue
			}

			if canAccess(i, j, grid, removed) {
				positions = append(positions, [2]int{i, j})
			}
		}
	}
	return positions
}

func removeCycles(grid []string, maxCycles int) int {
	cycle := 0
	removed := map[[2]int]bool{}

	rollsRemovedInCycle := -1
	for rollsRemovedInCycle != 0 && cycle < maxCycles {
		positions := accessedPositions(grid, removed)
		rollsRemovedInCycle = len(positions)

		for _, position := range positions {
			removed[position] = true
		}

		cycle++
	}

	if cycle == maxCycles && maxCycles > 1 {
		panic("max cycle reached")
	}

	return len(removed)
}

func main() {
	grid := parse(utils.FilePath())

	fmt.Println("p1:", removeCycles(grid, 1))
	fmt.Println("p2:", removeCycles(grid, 100))
}
