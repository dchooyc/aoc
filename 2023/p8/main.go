package main

import (
	"aoc/lib/m"
	"aoc/lib/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	instruct := input[0]
	nodes := getNodes(input[2:])
	first := steps("AAA", "ZZZ", instruct, nodes)
	second := ghosts(instruct, nodes)
	fmt.Println(first)
	fmt.Println(second)
}

func getNodes(lines []string) map[string][]string {
	nodes := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		node := parts[0]
		nodes[node] = strings.Split(parts[1][1:len(parts[1])-1], ", ")
	}

	return nodes
}

func steps(point, target, instruct string, nodes map[string][]string) int {
	steps, index := 0, 0

	for point != target {
		if instruct[index] == 'L' {
			point = nodes[point][0]
		} else {
			point = nodes[point][1]
		}
		steps++
		index++
		index %= len(instruct)
	}

	return steps
}

func end_steps(point, target, instruct string, nodes map[string][]string) int {
	steps, index := 0, 0

	for string(point[2]) != target {
		if instruct[index] == 'L' {
			point = nodes[point][0]
		} else {
			point = nodes[point][1]
		}
		steps++
		index++
		index %= len(instruct)
	}

	return steps
}

func ghosts(instruct string, nodes map[string][]string) int64 {
	points := []string{}

	for node := range nodes {
		if node[2] == 'A' {
			points = append(points, node)
		}
	}

	steps := make([]int, len(points))

	for i, point := range points {
		steps[i] = end_steps(point, "Z", instruct, nodes)
	}

	res := int64(steps[0])

	for i := 1; i < len(steps); i++ {
		res = m.Lcm64(res, int64(steps[i]))
	}

	return res
}
