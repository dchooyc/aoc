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
	levels := getLevels(input)
	p1 := part1(levels)
	p2 := part2(levels)
	fmt.Println(p1)
	fmt.Println(p2)
}

func part2(levels [][]int) int {
	res := 0
	for _, level := range levels {
		n := len(level)
		for i := 0; i < n; i++ {
			remLevel := make([]int, n-1)
			copy(remLevel, level[:i])
			copy(remLevel[i:], level[i+1:])
			if safe(remLevel) {
				res++
				break
			}
		}
	}
	return res
}

func part1(levels [][]int) int {
	res := 0
	for _, level := range levels {
		if safe(level) {
			res++
		}
	}
	return res
}

func getLevels(lines []string) [][]int {
	n := len(lines)
	res := make([][]int, n)
	for i, line := range lines {
		vals := strings.Split(line, " ")
		nums := make([]int, len(vals))
		for key, val := range vals {
			num, _ := strconv.Atoi(val)
			nums[key] = num
		}
		res[i] = nums
	}
	return res
}

func safe(nums []int) bool {
	isSafe := true
	inc := nums[0] < nums[1]
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		absDiff := m.Abs(diff)
		if (diff > 0 && !inc) || (diff < 0 && inc) || absDiff < 1 || absDiff > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}
