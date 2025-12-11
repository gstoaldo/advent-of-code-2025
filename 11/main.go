package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

type G map[string][]string

func parse(path string) G {
	input := G{}

	for _, line := range utils.ReadLines(path) {
		splitted := strings.Split(line, ": ")

		origin, destStr := splitted[0], splitted[1]
		dest := strings.Split(destStr, " ")

		input[origin] = dest
	}

	return input
}

func countPaths(input G) int {
	visited := map[string]bool{}

	var dfs func(curr string) int

	dfs = func(curr string) int {
		if curr == "out" {
			return 1
		}

		count := 0

		for _, n := range input[curr] {
			if visited[n] {
				continue
			}

			visited[n] = true
			count += dfs(n)
			visited[n] = false
		}

		return count
	}

	return dfs("you")
}

func main() {
	input := parse(utils.FilePath())
	fmt.Println("p1:", countPaths(input))
}
