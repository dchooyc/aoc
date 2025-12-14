package main

import (
	"aoc/lib/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	children := createChildren(input)
	part1 := countPaths("you", "out", children)
	part2 := svrPaths(children)
	fmt.Println(part1)
	fmt.Println(part2)
}

func svrPaths(children map[string][]string) int64 {
	sf := countPaths("svr", "fft", children)
	sd := countPaths("svr", "dac", children)
	fd := countPaths("fft", "dac", children)
	df := countPaths("dac", "fft", children)
	fo := countPaths("fft", "out", children)
	do := countPaths("dac", "out", children)

	return (sf * fd * do) + (sd * df * fo)
}

func countPaths(cur, dst string, children map[string][]string) int64 {
	memo := make(map[string]int64)
	return dfs(cur, dst, children, memo)
}

func dfs(cur, dst string, children map[string][]string, memo map[string]int64) int64 {
	if cur == dst {
		return 1
	}

	if val, ok := memo[cur]; ok {
		return val
	}

	var res int64

	for _, child := range children[cur] {
		res += dfs(child, dst, children, memo)
	}

	memo[cur] = res
	return res
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
