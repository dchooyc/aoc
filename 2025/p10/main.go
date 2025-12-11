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
	part1 := totPresses(input)
	fmt.Println(part1)
}

func totPresses(input []string) int {
	res := 0
	for _, line := range input {
		res += minPresses(line)
	}
	return res
}

func minPresses(line string) int {
	vals := strings.Split(line, " ")
	n := len(vals)
	lights, buttons := convLights(vals[0]), convButtons(vals[1:n-1])
	res := len(lights)
	dfs(make([]int, len(lights)), lights, buttons, 0, &res)
	return res
}

func dfs(cur, target []int, buttons [][]int, presses int, res *int) {
	if arrEqual(cur, target) {
		*res = m.Min(*res, presses)
		return
	}
	for i := 0; i < len(buttons); i++ {
		b := buttons[i]
		dfs(setButton(cur, b), target, buttons[i+1:], presses+1, res)
	}
}

func setButton(cur, b []int) []int {
	res := make([]int, len(cur))
	copy(res, cur)

	for _, pos := range b {
		res[pos] ^= 1
	}

	return res
}

func arrEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func convButtons(vals []string) [][]int {
	res := make([][]int, len(vals))

	for i, val := range vals {
		n := len(val)
		digits := strings.Split(val[1:n-1], ",")
		cur := make([]int, len(digits))

		for i, digit := range digits {
			num, _ := strconv.Atoi(digit)
			cur[i] = num
		}

		res[i] = cur
	}

	return res
}

func convLights(s string) []int {
	res := make([]int, len(s)-2)

	for i := 1; i < len(s)-1; i++ {
		if s[i] == '#' {
			res[i-1] = 1
		} else {
			res[i-1] = 0
		}
	}

	return res
}
