package main

import (
	"aoc/lib/utils"
	"fmt"
)

func main() {
	input := utils.ConvFile("input.txt")
	first := calVal(input)
	second := calValWords(input)
	fmt.Println(first)
	fmt.Println(second)
}

func calVal(lines []string) int {
	res := 0

	for _, line := range lines {
		nums := []int{}

		for i := 0; i < len(line); i++ {
			char := line[i]

			if char >= '0' && char <= '9' {
				nums = append(nums, int(char-'0'))
			}
		}
		res += (nums[0] * 10) + nums[len(nums)-1]
	}

	return res
}

func calValWords(lines []string) int {
	res := 0
	words := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range lines {
		nums := []int{}

		for i := 0; i < len(line); i++ {
			char := line[i]

			if char >= '0' && char <= '9' {
				nums = append(nums, int(char-'0'))
			} else {
				for end := i + 3; end <= len(line) && end <= i+5; end++ {
					check := line[i:end]
					if val, ok := words[check]; ok {
						nums = append(nums, val)
						break
					}
				}
			}
		}
		res += (nums[0] * 10) + nums[len(nums)-1]
	}

	return res
}
