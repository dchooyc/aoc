package main

import (
	"aoc/lib/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	ranges := createRanges(input)
	part1 := addPatterns(ranges)
	part2 := addAnyPatterns(ranges)
	fmt.Println(part1)
	fmt.Println(part2)
}

func addAnyPatterns(ranges [][]int) int64 {
	var res int64

	for _, r := range ranges {
		start, end := r[0], r[1]
		for i := start; i <= end; i++ {
			if isAnyRepeating(i) {
				res += int64(i)
			}
		}
	}

	return res
}

func isAnyRepeating(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}
		found := true
		pattern := s[:i]
		for j := i; j < n; j += i {
			if s[j:j+i] != pattern {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}

func addPatterns(ranges [][]int) int64 {
	var res int64

	for _, r := range ranges {
		start, end := r[0], r[1]
		for i := start; i <= end; i++ {
			if isRepeating(i) {
				res += int64(i)
			}
		}
	}

	return res
}

func isRepeating(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	if n%2 != 0 {
		return false
	}
	return s[:n/2] == s[n/2:]
}

func createRanges(input []string) [][]int {
	vals := strings.Split(input[0], ",")
	ranges := [][]int{}

	for _, val := range vals {
		nums := strings.Split(val, "-")
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		ranges = append(ranges, []int{start, end})
	}

	return ranges
}
