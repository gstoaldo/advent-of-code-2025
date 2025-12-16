package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2025/utils"
)

type Shape [][]bool

func (s Shape) Area() int {
	result := 0
	for _, row := range s {
		for _, val := range row {
			if val {
				result++
			}
		}
	}

	return result
}

type RegionData struct {
	H        int
	W        int
	ShapeQty []int
}

func (r RegionData) Area() int {
	return r.H * r.W
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

func solve(shapes []Shape, regionsData []RegionData) int {
	impossible := 0

	for _, regionData := range regionsData {
		totalShapeArea := 0
		for i, qty := range regionData.ShapeQty {
			totalShapeArea += qty * shapes[i].Area()
		}

		if totalShapeArea > regionData.Area() {
			impossible++
		}
	}

	return len(regionsData) - impossible
}

func main() {
	fmt.Println("p1:", solve(parse(utils.FilePath())))
}
