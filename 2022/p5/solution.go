package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := convFile("./input.txt")
	crates, moves := divide(input)
	instructs := instructions(moves)
	oneByOne := topCrates(crates, instructs)
	crane9001 := newCrane(crates, instructs)
	fmt.Println(oneByOne)
	fmt.Println(crane9001)
}

func topLayer(stacked [][]byte) string {
	res := []byte{}

	for i := 0; i < 9; i++ {
		if len(stacked[i]) > 0 {
			res = append(res, stacked[i][0])
		}
	}

	return string(res)
}

func newCrane(crates []string, instructs [][3]int) string {
	stacked := stackCrates(crates)

	for _, instruct := range instructs {
		q, s, d := instruct[0], instruct[1], instruct[2]

		load := stacked[s][:q]
		stacked[s] = stacked[s][q:]
		destination := make([]byte, len(stacked[d])+q)
		copy(destination, load)
		copy(destination[q:], stacked[d])
		stacked[d] = destination
	}

	return topLayer(stacked)
}

func topCrates(crates []string, instructs [][3]int) string {
	stacked := stackCrates(crates)

	for _, instruct := range instructs {
		q, s, d := instruct[0], instruct[1], instruct[2]

		for i := 0; i < q; i++ {
			target := stacked[s][0]
			stacked[s] = stacked[s][1:]
			destination := make([]byte, len(stacked[d])+1)
			destination[0] = target
			copy(destination[1:], stacked[d])
			stacked[d] = destination
		}
	}

	return topLayer(stacked)
}

func divide(lines []string) ([]string, []string) {
	for i, line := range lines {
		if len(line) == 0 {
			return lines[:i-1], lines[i+1:]
		}
	}
	return []string{}, []string{}
}

func instructions(moves []string) [][3]int {
	res := make([][3]int, len(moves))

	for i, move := range moves {
		a := strings.Split(move, " ")
		quant, _ := strconv.Atoi(a[1])
		src, _ := strconv.Atoi(a[3])
		dst, _ := strconv.Atoi(a[5])
		res[i][0] = quant
		res[i][1] = src - 1
		res[i][2] = dst - 1
	}

	return res
}

func stackCrates(crates []string) [][]byte {
	res := make([][]byte, 9)

	for _, crate := range crates {
		i, index := 1, 0

		for i < len(crate) {
			if crate[i] != ' ' {
				res[index] = append(res[index], crate[i])
			}
			index++
			i += 4
		}
	}

	return res
}

func convFile(src string) []string {
	res := []string{}

	file, err := os.Open(src)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
