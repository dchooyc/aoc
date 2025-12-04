package main

import (
	"aoc/lib/utils"
	"fmt"
)

func main() {
	input := utils.ConvFile("input.txt")
	part1 := accessRolls(input)
	part2 := accessAllRolls(input)
	fmt.Println(part1)
	fmt.Println(part2)
}

func accessAllRolls(input []string) int {
	res := 0
	neighbours := populateNeighbours(input)

	for q := populateQueue(neighbours); len(q) != 0; {
		res += len(q)
		for i := 0; i < len(q); i++ {
			node := q[i]
			for nodeNeighbour := range neighbours[node] {
				delete(neighbours[nodeNeighbour], node)
			}
			delete(neighbours, node)
		}
		q = populateQueue(neighbours)
	}

	return res
}

func populateQueue(neighbours map[int]map[int]bool) []int {
	res := []int{}
	for node, nodeNeighbours := range neighbours {
		if len(nodeNeighbours) < 4 {
			res = append(res, node)
		}
	}
	return res
}

func populateNeighbours(input []string) map[int]map[int]bool {
	neighbours := make(map[int]map[int]bool)
	m, n := len(input), len(input[0])

	for a := 0; a < m; a++ {
		for b := 0; b < n; b++ {
			if input[a][b] == '@' {
				neighbours[nodeVal(a, b, n)] = make(map[int]bool)
			}
		}
	}

	for i, line := range input {
		for j, char := range line {
			if char == '@' {
				cur := nodeVal(i, j, n)
				for x := i - 1; x <= i+1; x++ {
					for y := j - 1; y <= j+1; y++ {
						if !isValid(x, y, m, n) || (x == i && y == j) {
							continue
						}
						if input[x][y] == '@' {
							node := nodeVal(x, y, n)
							neighbours[cur][node] = true
							neighbours[node][cur] = true
						}
					}
				}
			}
		}
	}

	return neighbours
}

func nodeVal(i, j, n int) int {
	return (i * n) + j
}

func accessRolls(input []string) int {
	res := 0

	for i, line := range input {
		for j, char := range line {
			if char == '@' {
				val := surroundValue(input, i, j)
				if val < 4 {
					res++
				}
			}
		}
	}

	return res
}

func surroundValue(grid []string, x, y int) int {
	res := 0
	m, n := len(grid), len(grid[0])

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !isValid(i, j, m, n) || (i == x && j == y) {
				continue
			}
			if grid[i][j] == '@' {
				res++
			}
		}
	}

	return res
}

func isValid(i, j, m, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}
