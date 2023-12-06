package main

import (
	"aoc/lib/b"
	"aoc/lib/utils"
	"fmt"
)

func main() {
	input := utils.ConvFile("input.txt")
	first := beatMany(input)
	second := beatOne(input)
	fmt.Println(first)
	fmt.Println(second)
}

func beatOne(lines []string) int64 {
	time := getNum(lines[0])
	dist := getNum(lines[1])
	var cant, i, j int64 = 0, 1, time

	for ; i < time; i++ {
		if i*(time-i) <= dist {
			cant++
		} else {
			break
		}
	}

	for ; j >= 0; j-- {
		if j*(time-j) <= dist {
			cant++
		} else {
			break
		}
	}

	return time - cant
}

func beatMany(lines []string) int {
	times := getNums(lines[0])
	dists := getNums(lines[1])
	res := 1

	for i := 0; i < len(times); i++ {
		time := times[i]
		cnt := 0

		for t := 1; t < time; t++ {
			if t*(time-t) > dists[i] {
				cnt++
			}
		}
		res *= cnt
	}

	return res
}

func getNum(line string) int64 {
	var res int64

	for i := 0; i < len(line); i++ {
		if b.IsNum(line[i]) {
			res *= 10
			res += int64(line[i] - '0')
		}
	}

	return res
}

func getNums(line string) []int {
	res, cur := []int{}, 0

	for i := 0; i < len(line); i++ {
		if b.IsNum(line[i]) {
			cur *= 10
			cur += int(line[i] - '0')
		} else {
			if cur != 0 {
				res = append(res, cur)
				cur = 0
			}
		}
	}

	res = append(res, cur)
	return res
}
