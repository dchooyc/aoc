package main

import (
	"aoc/lib/b"
	"aoc/lib/m"
	"aoc/lib/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	seeds := stia(strings.Split(input[0], ": ")[1])
	mappings := getMappings(input[3:])
	first := m.ArrMin64(locBySeeds(seeds, mappings))
	second := minLocByRanges(seeds, mappings)
	fmt.Println(first)
	fmt.Println(second)
}

func seedsToRanges(seeds []int64) [][2]int64 {
	res := [][2]int64{}

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := start + seeds[i+1]
		res = append(res, [2]int64{start, end})
	}

	return res
}

func minLocByRanges(seeds []int64, mappings [][][]int64) int64 {
	ranges := seedsToRanges(seeds)

	for _, mapping := range mappings {
		newRanges := [][2]int64{}

		for i := 0; i < len(ranges); i++ {
			rStart := ranges[i][0]
			rEnd := ranges[i][1]
			overlap := false

			for _, cores := range mapping {
				desStart := cores[0]
				srcStart := cores[1]
				srcEnd := srcStart + cores[2]
				overlapStart := m.Max64(rStart, srcStart)
				overlapEnd := m.Min64(rEnd, srcEnd)

				if overlapStart < overlapEnd {
					rDesStart := desStart + (overlapStart - srcStart)
					rDesEnd := desStart + (overlapEnd - srcStart)
					newRanges = append(newRanges, [2]int64{rDesStart, rDesEnd})

					if overlapStart > rStart {
						ranges = append(ranges, [2]int64{rStart, overlapStart})
					}

					if overlapEnd < rEnd {
						ranges = append(ranges, [2]int64{overlapEnd, rEnd})
					}

					overlap = true
					break
				}
			}

			if !overlap {
				newRanges = append(newRanges, ranges[i])
			}
		}

		ranges = newRanges
	}

	min := ranges[0][0]

	for i := 1; i < len(ranges); i++ {
		min = m.Min64(min, ranges[i][0])
	}

	return min
}

func locBySeeds(seeds []int64, mappings [][][]int64) []int64 {
	res := make([]int64, len(seeds))
	copy(res, seeds)

	for _, mapping := range mappings {
		for i, val := range res {
			for _, cores := range mapping {
				desStart := cores[0]
				srcStart := cores[1]
				srcEnd := srcStart + cores[2]

				if val >= srcStart && val < srcEnd {
					diff := val - srcStart
					dest := desStart + diff
					res[i] = dest
					break
				}
			}
		}
	}

	return res
}

func getMappings(lines []string) [][][]int64 {
	mappings := [][][]int64{}
	curMapping := [][]int64{}

	for _, line := range lines {
		if len(line) != 0 && b.IsLetter(line[0]) {
			continue
		}

		if len(line) == 0 {
			mappings = append(mappings, curMapping)
			curMapping = [][]int64{}
		} else {
			curMapping = append(curMapping, stia(line))
		}
	}

	mappings = append(mappings, curMapping)
	return mappings
}

func stia(s string) []int64 {
	res := []int64{}
	var cur int64

	for i := 0; i < len(s); i++ {
		if b.IsNum(s[i]) {
			cur *= 10
			cur += int64(s[i] - '0')
		} else {
			res = append(res, cur)
			cur = 0
		}
	}

	res = append(res, cur)
	return res
}
