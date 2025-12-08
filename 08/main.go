package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

func parse(path string) (positions [][]int) {
	for _, line := range utils.ReadLines(path) {
		position := make([]int, 3)
		for i, numStr := range strings.Split(line, ",") {
			position[i] = utils.ToInt(numStr)
		}
		positions = append(positions, position)
	}

	return positions
}

func dist(a, b []int) float64 {
	squaredSum := 0
	for i := range a {
		squaredSum += (a[i] - b[i]) * (a[i] - b[i])
	}

	return math.Sqrt(float64(squaredSum))
}

func circuitProduct(positions [][]int, kmax int) int {
	type pair struct {
		i, j int
		dist float64
	}

	dists := []pair{}

	for i := 0; i < len(positions)-1; i++ {
		for j := i + 1; j < len(positions); j++ {
			dists = append(dists, pair{i, j, dist(positions[i], positions[j])})
		}
	}

	sort.Slice(dists, func(i, j int) bool { return dists[i].dist < dists[j].dist })

	circuits := map[int]int{}
	circuitId := 1

	for k, pair := range dists {
		if k >= kmax {
			break
		}

		idI, mergedI := circuits[pair.i]
		idJ, mergedJ := circuits[pair.j]

		if mergedI && mergedJ && idI == idJ {
			// Same circuit, nothing happens
			continue
		}

		if mergedI && mergedJ {
			// Merge two circuits. Pick circuit I and update circuit J ID
			for k, v := range circuits {
				if v == idJ {
					circuits[k] = idI
				}
			}
		}

		if !mergedI && !mergedJ {
			// create new circuit
			circuits[pair.i] = circuitId
			circuits[pair.j] = circuitId
			circuitId++
		}

		if !mergedI {
			circuits[pair.i] = circuits[pair.j]
		}

		if !mergedJ {
			circuits[pair.j] = circuits[pair.i]
		}
	}

	sizesMap := map[int]int{}
	for _, v := range circuits {
		sizesMap[v]++
	}

	sizes := []int{}
	for _, v := range sizesMap {
		sizes = append(sizes, v)
	}

	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })

	result := 1
	for _, v := range sizes[:3] {
		result *= v
	}

	return result
}

func main() {
	positions := parse(utils.FilePath())
	fmt.Println("p1:", circuitProduct(positions, 1000))
}
