package main

import (
	"aoc/lib/m"
	"aoc/lib/utils"
	"fmt"
)

func main() {
	input := utils.ConvFile("input.txt")
	part1, part2 := beamSplits(input)
	fmt.Println(part1)
	fmt.Println(part2)
}

func beamSplits(input []string) (int, int64) {
	res := 0
	beams := make([]bool, len(input[0]))
	times := make([]int64, len(input[0]))

	for i := 0; i < len(input[0]); i++ {
		if input[0][i] == 'S' {
			beams[i] = true
			times[i] = 1
			break
		}
	}

	for i := 1; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '^' && beams[j] {
				beams[j-1] = true
				beams[j+1] = true
				beams[j] = false
				times[j-1] += times[j]
				times[j+1] += times[j]
				times[j] = 0
				res++
			}
		}
	}

	return res, m.ArrSum64(times)
}
