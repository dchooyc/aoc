package main

import (
	"aoc/lib/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	junctions := makeJunctions(input)
	pairs := sortedPairs(junctions)
	neighbours := createNeighbours(pairs)
	part1 := threeLargest(circuitSizes(neighbours))
	part2 := lastTwo(junctions, pairs, neighbours)
	fmt.Println(part1)
	fmt.Println(part2)
}

func lastTwo(junctions, pairs [][]int64, neighbours map[int][]int) int64 {
	for i := 1000; i < len(pairs); i++ {
		p := pairs[i]
		a, b := int(p[0]), int(p[1])
		neighbours[a] = append(neighbours[a], b)
		neighbours[b] = append(neighbours[b], a)
		seen := make(map[int]bool)
		v := visit(0, seen, neighbours)
		if v == 1000 {
			return junctions[a][0] * junctions[b][0]
		}
	}
	return 0
}

func threeLargest(circuits []int) int {
	n := len(circuits)
	a := circuits[n-1]
	b := circuits[n-2]
	c := circuits[n-3]
	return a * b * c
}

func circuitSizes(neighbours map[int][]int) []int {
	res := []int{}
	seen := make(map[int]bool)

	for i := 0; i < 1000; i++ {
		if !seen[i] {
			res = append(res, visit(i, seen, neighbours))
		}
	}

	sort.Ints(res)
	return res
}

func visit(node int, seen map[int]bool, neighbours map[int][]int) int {
	seen[node] = true
	res := 1

	for _, neighbour := range neighbours[node] {
		if !seen[neighbour] {
			res += visit(neighbour, seen, neighbours)
		}
	}

	return res
}

func createNeighbours(pairs [][]int64) map[int][]int {
	neighbours := make(map[int][]int)

	for i := 0; i < 1000; i++ {
		p := pairs[i]
		a, b := int(p[0]), int(p[1])

		neighbours[a] = append(neighbours[a], b)
		neighbours[b] = append(neighbours[b], a)
	}

	return neighbours
}

func sortedPairs(junctions [][]int64) [][]int64 {
	res := [][]int64{}
	n := len(junctions)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := getDistance(junctions[i], junctions[j])
			res = append(res, []int64{int64(i), int64(j), d})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i][2] < res[j][2]
	})

	return res
}

func getDistance(p1, p2 []int64) int64 {
	var res int64
	for i := 0; i < len(p1); i++ {
		a, b := p1[i], p2[i]
		res += (a - b) * (a - b)
	}
	return res
}

func makeJunctions(input []string) [][]int64 {
	res := [][]int64{}

	for _, line := range input {
		vals := strings.Split(line, ",")
		nums := []int64{}
		for _, val := range vals {
			num, _ := strconv.Atoi(val)
			nums = append(nums, int64(num))
		}
		res = append(res, nums)
	}

	return res
}
