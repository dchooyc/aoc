package main

import (
	"aoc/lib/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	children := createChildren(input)
	part1 := 0
	countPaths("you", "out", &part1, children)
	fmt.Println(part1)
}

func countPaths(cur, dst string, res *int, children map[string][]string) {
	if cur == dst {
		*res++
		return
	}
	for _, child := range children[cur] {
		countPaths(child, dst, res, children)
	}
}

func createChildren(input []string) map[string][]string {
	children := make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, " ")
		parent := parts[0][:len(parts[0])-1]
		children[parent] = parts[1:]
	}

	return children
}
