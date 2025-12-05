package main

import (
	"aoc/lib/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	ranges, ingredients := rangesAndIngredients(input)
	ranges = tidyRanges(ranges)
	part1 := cntFreshIngredients(ingredients, ranges)
	part2 := possibleFresh(ranges)
	fmt.Println(part1)
	fmt.Println(part2)
}

func possibleFresh(ranges [][]int64) int64 {
	var res int64
	for _, r := range ranges {
		s, e := r[0], r[1]
		res += e - s + 1
	}
	return res
}

func tidyRanges(ranges [][]int64) [][]int64 {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	res := [][]int64{}
	start, end := ranges[0][0], ranges[0][1]

	for i := 1; i < len(ranges); i++ {
		cur := ranges[i]
		curStart, curEnd := cur[0], cur[1]

		if curStart > end {
			res = append(res, []int64{start, end})
			start, end = curStart, curEnd
		} else if curStart >= start && curStart <= end && curEnd > end {
			end = curEnd
		}
	}

	res = append(res, []int64{start, end})

	return res
}

func cntFreshIngredients(ingredients []int64, ranges [][]int64) int {
	res := 0
	for _, ingredient := range ingredients {
		if isFresh(ranges, ingredient) {
			res++
		}
	}
	return res
}

func isFresh(ranges [][]int64, ingredient int64) bool {
	for _, r := range ranges {
		start, end := r[0], r[1]
		if ingredient >= start && ingredient <= end {
			return true
		}
	}
	return false
}

func rangesAndIngredients(input []string) ([][]int64, []int64) {
	ranges := [][]int64{}
	ingredients := []int64{}
	change := false

	for _, line := range input {
		if line == "" {
			change = true
		}
		if change {
			val, _ := strconv.Atoi(line)
			ingredients = append(ingredients, int64(val))
		} else {
			vals := strings.Split(line, "-")
			start, _ := strconv.Atoi(vals[0])
			end, _ := strconv.Atoi(vals[1])
			ranges = append(ranges, []int64{int64(start), int64(end)})
		}
	}

	return ranges, ingredients
}
