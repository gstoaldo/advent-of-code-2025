package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func parse(filepath string) (result [][]int) {
	lines := utils.ReadLines(filepath)

	for _, line := range lines {
		dirStr := string(line[0])
		amount := utils.ToInt(string(line[1:]))

		dir := 1
		if dirStr == "L" {
			dir = -1
		}

		result = append(result, []int{dir, amount})
	}

	return result
}

func rotate(v0, dir, amount int) (int, int) {
	laps := amount / 100
	delta := dir * amount % 100
	final := v0 + delta

	if delta != 0 && v0 != 0 && (final <= 0 || final >= 100) {
		laps++
	}

	final = final % 100
	if final < 0 {
		final += 100
	}

	return final, laps
}

func password(input [][]int) (int, int) {
	timesAtZero := 0
	totalLaps := 0

	curr := 50
	for _, step := range input {
		next, laps := rotate(curr, step[0], step[1])
		curr = next
		totalLaps += laps

		if curr == 0 {
			timesAtZero++
		}
	}

	return timesAtZero, totalLaps
}

func main() {
	input := parse(utils.FilePath())

	timesAtZero, totalLaps := password(input)

	fmt.Println("p1:", timesAtZero)
	fmt.Println("p2:", totalLaps)
}
