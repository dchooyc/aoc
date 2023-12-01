package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := convFile("./input.txt")
	grid := genGrid(input)
	visible := countVisible(grid)
	maximum := maxScore(grid)
	fmt.Println(visible)
	fmt.Println(maximum)
}

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			res = max(res, checkScore(grid, i, j, m, n))
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countVisible(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := (m * 2) + (n * 2) - 4

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if checkCell(grid, i, j, m, n) {
				res++
			}
		}
	}

	return res
}

func checkScore(grid [][]int, i, j, m, n int) int {
	cell := grid[i][j]
	up, down, left, right := -1, -1, -1, -1
	u, d, l, r := i-1, i+1, j-1, j+1

	for ; u >= 0; u-- {
		if grid[u][j] >= cell {
			up = i - u
			break
		}
	}

	if up == -1 {
		up = i
	}

	for ; d < m; d++ {
		if grid[d][j] >= cell {
			down = d - i
			break
		}
	}

	if down == -1 {
		down = m - 1 - i
	}

	for ; l >= 0; l-- {
		if grid[i][l] >= cell {
			left = j - l
			break
		}
	}

	if left == -1 {
		left = j
	}

	for ; r < n; r++ {
		if grid[i][r] >= cell {
			right = r - j
			break
		}
	}

	if right == -1 {
		right = n - 1 - j
	}

	return up * down * left * right
}

func checkCell(grid [][]int, i, j, m, n int) bool {
	cell := grid[i][j]
	up, down, left, right := true, true, true, true
	u, d, l, r := i-1, i+1, j-1, j+1

	for ; u >= 0; u-- {
		if grid[u][j] >= cell {
			up = false
			break
		}
	}

	if up {
		return true
	}

	for ; d < m; d++ {
		if grid[d][j] >= cell {
			down = false
			break
		}
	}

	if down {
		return true
	}

	for ; l >= 0; l-- {
		if grid[i][l] >= cell {
			left = false
			break
		}
	}

	if left {
		return true
	}

	for ; r < n; r++ {
		if grid[i][r] >= cell {
			right = false
			break
		}
	}

	if right {
		return true
	}

	return false
}

func genGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))

	for i, line := range lines {
		grid[i] = make([]int, len(line))

		for j := 0; j < len(line); j++ {
			grid[i][j] = int(line[j] - '0')
		}
	}

	return grid
}

func convFile(src string) []string {
	res := []string{}

	file, err := os.Open(src)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
