package main

import (
	"aoc/lib/utils"
	"fmt"
)

func main() {
	input := utils.ConvFile("input.txt")
	part1 := totBatteries(input, 2)
	part2 := totBatteries(input, 12)
	fmt.Println(part1)
	fmt.Println(part2)
}

func totBatteries(input []string, b int) int {
	res := 0

	for _, line := range input {
		m := maxBattery(line, b)
		res += m
	}

	return res
}

func maxBattery(s string, l int) int {
	tally := make([][]int, 10)
	n := len(s)

	for i := 0; i < n; i++ {
		val := int(s[i] - '0')
		tally[val] = append(tally[val], i)
	}

	prevIndex := -1
	res := 0

	for i := l; i >= 1; i-- {
		vals := findMaxDigit(tally, prevIndex, n-i)
		index, digit := vals[0], vals[1]
		prevIndex = index
		res *= 10
		res += digit
	}

	return res
}

func findMaxDigit(tally [][]int, minIndex, maxIndex int) [2]int {
	for i := 9; i >= 1; i-- {
		target := -1
		if len(tally[i]) > 0 {
			for j, index := range tally[i] {
				if index < minIndex || index > maxIndex {
					continue
				}
				target = j
				break
			}
		}
		if target != -1 {
			index := tally[i][target]
			tally[i] = append(tally[i][:target], tally[i][target+1:]...)
			return [2]int{index, i}
		}
	}
	return [2]int{-1, -1}
}
