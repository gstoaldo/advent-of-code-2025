package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

type Machine struct {
	pattern []bool
	buttons [][]int
	joltage []int
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
		joltage := make([]int, len(joltageMatch))
		for i, numStr := range joltageMatch {
			joltage[i] = utils.ToInt(numStr)
		}

		machines = append(machines, Machine{
			pattern: pattern,
			buttons: buttons,
			joltage: joltage,
		})
	}

	return machines
}

func toggle(state []bool, button []int) []bool {
	result := append([]bool{}, state...)

	for _, b := range button {
		result[b] = !state[b]
	}

	return result
}

func matchPattern(state, pattern []bool) bool {
	for i := range state {
		if state[i] != pattern[i] {
			return false
		}
	}

	return true
}

func matchJoltage(state, joltage []int) bool {
	for i := range state {
		if state[i] != joltage[i] {
			return false
		}
	}

	return true
}

func minPressesPattern(machine Machine) (result int) {
	var run func(state []bool, btnId, count int) int

	run = func(state []bool, btnId, count int) int {
		if matchPattern(state, machine.pattern) {
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

func minPressesJoltage(machine Machine) (result int) {
	keyFunc := func(btnId int, state []int) string {
		return fmt.Sprintf("%v:%v", btnId, state)
	}

	cache := map[string]int{}

	var run func(state []int, btnId, count int) int

	run = func(state []int, btnId, count int) int {
		key := keyFunc(btnId, state)

		if val, ok := cache[key]; ok && val <= count {
			return val
		}

		if matchJoltage(state, machine.joltage) {
			cache[key] = count
			return count
		}

		if btnId == len(machine.buttons) {
			// no solution
			return math.MaxInt
		}

		// For each button, we have two choices: skip it or press it.
		skip := run(state, btnId+1, count)

		// If pressing the button would cause any counter to overflow, then there is no
		// point in pressing that button
		for _, b := range machine.buttons[btnId] {
			if (state[b] + 1) > machine.joltage[b] {
				return skip
			}
		}

		// Press but don't go to next btnId as the same button can be pressed more than once.
		press := run(increaseJoltage(state, machine.buttons[btnId]), btnId, count+1)

		result := min(skip, press)
		cache[key] = result

		return result
	}

	start := make([]int, len(machine.pattern))

	return run(start, 0, 0)
}

func solve(machines []Machine, minFunc func(Machine) int) (result int) {
	for i, machine := range machines {
		f := minPressesPattern(machine)

		if f == math.MaxInt {
			panic("no solution found")
		}

		fmt.Println("solved:", i, len(machines))

		result += minFunc(machine)
	}

	return result
}

func increaseJoltage(state []int, button []int) []int {
	result := append([]int{}, state...)
	for _, b := range button {
		result[b]++
	}

	return result
}

func main() {
	machines := parse(utils.FilePath())
	fmt.Println("p1:", solve(machines, minPressesPattern))
	fmt.Println("p2:", solve(machines, minPressesJoltage))
}
