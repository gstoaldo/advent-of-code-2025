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

func countPaths(input G, start, end string) int {
	visited := map[string]bool{}

	var dfs func(curr string) int

	dfs = func(curr string) int {
		if curr == end {
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

	return dfs(start)
}

func blocklist(input G) map[string]bool {
	result := map[string]bool{}
	for origin := range input {
		if countPaths(input, origin, "fft") == 0 && countPaths(input, origin, "dac") == 0 {
			result[origin] = true
		}
	}
	return result
}

func countPathsP2(input G, blocklist map[string]bool) int {
	visited := map[string]bool{}

	var dfs func(curr string) int

	dfs = func(curr string) int {
		if curr == "out" {
			if visited["dac"] && visited["fft"] {
				return 1
			}

			return 0
		}

		count := 0

		for _, n := range input[curr] {
			if visited[n] {
				continue
			}

			if blocklist[n] && !visited["fft"] && !visited["dac"] {
				continue
			}

			visited[n] = true
			count += dfs(n)
			visited[n] = false
		}

		return count
	}

	return dfs("svr")
}

func main() {
	input := parse(utils.FilePath())
	fmt.Println("p1:", countPaths(input, "you", "out"))
	fmt.Println("p2:", countPathsP2(input, blocklist(input)))
}
