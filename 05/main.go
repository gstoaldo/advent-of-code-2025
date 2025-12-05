package main

import (
	"fmt"
	"sort"
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

func overlaps(a, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]
}

func mergeIntervals(intervals [][]int) (result [][]int) {
	// sort intervals by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	curr := intervals[0]

	for i := 1; i < len(intervals); i++ {
		next := intervals[i]

		// If the two intervals overlap, merge them into a single interval
		// spanning from the minimum start to the maximum end.
		if overlaps(curr, next) {
			curr = []int{min(curr[0], next[0]), max(curr[1], next[1])}
			continue
		}

		result = append(result, curr)
		curr = next
	}

	result = append(result, curr)

	return result
}

func countFreshIntervals(intervals [][]int) (result int) {
	merged := mergeIntervals(intervals)

	for _, interval := range merged {
		result += interval[1] - interval[0] + 1
	}

	return result
}

func main() {
	intervals, ids := parse(utils.FilePath())
	fmt.Println("p1:", countFresh(intervals, ids))
	fmt.Println("p2:", countFreshIntervals(intervals))
}
