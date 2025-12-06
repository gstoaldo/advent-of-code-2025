package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func split(line string) []string {
	re := regexp.MustCompile(`\s+`) // split by one or more spaces
	return re.Split(strings.TrimSpace(line), -1)
}

func parseP1(path string) ([][]int, []string) {
	lines := utils.ReadLines(path)
	problemSize := len(lines) - 1
	nproblems := len(split(lines[0]))

	problems := make([][]int, nproblems)
	for i := range problems {
		problems[i] = make([]int, problemSize)
	}

	for i := range problemSize {
		lineNums := split(lines[i])
		for j, numStr := range lineNums {
			problems[j][i] = utils.ToInt(numStr)
		}
	}

	operations := split(lines[len(lines)-1])

	return problems, operations
}

func parseP2(path string) ([][]int, []string) {
	lines := utils.ReadLines(path)
	problemSize := len(lines) - 1

	problems := [][]int{}
	nums := []int{}

	// read chas top-down, right-left and concatenate each column
	// if column is empty, starts a new problem
	for j := len(lines[0]) - 1; j >= 0; j-- {
		numStr := ""
		for i := range problemSize {
			numStr += string(lines[i][j])
		}

		numStr = strings.TrimSpace(numStr)

		if numStr == "" {
			// empty column
			problems = append(problems, nums)
			nums = []int{}
		} else {
			nums = append(nums, utils.ToInt(numStr))
		}
	}

	problems = append(problems, nums)

	operations := split(lines[len(lines)-1])
	slices.Reverse(operations)

	return problems, operations
}

func total(problems [][]int, operations []string) (result int) {
	calc := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"*": func(a, b int) int { return a * b },
	}

	for i, nums := range problems {
		op := operations[i]

		problemResult := nums[0]
		for _, n := range nums[1:] {
			problemResult = calc[op](problemResult, n)
		}

		result += problemResult
	}

	return result
}

func main() {
	fmt.Println("p1:", total(parseP1(utils.FilePath())))
	fmt.Println("p2:", total(parseP2(utils.FilePath())))
}
