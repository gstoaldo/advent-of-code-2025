package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

type Shape [][]bool
type RegionData struct {
	H        int
	W        int
	ShapeQty []int
}

func parse(path string) (shapes []Shape, regionsData []RegionData) {
	chunks := strings.Split(utils.ReadFile(path), "\n\n")
	shapesChunks, regionsChunk := chunks[:len(chunks)-1], chunks[len(chunks)-1]

	for _, shapeChunk := range shapesChunks {
		lines := strings.Split(shapeChunk, "\n")
		shape := Shape{}
		for _, line := range lines[1:] {
			row := make([]bool, len(line))
			for i, c := range line {
				row[i] = c == '#'
			}
			shape = append(shape, row)
		}

		shapes = append(shapes, shape)
	}

	for line := range strings.SplitSeq(regionsChunk, "\n") {
		splitted := strings.Split(line, ": ")
		size := strings.Split(splitted[0], "x")
		W, H := size[0], size[1]
		qty := []int{}
		for numStr := range strings.SplitSeq(splitted[1], " ") {
			qty = append(qty, utils.ToInt(numStr))
		}

		regionsData = append(regionsData, RegionData{
			W:        utils.ToInt(W),
			H:        utils.ToInt(H),
			ShapeQty: qty,
		})
	}

	return shapes, regionsData
}

func rotate(shape [][]bool) [][]bool {
	H, W := len(shape), len(shape[0])

	result := make([][]bool, W)
	for i := range result {
		result[i] = make([]bool, H)
	}

	for i, row := range shape {
		for j, val := range row {
			result[j][W-i-1] = val
		}
	}

	return result
}

func flipH(shape [][]bool) [][]bool {
	H, W := len(shape), len(shape[0])

	result := make([][]bool, H)
	for i := range result {
		result[i] = make([]bool, W)
	}

	for i, row := range shape {
		copy(result[H-i-1], row)
	}

	return result
}

func flipV(shape [][]bool) [][]bool {
	H, W := len(shape), len(shape[0])

	result := make([][]bool, H)
	for i := range result {
		result[i] = make([]bool, W)
	}

	for i, row := range shape {
		for j, val := range row {
			result[i][W-j-1] = val
		}
	}

	return result
}

func print(shape [][]bool) {
	for _, line := range shape {
		fmt.Println(line)
	}
	fmt.Println("---")
}

func newRegion(regionData RegionData) [][]bool {
	region := make([][]bool, regionData.H)
	for i := range regionData.H {
		region[i] = make([]bool, regionData.W)
	}

	return region
}

func canAllocate(region [][]bool, shape Shape, di, dj int) bool {
	for si, line := range shape {
		for sj, val := range line {
			if !val {
				continue
			}

			i, j := si+di, sj+dj

			if i >= len(region) || j >= len(region[0]) {
				return false
			}

			if region[i][j] {
				return false
			}
		}
	}
	return true
}

func allocate(region [][]bool, shape Shape, di, dj int) {
	for si, line := range shape {
		for sj, val := range line {
			i, j := si+di, sj+dj
			if val {
				region[i][j] = val
			}
		}
	}
}

func deallocate(region [][]bool, shape Shape, di, dj int) {
	for si, line := range shape {
		for sj, val := range line {
			i, j := si+di, sj+dj
			if val {
				region[i][j] = !val
			}
		}
	}
}

func canFitAll(shapes []Shape, regionData RegionData) bool {
	region := newRegion(regionData)

	var run func(id int, qty int) bool

	run = func(id int, qty int) bool {
		if id == len(regionData.ShapeQty)-1 && qty == 0 {
			return true
		}

		if qty == 0 {
			return run(id+1, regionData.ShapeQty[id+1])
		}

		fmt.Println(id, qty)

		for di, row := range region {
			for dj := range row {
				original := shapes[id]
				shapes := []Shape{}

				for range 4 {
					roundShapes := []Shape{
						rotate(original),
						flipH(original),
						flipV(original),
					}
					shapes = append(shapes, roundShapes...)
					original = rotate(original)
				}

				for _, shape := range shapes {
					if !canAllocate(region, shape, di, dj) {
						continue
					}

					allocate(region, shape, di, dj)
					can := run(id, qty-1)
					if can {
						return can
					}
					deallocate(region, shape, di, dj)
				}
			}
		}
		return false
	}

	return run(0, regionData.ShapeQty[0])
}

func part1(shapes []Shape, regionsData []RegionData) (result int) {
	for _, regionData := range regionsData {
		if canFitAll(shapes, regionData) {
			result++
		}
	}
	return result
}

func main() {
	shapes, regionsData := parse(utils.FilePath())
	fmt.Println("p1:", part1(shapes, regionsData))
}
