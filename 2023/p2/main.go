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
	first := possible(input)
	second := power(input)
	fmt.Println(first)
	fmt.Println(second)
}

func possible(lines []string) int {
	res := 0
	expected := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for i, line := range lines {
		cubes, passed := getCubes(line), true

		for cube, cnt := range cubes {
			if cnt > expected[cube] {
				passed = false
				break
			}
		}

		if passed {
			res += (i + 1)
		}
	}

	return res
}

func power(lines []string) int {
	res := 0

	for _, line := range lines {
		cubes := getCubes(line)
		pow := 1

		for _, cnt := range cubes {
			pow *= cnt
		}

		res += pow
	}

	return res
}

func getCubes(line string) map[string]int {
	parts := strings.Split(line, ": ")
	handfuls := strings.Split(parts[1], "; ")
	cubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, handful := range handfuls {
		colours := strings.Split(handful, ", ")

		for _, colour := range colours {
			vals := strings.Split(colour, " ")
			cnt, err := strconv.Atoi(vals[0])
			if err != nil {
				fmt.Println(err)
			}
			col := vals[1]
			cubes[col] = m.Max(cubes[col], cnt)
		}
	}

	return cubes
}
