package main

import (
	"aoc/lib/utils"
	"fmt"
	"strconv"
)

func main() {
	input := utils.ConvFile("input.txt")
	first := sumParts(input)
	second := sumGears(input)
	fmt.Println(first)
	fmt.Println(second)
}

func sumParts(lines []string) int {
	res := 0

	for i, line := range lines {
		cur, adjacent := 0, false
		for j := 0; j < len(line); j++ {
			char := line[j]
			if isNum(char) {
				cur *= 10
				cur += int(line[j] - '0')
				if adjSymbol(lines, i, j) {
					adjacent = true
				}
			} else {
				if cur != 0 && adjacent {
					res += cur
				}
				cur, adjacent = 0, false
			}
		}
		if cur != 0 && adjacent {
			res += cur
		}
		cur, adjacent = 0, false
	}

	return res
}

func adjSymbol(lines []string, i, j int) bool {
	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}

	for _, dir := range dirs {
		x, y := i+dir[0], j+dir[1]
		if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[0]) {
			char := lines[x][y]
			if char != '.' && !isNum(char) {
				return true
			}
		}
	}

	return false
}

func isNum(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func sumGears(lines []string) int {
	gearToParts := make(map[string][]int)

	for i, line := range lines {
		cur, adjacentGears := 0, make(map[string]bool)
		for j := 0; j < len(line); j++ {
			char := line[j]
			if isNum(char) {
				cur *= 10
				cur += int(line[j] - '0')
				ag := adjGears(lines, i, j)
				for _, g := range ag {
					adjacentGears[g] = true
				}
			} else {
				if cur != 0 && len(adjacentGears) != 0 {
					for g, _ := range adjacentGears {
						gearToParts[g] = append(gearToParts[g], cur)
					}
				}
				cur, adjacentGears = 0, make(map[string]bool)
			}
		}
		if cur != 0 && len(adjacentGears) != 0 {
			for g, _ := range adjacentGears {
				gearToParts[g] = append(gearToParts[g], cur)
			}
		}
		cur, adjacentGears = 0, make(map[string]bool)
	}

	res := 0

	for _, parts := range gearToParts {
		if len(parts) == 2 {
			res += parts[0] * parts[1]
		}
	}

	return res
}

func adjGears(lines []string, i, j int) []string {
	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}

	gears := []string{}

	for _, dir := range dirs {
		x, y := i+dir[0], j+dir[1]
		if x >= 0 && x < len(lines) && y >= 0 && y < len(lines[0]) {
			char := lines[x][y]
			if char == '*' {
				gears = append(gears, strconv.Itoa(x)+","+strconv.Itoa(y))
			}
		}
	}

	return gears
}
