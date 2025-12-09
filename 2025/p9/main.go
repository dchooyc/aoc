package main

import (
	"aoc/lib/m"
	"aoc/lib/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	points := convertPoints(input)
	part1 := maxRectangle(points)
	fmt.Println(part1)
}

func maxRectangle(points [][]int) int {
	res := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			res = m.Max(res, rectSize(points[i], points[j]))
		}
	}

	return res
}

func rectSize(a, b []int) int {
	x1, y1 := a[0], a[1]
	x2, y2 := b[0], b[1]
	diff1 := m.Abs(x1-x2) + 1
	diff2 := m.Abs(y1-y2) + 1
	return diff1 * diff2
}

func convertPoints(input []string) [][]int {
	res := [][]int{}

	for _, line := range input {
		coords := strings.Split(line, ",")
		r, _ := strconv.Atoi(coords[1])
		c, _ := strconv.Atoi(coords[0])
		res = append(res, []int{r, c})
	}

	return res
}
