package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func parse(path string) []string {
	return utils.ReadLines(path)
}

func maxJoltageP1(s string) int {
	// Traverse the string from left to right (up to the punultimate digit) and pick
	// the digit with the highest value.
	// Then, traverse the string from right to left (up to the picked left pointer)
	// and pick the digit with highest value.
	maxL := 0
	maxR := len(s) - 1

	for i := 1; i < len(s)-1; i++ {
		if utils.ToInt(string(s[i])) > utils.ToInt(string(s[maxL])) {
			maxL = i
		}
	}

	for i := len(s) - 2; i > maxL; i-- {
		if utils.ToInt(string(s[i])) > utils.ToInt(string(s[maxR])) {
			maxR = i
		}
	}

	return utils.ToInt(string(s[maxL]) + string(s[maxR]))
}

func maxJoltageP2(s string) int {
	// Start by considering a window with the last 12 digits. Then, look to the left and
	// get the digit with the highest value. The next iteration can look left up to the previous picked digit.

	digits := ""
	maxL := 0

	for r := len(s) - 12; r < len(s); r++ {

		for l := maxL; l <= r; l++ {
			if utils.ToInt(string(s[l])) > utils.ToInt(string(s[maxL])) {
				maxL = l
			}
		}

		digits += string(s[maxL])
		maxL++
	}

	return utils.ToInt(digits)
}

func sumMaxJoltage(input []string, maxJoltage func(s string) int) (result int) {
	for _, bank := range input {
		result += maxJoltage(bank)
	}

	return result
}

func main() {
	input := parse(utils.FilePath())
	fmt.Println("p1:", sumMaxJoltage(input, maxJoltageP1))
	fmt.Println("p2:", sumMaxJoltage(input, maxJoltageP2))
}
