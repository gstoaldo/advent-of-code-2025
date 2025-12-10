package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

type Machine struct {
	pattern  []bool
	buttons  [][]int
	joltages []int
}

func parse(path string) (machines []Machine) {
	patternRe := regexp.MustCompile(`\[(.+)\]`)
	buttonsRe := regexp.MustCompile(`\(([\d|,]+)\)`)
	joltageRe := regexp.MustCompile(`\{(.*)\}`)

	for _, line := range utils.ReadLines(path) {
		patternMatch := patternRe.FindStringSubmatch(line)[1]

		pattern := make([]bool, len(patternMatch))
		for i, char := range patternMatch {
			pattern[i] = char == '#'
		}

		buttonsMatch := buttonsRe.FindAllStringSubmatch(line, -1)
		buttons := make([][]int, len(buttonsMatch))
		for i, group := range buttonsMatch {
			button := []int{}
			for numStr := range strings.SplitSeq(group[1], ",") {
				button = append(button, utils.ToInt(numStr))
			}

			buttons[i] = button
		}

		joltageMatch := strings.Split(joltageRe.FindStringSubmatch(line)[1], ",")
		joltages := make([]int, len(joltageMatch))
		for i, numStr := range joltageMatch {
			joltages[i] = utils.ToInt(numStr)
		}

		machines = append(machines, Machine{
			pattern:  pattern,
			buttons:  buttons,
			joltages: joltages,
		})
	}

	return machines
}

func toggle(pattern []bool, button []int) []bool {
	result := append([]bool{}, pattern...)

	for _, b := range button {
		result[b] = !pattern[b]
	}

	return result
}

func match(state, pattern []bool) bool {
	for i := range state {
		if state[i] != pattern[i] {
			return false
		}
	}

	return true
}

func minPresses(machine Machine) (result int) {
	var run func(state []bool, btnId, count int) int

	run = func(state []bool, btnId, count int) int {
		if match(state, machine.pattern) {
			return count
		}

		if btnId == len(machine.buttons) {
			// no solution
			return math.MaxInt
		}

		// For each button, we have two choices: skip it or press it.
		// There is no point in pressing the same button more than once,
		// since pressing it twice returns to the original state.
		skip := run(state, btnId+1, count)
		press := run(toggle(state, machine.buttons[btnId]), btnId+1, count+1)

		return min(skip, press)
	}

	start := make([]bool, len(machine.pattern))

	return run(start, 0, 0)
}

func part1(machines []Machine) (result int) {
	for _, machine := range machines {
		f := minPresses(machine)

		if f == math.MaxInt {
			panic("no solution found")
		}

		result += minPresses(machine)
	}

	return result
}

func main() {
	machines := parse(utils.FilePath())
	fmt.Println("p1:", part1(machines))
}
