package main

import (
	"aoc/lib/m"
	"aoc/lib/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	a, b := getLists(input)
	p1 := part1(a, b)
	p2 := part2(a, b)
	fmt.Println(p1)
	fmt.Println(p2)
}

func part1(a, b []int) int {
	res := 0
	for i, val := range a {
		res += m.Abs(val - b[i])
	}
	return res
}

func part2(a, b []int) int {
	m1 := make(map[int]bool)
	m2 := make(map[int]int)
	for _, num := range a {
		m1[num] = true
	}
	for _, num := range b {
		m2[num]++
	}
	res := 0
	for k := range m1 {
		res += k * m2[k]
	}
	return res
}

func getLists(lines []string) ([]int, []int) {
	n := len(lines)
	list1 := make([]int, n)
	list2 := make([]int, n)

	for i, line := range lines {
		parts := strings.Split(line, "   ")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		list1[i] = a
		list2[i] = b
	}

	sort.Ints(list1)
	sort.Ints(list2)
	return list1, list2
}
