package main

import (
	"aoc/lib/utils"
	"fmt"
	"strconv"
)

func main() {
	input := utils.ConvFile("input.txt")
	lines := [][]byte{}

	for _, line := range input {
		lines = append(lines, []byte(line))
	}

	m, n := len(lines), len(lines[0])

	start := find(lines, 'S')
	dirs := map[string][]int{
		"u": {-1, 0},
		"d": {1, 0},
		"l": {0, -1},
		"r": {0, 1},
	}
	charDirs := map[byte][][]int{
		'F': {dirs["r"], dirs["d"]},
		'7': {dirs["l"], dirs["d"]},
		'J': {dirs["u"], dirs["l"]},
		'L': {dirs["u"], dirs["r"]},
		'|': {dirs["u"], dirs["d"]},
		'-': {dirs["l"], dirs["r"]},
	}

	points := [][]int{start}

	for sym, dir := range dirs {
		x := dir[0] + start[0]
		y := dir[1] + start[1]
		if !isValid(x, y, m, n) {
			continue
		}
		char := lines[x][y]
		a := sym == "u" && (char == 'F' || char == '7' || char == '|')
		b := sym == "d" && (char == 'J' || char == 'L' || char == '|')
		c := sym == "l" && (char == 'L' || char == 'F' || char == '-')
		d := sym == "r" && (char == '7' || char == 'J' || char == '-')
		if a || b || c || d {
			points = append(points, []int{x, y})
			break
		}
	}

	next := points[1]
	points = append(points, next)
	seen := make(map[string]bool)
	seen[getKey(start[0], start[1])] = true
	seen[getKey(next[0], next[1])] = true

	findNext(seen, next, lines, &points, charDirs)
	fmt.Println(len(points) / 2)
}

func findNext(
	seen map[string]bool,
	cur []int,
	lines [][]byte,
	points *[][]int,
	charDirs map[byte][][]int,
) {
	x, y := cur[0], cur[1]
	char := lines[x][y]
	dirs := charDirs[char]
	next := []int{}
	for _, dir := range dirs {
		a := x + dir[0]
		b := y + dir[1]
		key := getKey(a, b)
		if seen[key] {
			continue
		} else {
			seen[key] = true
			next = []int{a, b}
			break
		}
	}
	if len(next) == 0 {
		return
	}
	*points = append(*points, next)
	findNext(seen, next, lines, points, charDirs)
}

func getKey(a, b int) string {
	return strconv.Itoa(a) + "," + strconv.Itoa(b)
}

func isValid(i, j, m, n int) bool {
	a := i >= 0 && i < m
	b := j >= 0 && j < n
	return a && b
}

func find(lines [][]byte, char byte) []int {
	res := []int{}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == char {
				res = append(res, i)
				res = append(res, j)
				break
			}
		}
	}

	return res
}
