package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func parse(path string) (result [][]int) {
	file := utils.ReadFile(path)

	for interval := range strings.SplitSeq(file, ",") {
		parts := strings.Split(interval, "-")
		start, end := utils.ToInt(parts[0]), utils.ToInt(parts[1])
		result = append(result, []int{start, end})
	}

	return result
}

func isInvalidP1(n int) bool {
	nStr := strconv.Itoa(n)

	if len(nStr) == 1 || len(nStr)%2 != 0 {
		return false
	}

	return nStr[:len(nStr)/2] == nStr[len(nStr)/2:]
}

func isInvalidP2(n int) bool {
	nStr := strconv.Itoa(n)

	if len(nStr) == 1 {
		return false
	}

	for width := 1; width <= len(nStr)/2; width++ {
		for x := width; x <= len(nStr)-width; x += width {
			left := nStr[x-width : x]
			right := nStr[x : x+width]

			if left == right && x+width == len(nStr) {
				return true
			}

			if left != right {
				break
			}
		}
	}

	return false
}

func countInvalid(input [][]int, isInvalid func(n int) bool) (result int) {
	for _, interval := range input {
		for n := interval[0]; n <= interval[1]; n++ {
			if isInvalid(n) {
				result += n
			}
		}
	}

	return result
}

func main() {
	input := parse(utils.FilePath())
	fmt.Println("p1:", countInvalid(input, isInvalidP1))
	fmt.Println("p2:", countInvalid(input, isInvalidP2))
}
