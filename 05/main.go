package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func parse(path string) (intervals [][]int, ids []int) {
	chunks := strings.Split(utils.ReadFile(path), "\n\n")
	rangeChunk, idChunk := chunks[0], chunks[1]

	for line := range strings.SplitSeq(rangeChunk, "\n") {
		parts := strings.Split(line, "-")
		intervals = append(intervals, []int{utils.ToInt(parts[0]), utils.ToInt(parts[1])})
	}

	for line := range strings.SplitSeq(idChunk, "\n") {
		ids = append(ids, utils.ToInt(line))
	}

	return intervals, ids
}

func countFresh(intervals [][]int, ids []int) (result int) {
	for _, id := range ids {
		for _, interval := range intervals {
			if id >= interval[0] && id <= interval[1] {
				result++
				break
			}
		}
	}

	return result
}

func main() {
	intervals, ids := parse(utils.FilePath())
	fmt.Println("p1:", countFresh(intervals, ids))
}
