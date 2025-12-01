package main

import (
	"aoc/lib/utils"
	"fmt"
	"strconv"
)

func main() {
	input := utils.ConvFile("input.txt")
	part1 := numZeroes(input)
	part2 := numPassZeroes(input)
	fmt.Println(part1)
	fmt.Println(part2)
}

func numPassZeroes(rots []string) int {
	cur := 50
	res := 0

	for _, rot := range rots {
		dir := rot[0]
		dis, _ := strconv.Atoi(rot[1:])
		res += dis / 100
		dis %= 100

		if dis == 0 {
			continue
		}

		if dir == 'L' {
			if dis >= cur && cur != 0 {
				res++
			}
			cur += 100
			cur -= dis
		} else {
			cur += dis
			if cur >= 100 {
				res++
			}
		}

		cur %= 100
	}

	return res
}

func numZeroes(rots []string) int {
	cur := 50
	res := 0

	for _, rot := range rots {
		dir := rot[0]
		dis, _ := strconv.Atoi(rot[1:])

		if dir == 'L' {
			cur += 100
			cur -= dis
		} else {
			cur += dis
		}

		cur %= 100

		if cur == 0 {
			res++
		}
	}

	return res
}
