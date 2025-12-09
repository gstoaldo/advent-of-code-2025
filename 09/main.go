package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func parse(path string) (input [][]int) {
	for _, line := range utils.ReadLines(path) {
		splited := strings.Split(line, ",")
		input = append(input, []int{utils.ToInt(splited[0]), utils.ToInt(splited[1])})
	}

	return input
}

func area(p1, p2 []int) int {
	dx := utils.Abs(p2[0]-p1[0]) + 1
	dy := utils.Abs(p2[1]-p1[1]) + 1
	return dx * dy
}

func maxAreaP1(input [][]int) (result int) {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			result = max(result, area(input[i], input[j]))
		}
	}

	return result
}

// cache makes a HUGE difference
var cache = make(map[[2]int]bool)

func pointIsInside(input [][]int, p []int) bool {
	// (Ray casting algorithm)
	// Count the number of vertical edges to the right.
	// The vertice is inside if the count is odd.
	if val, ok := cache[[2]int{p[0], p[1]}]; ok {
		return val
	}

	count := 0
	wrapped := make([][]int, len(input)+1)
	copy(wrapped, input)
	wrapped[len(input)] = input[0]

	for i := 0; i < len(wrapped)-1; i++ {
		j := i + 1

		minx, maxx := min(wrapped[i][0], wrapped[j][0]), max(wrapped[i][0], wrapped[j][0])
		miny, maxy := min(wrapped[i][1], wrapped[j][1]), max(wrapped[i][1], wrapped[j][1])

		// In the perimeter
		if p[0] >= minx && p[0] <= maxx && p[1] >= miny && p[1] <= maxy {
			cache[[2]int{p[0], p[1]}] = true
			return true
		}

		if maxx > p[0] && p[1] >= miny && p[1] < maxy {
			count++
		}
	}

	result := count%2 != 0

	cache[[2]int{p[0], p[1]}] = result

	return result
}

func areaIsInside(input [][]int, p1, p2 []int) bool {
	// Check if the area (given by the opposite vertices p1, p2) is inside the input area.
	// Check just the points in the area perimeter, don't have to check every single point inside the area.
	minx, maxx := min(p1[0], p2[0]), max(p1[0], p2[0])
	miny, maxy := min(p1[1], p2[1]), max(p1[1], p2[1])

	for x := minx; x <= maxx; x++ {
		if !pointIsInside(input, []int{x, miny}) || !pointIsInside(input, []int{x, maxy}) {
			return false
		}
	}

	for y := miny; y <= maxy; y++ {
		if !pointIsInside(input, []int{minx, y}) || !pointIsInside(input, []int{maxx, y}) {
			return false
		}
	}

	return true
}

func maxAreaP2(input [][]int) (result int) {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			p1, p2 := input[i], input[j]

			if areaIsInside(input, p1, p2) {
				result = max(result, area(p1, p2))
			}
		}
	}

	return result
}

func main() {
	input := parse(utils.FilePath())
	fmt.Println("p1:", maxAreaP1(input))
	fmt.Println("p2:", maxAreaP2(input))
}
