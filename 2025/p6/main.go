package main

import (
	"aoc/lib/utils"
	"fmt"
)

func main() {
	input := utils.ConvFile("input.txt")
	vals, syms := valsAndSyms(input)
	part1 := grandTotal(vals, syms)
	part2 := cephMath(input)
	fmt.Println(part1)
	fmt.Println(part2)
}

func cephMath(input []string) int64 {
	var res int64
	vals := []int{}

	for i := len(input[0]) - 1; i >= 0; i-- {
		val := 0

		for j := 0; j < 4; j++ {
			char := input[j][i]
			if char != ' ' {
				val *= 10
				val += int(char - '0')
			}
		}

		vals = append(vals, val)
		sym := input[4][i]

		if sym != ' ' {
			if sym == '*' {
				res += int64(multi(vals))
			} else {
				res += int64(sum(vals))
			}
			vals = []int{}
			i--
		}
	}

	return res
}

func multi(vals []int) int {
	res := vals[0]
	for i := 1; i < len(vals); i++ {
		res *= vals[i]
	}
	return res
}

func sum(vals []int) int {
	res := 0
	for _, val := range vals {
		res += val
	}
	return res
}

func grandTotal(vals [][]int, syms []byte) int64 {
	var res int64

	for i := 0; i < len(syms); i++ {
		cur := vals[0][i]
		for j := 1; j < 4; j++ {
			if syms[i] == '+' {
				cur += vals[j][i]
			} else {
				cur *= vals[j][i]
			}
		}
		res += int64(cur)
	}

	return res
}

func valsAndSyms(input []string) ([][]int, []byte) {
	vals := [][]int{}
	syms := []byte{}

	for i := 0; i < 4; i++ {
		row := []int{}
		cur := 0
		var prev byte

		for j := 0; j < len(input[i]); j++ {
			char := input[i][j]
			if char == ' ' {
				if prev != ' ' {
					row = append(row, cur)
					cur = 0
				}
			} else {
				cur *= 10
				cur += int(char - '0')
			}
			prev = char
		}
		row = append(row, cur)
		vals = append(vals, row)
	}

	for i := 0; i < len(input[4]); i++ {
		char := input[4][i]
		if char != ' ' {
			syms = append(syms, char)
		}
	}

	return vals, syms
}
