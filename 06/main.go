package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func split(line string) []string {
	re := regexp.MustCompile(`\s+`) // split by one or more spaces
	return re.Split(strings.TrimSpace(line), -1)
}

func parse(path string) ([][]int, []string) {
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
	problems, operations := parse(utils.FilePath())
	fmt.Println("p1:", total(problems, operations))
}
