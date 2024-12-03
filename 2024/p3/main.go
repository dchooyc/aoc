package main

import (
	"aoc/lib/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	p1 := part1(input)
	p2 := part2(input)
	fmt.Println(p1)
	fmt.Println(p2)
}

func part1(lines []string) int {
	res := 0
	for _, line := range lines {
		res += mull(line)
	}
	return res
}

func part2(lines []string) int {
	newLines := doDont(lines)
	return part1(newLines)
}

func doDont(lines []string) []string {
	res := []string{}
	line := strings.Join(lines, "")
	parts := strings.Split(line, "don't()")
	res = append(res, parts[0])
	for i := 1; i < len(parts); i++ {
		part := parts[i]
		lines := strings.Split(part, "do()")
		res = append(res, lines[1:]...)
	}
	return res
}

func mull(s string) int {
	res := 0
	parts := strings.Split(s, "mul(")

	for i := 1; i < len(parts); i++ {
		part := parts[i]
		close := strings.Split(part, ")")
		if len(close) < 2 {
			continue
		}
		front := close[0]
		nums := strings.Split(front, ",")
		if len(nums) != 2 {
			continue
		}
		a, err1 := strconv.Atoi(nums[0])
		b, err2 := strconv.Atoi(nums[1])
		if err1 == nil && err2 == nil {
			res += a * b
		}
	}

	return res
}
