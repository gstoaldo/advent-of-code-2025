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

func maxArea(input [][]int) (result int) {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			dx := utils.Abs(input[j][0]-input[i][0]) + 1
			dy := utils.Abs(input[j][1]-input[i][1]) + 1
			area := dx * dy
			result = max(result, area)
		}
	}

	return result
}

func main() {
	input := parse(utils.FilePath())
	fmt.Println("p1:", maxArea(input))
}
