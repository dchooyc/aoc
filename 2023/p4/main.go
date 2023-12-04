package main

import (
	"aoc/lib/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	matches := cardMatches(input)
	first := points(matches)
	second := copies(matches)
	fmt.Println(first)
	fmt.Println(second)
}

func copies(matches []int) int {
	n := len(matches)
	res, copies := 0, make([]int, n)

	for i := 0; i < n; i++ {
		copies[i]++
		for j := i + 1; j < n && j <= i+matches[i]; j++ {
			copies[j] += copies[i]
		}
		res += copies[i]
	}

	return res
}

func points(matches []int) int {
	res := 0

	for _, match := range matches {
		if match != 0 {
			res += 1 << (match - 1)
		}
	}

	return res
}

func cardMatches(lines []string) []int {
	res := make([]int, len(lines))

	for i, line := range lines {
		res[i] = getMatches(line)
	}

	return res
}

func getMatches(line string) int {
	parts := strings.Split(line, ": ")
	numbers := strings.Split(parts[1], " | ")
	winners := getNumbers(numbers[0])
	scratch := getNumbers(numbers[1])
	matches := 0

	for num := range scratch {
		if winners[num] {
			matches++
		}
	}

	return matches
}

func getNumbers(numbers string) map[int]bool {
	res, cur := make(map[int]bool), 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == ' ' {
			if cur != 0 {
				res[cur] = true
			}
			cur = 0
		} else {
			cur *= 10
			cur += int(numbers[i] - '0')
		}
	}

	if cur != 0 {
		res[cur] = true
	}

	return res
}
