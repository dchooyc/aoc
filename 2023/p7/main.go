package main

import (
	"aoc/lib/m"
	"aoc/lib/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := utils.ConvFile("input.txt")
	first := winnings(input, false)
	second := winnings(input, true)
	fmt.Println(first)
	fmt.Println(second)
}

func winnings(lines []string, joker bool) int {
	encodedToBid := make(map[string]int)
	handsByTypes := make([][]string, 7)
	var encoding map[byte]byte

	if joker {
		encoding = getEncodingJoker()
	} else {
		encoding = getEncoding()
	}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		encoded, handType := process(hand, encoding, joker)
		// assuming a specific hand always has same bid
		encodedToBid[encoded] = bid
		handsByTypes[handType] = append(handsByTypes[handType], encoded)
	}

	res, rank := 0, 1

	for _, hands := range handsByTypes {
		sort.Strings(hands)
		for _, hand := range hands {
			res += rank * encodedToBid[hand]
			rank++
		}
	}

	return res
}

func process(hand string, encoding map[byte]byte, joker bool) (string, int) {
	encoded := make([]byte, 5)
	tally, max := make(map[byte]int), 0

	for i := 0; i < 5; i++ {
		char := hand[i]
		encoded[i] = encoding[char]
		tally[char]++
		max = m.Max(max, tally[char])
	}

	s := string(encoded)

	if joker {
		return s, getTypeJoker(tally, max)
	}

	return s, getType(tally, max)
}

func getTypeJoker(tally map[byte]int, max int) int {
	if max == 5 {
		return 6
	} else if max == 4 {
		if tally['J'] > 0 {
			return 6
		}
		return 5
	} else if max == 3 {
		if len(tally) == 2 {
			if tally['J'] > 0 {
				return 6
			}
			return 4
		}
		if tally['J'] > 0 {
			return 5
		}
		return 3
	} else if max == 2 {
		if len(tally) == 3 {
			if tally['J'] == 2 {
				return 5
			} else if tally['J'] == 1 {
				return 4
			}
			return 2
		}
		if tally['J'] > 0 {
			return 3
		}
		return 1
	}
	if tally['J'] > 0 {
		return 1
	}
	return 0
}

func getType(tally map[byte]int, max int) int {
	if max == 5 {
		return 6
	} else if max == 4 {
		return 5
	} else if max == 3 {
		if len(tally) == 2 {
			return 4
		}
		return 3
	} else if max == 2 {
		if len(tally) == 3 {
			return 2
		}
		return 1
	}
	return 0
}

func getEncoding() map[byte]byte {
	cards := []byte{'A', 'K', 'Q', 'J', 'T'}
	res := make(map[byte]byte)

	for i := 0; i < 8; i++ {
		res[byte(i+2)+'0'] = 'a' + byte(i)
	}

	for i := 0; i < 5; i++ {
		res[cards[i]] = 'a' + byte(12-i)
	}

	return res
}

func getEncodingJoker() map[byte]byte {
	cards := []byte{'A', 'K', 'Q', 'T'}
	res := make(map[byte]byte)
	res['J'] = 'a'

	for i := 0; i < 8; i++ {
		res[byte(i+2)+'0'] = 'a' + byte(i+1)
	}

	for i := 0; i < 4; i++ {
		res[cards[i]] = 'a' + byte(12-i)
	}

	return res
}
