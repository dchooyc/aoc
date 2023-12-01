package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := convFile("./input.txt")
	score := calcScore(input)
	outcome := calcScoreByOutcome(input)

	fmt.Println("This is the wrongly interpreted score:")
	fmt.Println(score)
	fmt.Println("This is the correctly interpreted score:")
	fmt.Println(outcome)
}

func calcScore(plays []string) int {
	score := 0
	conv := map[byte]byte{'A': 'X', 'B': 'Y', 'C': 'Z'}
	wins := map[byte]byte{'X': 'Z', 'Y': 'X', 'Z': 'Y'}

	for _, play := range plays {
		oppo, me, cur := conv[play[0]], play[2], 0

		if me == 'X' {
			cur += 1
		} else if me == 'Y' {
			cur += 2
		} else {
			cur += 3
		}

		if oppo == me {
			cur += 3
		} else if wins[me] == oppo {
			cur += 6
		}

		score += cur
	}

	return score
}

func calcScoreByOutcome(plays []string) int {
	score := 0
	loses := map[byte]int{'A': 3, 'B': 1, 'C': 2}
	value := map[byte]int{'A': 1, 'B': 2, 'C': 3}
	wins := map[byte]int{'A': 2, 'B': 3, 'C': 1}

	for _, play := range plays {
		oppo, end, cur := play[0], play[2], 0

		if end == 'X' {
			cur += loses[oppo]
		} else if end == 'Y' {
			cur += value[oppo] + 3
		} else {
			cur += wins[oppo] + 6
		}

		score += cur
	}

	return score
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
