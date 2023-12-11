package main

import (
	"aoc/lib/utils"
	"fmt"
)

func main() {
	input := utils.ConvFile("input.txt")
	first := oasis(input, false)
	second := oasis(input, true)
	fmt.Println(first)
	fmt.Println(second)
}

func oasis(lines []string, backwards bool) int64 {
	var res int64

	for _, line := range lines {
		nums := utils.Stia64(line)
		extr := extrapolate(nums, backwards)

		if backwards {
			res += nums[0] - extr[0]
		} else {
			res += nums[len(nums)-1] + extr[len(extr)-1]
		}
	}

	return res
}

func extrapolate(nums []int64, backwards bool) []int64 {
	res := []int64{}
	allZero := true

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		res = append(res, diff)
		if diff != 0 {
			allZero = false
		}
	}

	if allZero {
		return res
	}

	extr := extrapolate(res, backwards)

	if backwards {
		front := res[0] - extr[0]
		newRes := make([]int64, len(res)+1)
		copy(newRes[1:], res)
		newRes[0] = front
		return newRes
	}

	back := res[len(res)-1] + extr[len(extr)-1]
	res = append(res, back)

	return res
}
